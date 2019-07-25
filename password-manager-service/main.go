package main

import (
    pb "github.com/jonb377/website/password-manager-service/proto/password-manager"
    auth "github.com/jonb377/website/common"
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


    db.AutoMigrate(&PasswordEntry{})

    srv := micro.NewService(
        micro.Name("go.micro.api.password_manager"),
        micro.Version("latest"),
        micro.WrapHandler(auth.AuthWrapper),
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

