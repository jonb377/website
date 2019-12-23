package storage

import (
    pb "github.com/jonb377/website/storage-service/proto/storage"
    moni "github.com/jonb377/website/monitor-service/monitor"
    "fmt"
    "github.com/micro/go-micro"
    "github.com/micro/go-micro/server"
    "log"
    _ "github.com/micro/go-plugins/broker/rabbitmq"
    _ "github.com/micro/go-plugins/registry/kubernetes"
)

const serviceName = "go.micro.api.storage"

func RunStorageService() {
    db, err := CreateConnection()

    if err != nil {
        log.Fatalf("Could not connect to DB: %v", err)
    }

    defer db.Close()


    db.AutoMigrate(&Blob{})

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

    if err := pb.RegisterStorageHandler(srv.Server(), &StorageService{
        db: db,
    }); err != nil {
        fmt.Println(err)
    }

    if err := srv.Run(); err != nil {
        fmt.Println(err)
    }
}
