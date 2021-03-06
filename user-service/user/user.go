package user

import (
    pb "github.com/jonb377/website/user-service/proto/user"
    moni "github.com/jonb377/website/monitor-service/monitor"
    "fmt"
    "github.com/micro/go-micro"
    "github.com/micro/go-micro/server"
    "github.com/micro/micro/plugin"
    rpc "github.com/micro/go-plugins/micro/disable_rpc"
    "log"
    _ "github.com/micro/go-plugins/broker/rabbitmq"
    _ "github.com/micro/go-plugins/registry/kubernetes"
)

const serviceName = "go.micro.api.user"

func RunUserService() {
    // Disable requests to /rpc
    plugin.Register(rpc.NewPlugin())

    db, err := CreateConnection()

    if err != nil {
        log.Fatalf("Could not connect to DB: %v", err)
    }

    defer db.Close()

    db.AutoMigrate(&User{})
    db.AutoMigrate(&Device{})
    db.AutoMigrate(&AccessKey{})
    db.AutoMigrate(&UserKey{})

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

    if err := pb.RegisterUserHandler(srv.Server(), &service{db}); err != nil {
        fmt.Println(err)
    }

    if err := srv.Run(); err != nil {
        fmt.Println(err)
    }
}
