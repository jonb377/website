package passwordmanager

import (
    pb "github.com/jonb377/website/password-manager-service/proto/password-manager"
    common "github.com/jonb377/website/common"
    "fmt"
    "github.com/micro/go-micro"
    "github.com/micro/go-micro/server"
    "log"
)

const serviceName = "go.micro.api.passwordmanager"

func RunPasswordManagerService() {
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
        server.Server(
            server.Name(serviceName),
            server.Address(":8080"),
        ),
    )

    srv.Init()

    if err := pb.RegisterPasswordManagerServiceHandler(srv.Server(), &PasswordManagerService{
        db: db,
    }); err != nil {
        fmt.Println(err)
    }

    if err := srv.Run(); err != nil {
        fmt.Println(err)
    }
}
