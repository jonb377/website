package auth

import (
    userProto "github.com/jonb377/website/user-service/proto/user"
    pb "github.com/jonb377/website/auth-service/proto/auth"
    "fmt"
    "github.com/micro/go-micro"
    "github.com/micro/go-micro/client"
    "github.com/micro/go-micro/server"
    "github.com/micro/micro/plugin"
    rpc "github.com/micro/go-plugins/micro/disable_rpc"
    "log"
    _ "github.com/micro/go-plugins/broker/rabbitmq"
    _ "github.com/micro/go-plugins/registry/kubernetes"
)

const serviceName = "go.micro.api.auth"

func RunAuthService() {
    // Disable requests to /rpc
    plugin.Register(rpc.NewPlugin())

    db, err := CreateConnection()

    if err != nil {
        log.Fatalf("Could not connect to DB: %v", err)
    }

    defer db.Close()


    db.AutoMigrate(&Session{})

    srv := micro.NewService(
        micro.Name(serviceName),
        micro.Version("latest"),
        micro.Server(
            server.NewServer(
                server.Name(serviceName),
                server.Address(":8080"),
            ),
        ),
    )

    srv.Init()

    if err := pb.RegisterAuthHandler(srv.Server(), &AuthService{
        db: nil,
        userClient: userProto.NewUserService("go.micro.api.user", client.DefaultClient),
        tokenService: &TokenService{},
    }); err != nil {
        fmt.Println(err)
    }

    if err := srv.Run(); err != nil {
        fmt.Println(err)
    }
}
