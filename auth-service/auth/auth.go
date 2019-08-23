package auth

import (
    userProto "github.com/jonb377/website/user-service/proto/user"
    pb "github.com/jonb377/website/auth-service/proto/auth"
    moni "github.com/jonb377/website/monitor-service/monitor"
    "fmt"
    "github.com/micro/go-micro"
    "github.com/micro/go-micro/client"
    "github.com/micro/go-micro/server"
    "log"
    _ "github.com/micro/go-plugins/broker/rabbitmq"
    _ "github.com/micro/go-plugins/registry/kubernetes"
)

const serviceName = "go.micro.api.auth"

func RunAuthService() {
    log.Println("Registering service ", serviceName)

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
        micro.WrapHandler(moni.RPCTraceWrapper(serviceName)),
    )

    srv.Init()

    if err := pb.RegisterAuthHandler(srv.Server(), &AuthService{
        db: db,
        userClient: userProto.NewUserService("go.micro.api.user", client.DefaultClient),
        tokenService: &TokenService{},
    }); err != nil {
        fmt.Println(err)
    }

    if err := srv.Run(); err != nil {
        fmt.Println(err)
    }
}
