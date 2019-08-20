package monitor

import (
    moniProto "github.com/jonb377/website/monitor-service/proto/monitor"
    "github.com/micro/go-micro/client"
    "net/http"
    "context"
    "strings"
)

// Send a push notification with activity details
func MonitorNotifyWrapper(next http.Handler) http.Handler {
    moniClient := moniProto.NewMonitorService("go.micro.api.monitor", client.DefaultClient)
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if !strings.HasPrefix(r.URL.Path, "/amy") && !strings.HasPrefix(r.URL.Path, "/media") {
            var ip string
            if len(r.Header["X-Real-Ip"]) > 0 {
                ip = r.Header["X-Real-Ip"][0]
            } else if len(r.Header["X-Forwarded-For"]) > 0{
                ip = r.Header["X-Forwarded-For"][0]
            }

            var host string
            if len(r.Header["X-Host"]) > 0 {
                host = r.Header["X-Host"][0]
            }

            req := moniProto.Activity{
                Ip: ip,
                Route: r.URL.Path,
                Host: host,
            }
            go moniClient.NotifyActivity(context.Background(), &req)
        }
        next.ServeHTTP(w, r)
    })
}
