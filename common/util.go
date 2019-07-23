package main

import (
	authService "../auth-service/proto/auth"
    "context"
    "errors"
    "github.com/micro/go-micro/client"
    "github.com/micro/go-micro/metadata"
    "github.com/micro/go-micro/server"
    "log"
)


type myRequest struct {
    server.Request
    sessionKey []byte
}

/*
Wrap the Request Read method to decrypt the body
 */
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


func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
    return func(ctx context.Context, req server.Request, resp interface{}) error {
        meta, ok := metadata.FromContext(ctx)
        if !ok {
            return errors.New("no auth meta-data found in request")
        }

        var err error
        tokenString, ok := meta["Token"]
        if ok {
            device, ok := meta["Device"]
            if !ok {
                return errors.New("device header messing")
            }
            log.Println("Authenticating with token: ", tokenString)

            authClient := authService.NewAuthService("auth", client.DefaultClient)
            tokenResponse, err := authClient.ValidateToken(context.Background(), &authService.ValidateTokenRequest{
                Token: tokenString,
                Device: device,
            })
            if err != nil {
                return err
            }
            err = fn(ctx, &myRequest{
                Request: req,
                sessionKey: tokenResponse.SessionKey,
            }, resp)
        } else {
            err = fn(ctx, req, resp)
        }
        return err
    }
}
