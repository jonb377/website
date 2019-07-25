package main

import (
    userProto "github.com/jonb377/website/user-service/proto/user"
    pb "github.com/jonb377/website/auth-service/proto/auth"
    "fmt"
    "github.com/micro/go-micro"
    "github.com/micro/go-micro/client"
    "log"
)

func main() {
    db, err := CreateConnection()

    if err != nil {
        log.Fatalf("Could not connect to DB: %v", err)
    }

    defer db.Close()


    db.AutoMigrate(&Session{})

    srv := micro.NewService(
        micro.Name("go.micro.api.auth"),
        micro.Version("latest"),
    )

    srv.Init()



    if err := pb.RegisterAuthServiceHandler(srv.Server(), &AuthService{
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

