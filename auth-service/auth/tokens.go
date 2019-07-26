package auth

import (
    "crypto/rand"
    "github.com/dgrijalva/jwt-go"
    "time"
)

var tokenkey []byte

func init() {
    tokenkey = make([]byte, 16)
    if _, err := rand.Read(tokenkey); err != nil {
        panic("Failed to generate key")
    }
}

type CustomClaims struct {
    jwt.StandardClaims
    SessionId string
    Username string
    Device string
}

type Authable interface {
    Decode(token string) (*CustomClaims, error)
    Encode(session *Session) (string, error)
}

type TokenService struct {}

func (srv *TokenService) Decode(tokenString string) (*CustomClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        return tokenkey, nil
    })

    if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
        return claims, nil
    } else {
        return nil, err
    }
}

func (srv *TokenService) Encode(session *Session) (string, error) {
    expireToken := time.Now().Add(time.Hour).Unix()

    claims := CustomClaims{
        SessionId: session.SessionID,
        Username: session.Username,
        Device: session.Device,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expireToken,
            Issuer: "auth",
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    return token.SignedString(tokenkey)
}
