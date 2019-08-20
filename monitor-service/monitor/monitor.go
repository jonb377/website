package monitor

import (
    notiProto "github.com/jonb377/website/notifications-service/proto/notifications"
    userProto "github.com/jonb377/website/user-service/proto/user"
    pb "github.com/jonb377/website/monitor-service/proto/monitor"
    "fmt"
    "github.com/micro/go-micro"
    "github.com/micro/go-micro/client"
    "github.com/micro/go-micro/server"
    "log"
    "os"
    "strings"
    "context"
    "io/ioutil"
    _ "github.com/micro/go-plugins/broker/rabbitmq"
    _ "github.com/micro/go-plugins/registry/kubernetes"
)

const serviceName = "go.micro.api.monitor"

func RunMonitorService() {
    log.Println("Registering service ", serviceName)

    db, err := CreateConnection()

    if err != nil {
        log.Fatalf("Could not connect to DB: %v", err)
    }

    defer db.Close()

    db.AutoMigrate(&Request{})

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

    // Read the ipstack access key
    ipstackFile, err := os.Open("/var/srv/ipstack_access_key.txt")
    if err != nil {
        log.Fatalf("Failed to open access key file: %v\n", err.Error())
    }
    data, _ := ioutil.ReadAll(ipstackFile)
    accessKey := strings.TrimSpace(string(data))

    // Get admin users
    userClient := userProto.NewUserService("go.micro.api.user", client.DefaultClient)
    admins, err := userClient.GetAdmins(context.Background(), &userProto.Empty{})
    if err != nil {
        log.Fatalf("Failed to load users: %v\n", err.Error())
    }


    notiClient := notiProto.NewNotificationsService("go.micro.api.notifications", client.DefaultClient)

    if err := pb.RegisterMonitorHandler(srv.Server(), &MonitorService{
                db: db,
                notiClient: notiClient,
                admins: admins.Admins,
                ipstackAccessKey: accessKey,
            }); err != nil {
        fmt.Println(err)
    }

    if err := srv.Run(); err != nil {
        fmt.Println(err)
    }
}
