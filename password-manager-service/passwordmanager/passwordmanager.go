package passwordmanager

import (
    pb "github.com/jonb377/website/password-manager-service/proto/password-manager"
    "fmt"
    "github.com/micro/go-micro"
    "github.com/micro/go-micro/server"
    "log"
    _ "github.com/micro/go-plugins/broker/rabbitmq"
    _ "github.com/micro/go-plugins/registry/kubernetes"
)

const serviceName = "go.micro.api.passwordmanager"

func RunPasswordManagerService() {
    db, err := CreateConnection()

    if err != nil {
        log.Fatalf("Could not connect to DB: %v", err)
    }

    defer db.Close()


    db.AutoMigrate(&PasswordEntry{})
    db.Model(&PasswordEntry{}).AddUniqueIndex("user_domain_username", "user", "domain", "username")

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

    if err := pb.RegisterPasswordManagerHandler(srv.Server(), &PasswordManagerService{
        db: db,
    }); err != nil {
        fmt.Println(err)
    }

    if err := srv.Run(); err != nil {
        fmt.Println(err)
    }
}
