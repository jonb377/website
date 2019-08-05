package user

import (
    router "github.com/jonb377/website/router-service/router"
    proto "github.com/jonb377/website/user-service/proto/user"
    "github.com/gorilla/mux"
    "github.com/micro/go-micro/client"
)

func Route(r *mux.Router) {
    user := proto.NewUserService("go.micro.api.user", client.DefaultClient)
    r.HandleFunc(
        "/register",
        router.RPCCall(
            user.Register,
            &proto.RegisterRequest{},
        ),
    ).Methods("PUT")
    r.HandleFunc(
        "/devices/add",
        router.RPCCall(
            user.AddDevice,
            &proto.Empty{},
        ),
    ).Methods("PUT")
    r.HandleFunc(
        "/access-key",
        router.RPCCall(
            user.GetAccessKey,
            &proto.Empty{},
        ),
    ).Methods("POST")
    r.Use(router.AuthWrapper)
    r.Use(router.LogWrapper)
}
