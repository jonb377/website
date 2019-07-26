package main

import (
    userProto "github.com/jonb377/website/user-service/proto/user"
    pb "github.com/jonb377/website/auth-service/proto/auth"
    auth "github.com/jonb377/website/auth-service/auth"
    "fmt"
    "github.com/micro/go-micro"
    "github.com/micro/go-micro/client"
    "log"
)

func main() {
    db, err := auth.CreateConnection()

    if err != nil {
        log.Fatalf("Could not connect to DB: %v", err)
    }

    defer db.Close()


    db.AutoMigrate(&auth.Session{})

    srv := micro.NewService(
        micro.Name("go.micro.api.auth"),
        micro.Version("latest"),
    )

    srv.Init()



    if err := pb.RegisterAuthServiceHandler(srv.Server(), &auth.AuthService{
        db: db,
        userClient: userProto.NewUserService("go.micro.api.user", client.DefaultClient),
        tokenService: &auth.TokenService{},
    }); err != nil {
        fmt.Println(err)
    }

    if err := srv.Run(); err != nil {
        fmt.Println(err)
    }
}

