package main

import (
    //jbrouter "github.com/jonb377/website/router-service/router"
	"github.com/micro/go-micro/util/log"
    "github.com/gorilla/mux"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/api/resolver"
	"github.com/micro/go-micro/api/router"
    "github.com/micro/go-micro/registry"
    "github.com/micro/go-micro/registry/consul"
	ahandler "github.com/micro/go-micro/api/handler"
	rrmicro "github.com/micro/go-micro/api/resolver/micro"
	arpc "github.com/micro/go-micro/api/handler/rpc"
	regRouter "github.com/micro/go-micro/api/router/registry"
	httpapi "github.com/micro/go-micro/api/server/http"
)

const (
    Name        = "go.micro.api"
    Address     = ":8080"
    Handler     = "api"
    Resolver    = "micro"
    APIPath     = "/api"
    Namespace   = "go.micro.api"
    serviceName = "go.micro.api.router"
)

func main() {
    r := mux.NewRouter()

    // Register auth middleware
    //r.Use(jbrouter.AuthWrapper)

    // Create the service
    service := micro.NewService(
        micro.Registry(
            consul.NewRegistry(
                registry.Addrs("consul"),
            ),
        ),
        micro.Name(serviceName),
    )

    ropts := []resolver.Option{
        resolver.WithNamespace(Namespace),
        resolver.WithHandler(Handler),
    }

    rr := rrmicro.NewResolver(ropts...)

    log.Logf("Registering API RPC Handler at %s", APIPath)
    rt := regRouter.NewRouter(
        router.WithNamespace(Namespace),
        router.WithHandler(arpc.Handler),
        router.WithResolver(rr),
        router.WithRegistry(service.Options().Registry),
    )
    rp := arpc.NewHandler(
        ahandler.WithNamespace(Namespace),
        ahandler.WithRouter(rt),
        ahandler.WithService(service),
    )
    r.PathPrefix(APIPath).Handler(rp)

    api := httpapi.NewServer(Address)
	api.Init()
	api.Handle("/", r)

	// Start API
	if err := api.Start(); err != nil {
		log.Fatal(err)
	}

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

	// Stop API
	if err := api.Stop(); err != nil {
		log.Fatal(err)
	}
}
