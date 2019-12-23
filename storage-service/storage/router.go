package storage

import (
    router "github.com/jonb377/website/router-service/router"
    proto "github.com/jonb377/website/storage-service/proto/storage"
    "github.com/gorilla/mux"
    "github.com/micro/go-micro/client"
)

func Route(r *mux.Router) {
    storage := proto.NewStorageService("go.micro.api.storage", client.DefaultClient)
    r.HandleFunc(
        "/{olduri}/move-to/{newuri}",
        router.AuthenticatedRPCCall(
            storage.Rename,
            &proto.Empty{},
        ),
    ).Methods("POST")
    r.HandleFunc(
        "/{uri}",
        router.AuthenticatedRPCCall(
            storage.SaveBlob,
            &proto.Blob{},
        ),
    ).Methods("POST")
    r.HandleFunc(
        "/{uri}",
        router.AuthenticatedRPCCall(
            storage.DeleteBlob,
            &proto.Blob{},
        ),
    ).Methods("DELETE")
    r.HandleFunc(
        "/",
        router.AuthenticatedRPCCall(
            storage.Sync,
            &proto.SyncRequest{},
        ),
    ).Methods("GET")
    r.Use(router.AuthWrapper)
}
