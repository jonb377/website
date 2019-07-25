package common

import (
    authService "github.com/jonb377/website/auth-service/proto/auth"
    "context"
    "errors"
    "github.com/micro/go-micro/client"
    "github.com/micro/go-micro/metadata"
    "github.com/micro/go-micro/server"
    "log"
)


type AuthCheck interface {
    Authenticated() bool
}

type authenticated struct {
    a bool
}

func (a *authenticated) Authenticated() bool {
    return a.a
}


type myRequest struct {
    server.Request
    sessionKey []byte
}

// Wrap the Request Read method to decrypt the body
func (r *myRequest) Read() ([]byte, error) {
    cipher := NewAESCipher(r.sessionKey)
    body, err := r.Request.Read()
    if err != nil {
        return nil, err
    }
    res, err := cipher.Decrypt(body)
    if err != nil {
        return nil, err
    }
    return res, nil
}


// Wraps inbound requests. Decrypts inbound and outbound data if a session is active.
// Adds "authenticated" to the context, which indicates whether an active session was found.
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
    return func(ctx context.Context, req server.Request, resp server.Response) error {
        meta, ok := metadata.FromContext(ctx)
        if !ok {
            return errors.New("no auth meta-data found in request")
        }

        var err error
        tokenString, ok := meta["Token"]
        if ok {
            device, ok := meta["Device"]
            if !ok {
                return errors.New("device header missing")
            }
            log.Println("Authenticating with token: ", tokenString)

            authClient := authService.NewAuthService("auth", client.DefaultClient)
            tokenResponse, err := authClient.ValidateToken(context.Background(), &authService.ValidateTokenRequest{
                Token: tokenString,
                Device: device,
            })
            if err != nil {
                log.Println("Authentication failed: ", err)
                return err
            }
            err = fn(context.WithValue(ctx, "authenticated", &authenticated{true}), &myRequest{
                Request: req,
                sessionKey: tokenResponse.SessionKey,
            }, resp)
            if err == nil {
                // Encrypt the response
            }
        } else {
            err = fn(context.WithValue(ctx, "authenticated", &authenticated{false}), req, resp)
        }
        return err
    }
}
