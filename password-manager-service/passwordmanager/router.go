package passwordmanager

import (
    router "github.com/jonb377/website/router-service/router"
    proto "github.com/jonb377/website/password-manager-service/proto/password-manager"
    "github.com/gorilla/mux"
    "github.com/micro/go-micro/client"
)

func Route(r *mux.Router) {
    passman := proto.NewPasswordManagerService("go.micro.api.passwordmanager", client.DefaultClient)
    r.HandleFunc(
        "/update",
        router.AuthenticatedRPCCall(
            passman.UpdatePassword,
            &proto.UpdatePasswordRequest{},
        ),
    ).Methods("POST")
    r.HandleFunc(
        "/list",
        router.AuthenticatedRPCCall(
            passman.ListPasswords,
            &proto.ListPasswordRequest{},
        ),
    ).Methods("POST")
    r.Use(router.AuthWrapper)
}
