package auth

import (
    router "github.com/jonb377/website/router-service/router"
    proto "github.com/jonb377/website/auth-service/proto/auth"
    "github.com/gorilla/mux"
    "github.com/micro/go-micro/client"
)

func Route(r *mux.Router) {
    auth := proto.NewAuthService("go.micro.api.auth", client.DefaultClient)
    r.HandleFunc(
        "/connection/create",
        router.RPCCall(
            auth.CreateConnection,
            &proto.CreateConnectionRequest{},
        ),
    ).Methods("POST")
    r.HandleFunc(
        "/connection/challenge",
        router.RPCCall(
            auth.ConnectionChallenge,
            &proto.ConnectionChallengeRequest{},
        ),
    ).Methods("POST")
    r.HandleFunc(
        "/connection/close",
        router.RPCCall(
            auth.CloseConnection,
            &proto.Empty{},
        ),
    ).Methods("POST")
    r.Use(router.AuthWrapper)
    r.Use(router.LogWrapper)
}
