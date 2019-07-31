package router

import (
    authService "github.com/jonb377/website/auth-service/proto/auth"
    "github.com/micro/go-micro/client"
    "io/ioutil"
    "net/http"
    "context"
    "bytes"
    "log"
    "fmt"
)

type myResponseWriter struct {
    headers http.Header
    body []byte
    status int
}

func (w *myResponseWriter) Header() http.Header {
    return w.headers
}

func (w *myResponseWriter) Write(data []byte) (int, error) {
    w.body = append(w.body, data...)
    return len(data), nil
}

func (w *myResponseWriter) WriteHeader(code int) {
    w.status = code
}


// Wraps inbound requests. Decrypts inbound and outbound data if a session is active.
// Adds "authenticated", which indicates whether an active session was found.
func AuthWrapper(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        var err error
        fmt.Println("Calling the auth wrapper")

        // Disallow session-id header
        if sessid := r.Header["Session-Id"]; len(sessid) != 0 {
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
        tokenCookie, err := r.Cookie("Token")
        if err != nil {
            // No active token in request
            log.Println("No auth token provided")
            next.ServeHTTP(w, r)
            return
        }
        tokenString := tokenCookie.String()

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
        r.Body = ioutil.NopCloser(bytes.NewReader(decrypted))

        // Use a dummy response writer for the rest of the call
        respWriter := myResponseWriter{}
        next.ServeHTTP(&respWriter, r)

        // Copy from the dummy
        if respWriter.status != 0 {
            // WriteHeader explicitly called
            w.WriteHeader(respWriter.status)
        }
        for k, v := range respWriter.headers {
            w.Header()[k] = v
        }
        encrypted, err := cipher.Encrypt(respWriter.body)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        w.Write(encrypted)
    })
}
