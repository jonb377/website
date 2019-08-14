package main

import (
    auth "github.com/jonb377/website/auth-service/auth"
    user "github.com/jonb377/website/user-service/user"
    passman "github.com/jonb377/website/password-manager-service/passwordmanager"
    websrv "github.com/jonb377/website/web-service/web"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
    "github.com/micro/go-micro/registry"
    "github.com/micro/go-micro/registry/consul"
    "github.com/gorilla/mux"
)

const (
    serviceName = "go.micro.api.router"
)

func main() {
    // Create the service
    service := web.NewService(
        web.Registry(
            consul.NewRegistry(
                registry.Addrs("consul"),
            ),
        ),
        web.Name(serviceName),
        web.Address("0.0.0.0:8080"),
    )

    r := mux.NewRouter()

    // Register service handlers
    api := r.PathPrefix("/api").Subrouter()
    auth.Route(api.PathPrefix("/auth").Subrouter())
    user.Route(api.PathPrefix("/user").Subrouter())
    passman.Route(api.PathPrefix("/passwordmanager").Subrouter())

    // Register web handler
    websrv.Route(r.PathPrefix("/").Subrouter())

    service.Handle("/", r)

    // Initialize server
    if err := service.Init(); err != nil {
        log.Fatal(err)
    }

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
