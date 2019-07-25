package main

import (
    pb "github.com/jonb377/website/user-service/proto/user"
    auth "github.com/jonb377/website/auth-service"
    "fmt"
    "github.com/micro/go-micro"
    "log"
)

func main() {
    db, err := CreateConnection()

    if err != nil {
        log.Fatalf("Could not connect to DB: %v", err)
    }

    defer db.Close()

    db.AutoMigrate(&User{})
    db.AutoMigrate(&Device{})

    repo := &UserRepository{db}

    srv := micro.NewService(
        micro.Name("go.micro.api.user"),
        micro.Version("latest"),
        micro.WrapHandler(auth.AuthWrapper),
    )

    srv.Init()

    if err := pb.RegisterUserServiceHandler(srv.Server(), &service{repo}); err != nil {
        fmt.Println(err)
    }

    if err := srv.Run(); err != nil {
        fmt.Println(err)
    }
}
