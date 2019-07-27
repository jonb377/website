package user

import (
    pb "github.com/jonb377/website/user-service/proto/user"
    common "github.com/jonb377/website/common"
    "fmt"
    "github.com/micro/go-micro"
    "github.com/micro/go-micro/server"
    "log"
)

const serviceName = "go.micro.api.user"

func RunUserService() {
    db, err := CreateConnection()

    if err != nil {
        log.Fatalf("Could not connect to DB: %v", err)
    }

    defer db.Close()

    db.AutoMigrate(&User{})
    db.AutoMigrate(&Device{})

    repo := &UserRepository{db}

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

    if err := pb.RegisterUserServiceHandler(srv.Server(), &service{repo}); err != nil {
        fmt.Println(err)
    }

    if err := srv.Run(); err != nil {
        fmt.Println(err)
    }
}
