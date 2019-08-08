package router

import (
    authService "github.com/jonb377/website/auth-service/proto/auth"
    "github.com/micro/go-micro/util/ctx"
    "github.com/micro/go-micro/metadata"
    "github.com/micro/go-micro/client"
    "encoding/json"
    "reflect"
    "io/ioutil"
    "net/http"
    "context"
    "errors"
    "bytes"
    "time"
    "log"
    "fmt"
)

type authResponseWriter struct {
    headers http.Header
    body []byte
    status int
}

func (w *authResponseWriter) Header() http.Header {
    return w.headers
}

func (w *authResponseWriter) Write(data []byte) (int, error) {
    w.body = append(w.body, data...)
    return len(data), nil
}

func (w *authResponseWriter) WriteHeader(code int) {
    w.status = code
}

type logResponseWriter struct {
    http.ResponseWriter
    status int
}

func (w logResponseWriter) WriteHeader(code int) {
    w.status = code
    w.ResponseWriter.WriteHeader(code)
}

func LogWrapper(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        logWriter := logResponseWriter{w, 200}
        next.ServeHTTP(logWriter, r)
        duration := time.Now().Sub(start)
        log.Println("api: ", r.URL.Path, "\tstatus: ", logWriter.status, "\tduration: ", duration.String())
    })
}


// Wraps inbound requests. Decrypts inbound and outbound data if a session is active.
// Adds "authenticated", which indicates whether an active session was found.
func AuthWrapper(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        var err error
        log.Println("Calling the auth wrapper ", r.URL.Path, fmt.Sprintf("%p", r))

        // Disallow session-id and username header
        if sessid := r.Header["Session-Id"]; len(sessid) != 0 {
            http.Error(w, "Session-Id header not allowed", http.StatusBadRequest)
            return
        }
        if username := r.Header["Username"]; len(username) != 0 {
            http.Error(w, "Session-Id header not allowed", http.StatusBadRequest)
            return
        }

        // Check the device header
        deviceHeader := r.Header["Device"]
        if len(deviceHeader) != 1 {
            http.Error(w, "device header missing", http.StatusBadRequest)
            return
        }
        device := deviceHeader[0]
        tokenHeader := r.Header["Token"]
        if len(tokenHeader) != 1 {
            // No active token in request
            log.Println("No auth token provided")
            next.ServeHTTP(w, r)
            return
        }
        tokenString := tokenHeader[0]

        log.Println("Authenticating with token: ", tokenString)

        // Get the data from the token
        authClient := authService.NewAuthService("go.micro.api.auth", client.DefaultClient)
        tokenResponse, err := authClient.ValidateToken(context.Background(), &authService.ValidateTokenRequest{
            Token: tokenString,
            Device: device,
        })
        if err != nil {
            log.Println("Authentication failed: ", err)
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        // Set the session-id header to indicate successful authentication
        r.Header["Session-Id"] = []string{tokenResponse.SessionID}
        r.Header["Username"] = []string{tokenResponse.Username}

        if len(tokenResponse.SessionKey) == 0 {
            if r.URL.Path == "/api/auth/connection/challenge" {
                // No active token in request
                log.Println("Token not active")
                next.ServeHTTP(w, r)
                return
            } else {
                log.Println("Session key not set: ", r.URL.Path)
                http.Error(w, "Session key not set", http.StatusBadRequest)
                return
            }
        }
        // Use the session key to decrypt the data
        cipher := NewAESCipher(tokenResponse.SessionKey)
        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        // Overwrite the old body with the decrypted value
        decrypted, err := cipher.Decrypt(body)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        log.Println("Decrypted message body ", string(decrypted))
        r.Body = ioutil.NopCloser(bytes.NewReader(decrypted))

        // Use a dummy response writer for the rest of the call
        respWriter := authResponseWriter{headers: make(map[string][]string)}
        next.ServeHTTP(&respWriter, r)

        // Copy from the dummy
        if respWriter.status != 0 {
            // WriteHeader explicitly called
            w.WriteHeader(respWriter.status)
        }
        for k, v := range respWriter.headers {
            log.Println("header ", k, ": ", v)
            if k != "Content-Length" {
                w.Header()[k] = v
            }
        }
        if len(respWriter.body) == 0 {
            return
        }
        log.Println("encrypting message ", string(respWriter.body), "\tlength ", len(respWriter.body))
        encrypted, err := cipher.Encrypt(respWriter.body)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        log.Println("encrypted message ", string(encrypted), "\tlength ", len(encrypted))
        w.Write(encrypted)
    })
}

// Convert an HTTP request to an internal RPC call
func RPCCall(f interface{}, req interface{}) func(http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Println("In RPCCall")
        log.Printf("Function type: %T\n", f)

        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
            log.Println("Error: ", err.Error())
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        json.Unmarshal(body, req)
        log.Println("Body: ", req)

        cx := ctx.FromRequest(r)
        vf := reflect.ValueOf(f)
        vres := vf.Call([]reflect.Value{reflect.ValueOf(cx), reflect.ValueOf(req)})
        resp, ierr := vres[0].Interface(), vres[1].Interface()

        if ierr != nil {
            log.Println("Error: ", ierr.(error).Error())
            http.Error(w, ierr.(error).Error(), http.StatusInternalServerError)
            return
        }
        data, err := json.Marshal(resp)
        if err != nil {
            log.Println("Error: ", err.Error())
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        log.Println("Result: ", string(data))
        w.Write(data)
    }
}


func RequireAuth(ctx context.Context) (string, string, string, error) {
    md, ok := metadata.FromContext(ctx)
    if !ok {
        md = metadata.Metadata{}
    }
    session_id := md["Session-Id"]
    username := md["Username"]
    device := md["Device"]
    if session_id == "" || username == "" {
        return "", "", "", errors.New("unauthorized")
    }
    return username, device, session_id, nil
}
