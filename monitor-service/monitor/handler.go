package monitor

import (
    notiProto "github.com/jonb377/website/notifications-service/proto/notifications"
    pb "github.com/jonb377/website/monitor-service/proto/monitor"
    "context"
    "github.com/jinzhu/gorm"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "fmt"
    "time"
    "log"
)

type MonitorService struct {
    db                  *gorm.DB
    notiClient          notiProto.NotificationsService
    admins              []string
    ipstackAccessKey    string
}

func (s *MonitorService) NotifyActivity(ctx context.Context, req *pb.Activity, _ *pb.Empty) error {
    // Create an entry in the database
    request := Request{
        Ip: req.Ip,
        Route: req.Route,
        CreatedAt: time.Now().Unix(),
    }
    if err := s.db.Create(&request).Error; err != nil {
        return err
    }

    var ipdata map[string]interface{} = nil
    if req.Ip != "" {
        // Get information about the ip address
        resp, err := http.Get(fmt.Sprintf("http://api.ipstack.com/%v?access_key=%v", req.Ip, s.ipstackAccessKey))
        if err == nil {
            data, err := ioutil.ReadAll(resp.Body)
            if err == nil {
                json.Unmarshal(data, &ipdata)
            } else {
                log.Printf("Error: %v\n", err.Error())
            }
        } else {
            log.Printf("Error: %v\n", err.Error())
        }
    }
    if ipdata == nil {
        log.Println("Failed to get IP address info :(")
        ipdata = map[string]interface{}{"country_name": "null", "region_name": "null", "city_name": "null"}
    }


    // Send a push notification about the activity
    for _, username := range s.admins {
        message := notiProto.SendNotificationRequest{
            Title: "New Activity",
            Body: fmt.Sprintf(
                "host: %v\nip: %v\nroute: %v\ncountry: %v\nregion: %v\ncity: %v",
                req.Host,
                req.Ip,
                req.Route,
                ipdata["country_name"],
                ipdata["region_name"],
                ipdata["city"],
            ),
            Username: username,
        }
        if _, err := s.notiClient.SendNotification(context.Background(), &message); err != nil {
            log.Println("Error pushing to ", username, ": ", err.Error())
        }
    }
    return nil
}
