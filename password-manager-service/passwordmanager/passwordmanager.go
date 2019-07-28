package passwordmanager

import (
    pb "github.com/jonb377/website/password-manager-service/proto/password-manager"
    common "github.com/jonb377/website/common"
    "fmt"
    "github.com/micro/go-micro"
    "github.com/micro/go-micro/server"
    "github.com/micro/micro/plugin"
    rpc "github.com/micro/go-plugins/micro/disable_rpc"
    "log"
    _ "github.com/micro/go-plugins/broker/rabbitmq"
    _ "github.com/micro/go-plugins/registry/kubernetes"
)

const serviceName = "go.micro.api.passwordmanager"

func RunPasswordManagerService() {
    // Disable requests to /rpc
    plugin.Register(rpc.NewPlugin())

    db, err := CreateConnection()

    if err != nil {
        log.Fatalf("Could not connect to DB: %v", err)
    }

    defer db.Close()


    db.AutoMigrate(&PasswordEntry{})

    srv := micro.NewService(
        micro.Name(serviceName),
        micro.Version("latest"),
        micro.WrapHandler(common.AuthWrapper),
        micro.Server(
            server.NewServer(
                server.Name(serviceName),
                server.Address(":8080"),
            ),
        ),
    )

    srv.Init()

    if err := pb.RegisterPasswordManagerHandler(srv.Server(), &PasswordManagerService{
        db: db,
    }); err != nil {
        fmt.Println(err)
    }

    if err := srv.Run(); err != nil {
        fmt.Println(err)
    }
}
