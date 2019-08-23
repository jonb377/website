package monitor

import (
    pb "github.com/jonb377/website/monitor-service/proto/monitor"
    "github.com/micro/go-micro/client"
    "github.com/micro/go-micro/metadata"
    "github.com/micro/go-micro/server"
    "net/http"
    "context"
    "strings"
    "log"
)

type statusWrapper struct {
    http.ResponseWriter
    status int32
    length int64
}

func (w *statusWrapper) WriteHeader(status int) {
    w.status = int32(status)
    w.ResponseWriter.WriteHeader(status)
}

func (w *statusWrapper) Write(b []byte) (int, error) {
    if w.status == 0 {
        w.status = 200
    }
    n, err := w.ResponseWriter.Write(b)
    w.length += int64(n)
    return n, err
}

func getIp(r *http.Request) string {
    var ip string
    if len(r.Header["X-Real-Ip"]) > 0 {
        ip = r.Header["X-Real-Ip"][0]
    } else if len(r.Header["X-Forwarded-For"]) > 0{
        ip = r.Header["X-Forwarded-For"][0]
    }
    return ip
}

// Send a push notification with activity details
func MonitorNotifyWrapper(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if !strings.HasPrefix(r.URL.Path, "/amy") && !strings.HasPrefix(r.URL.Path, "/media") {
            moniClient := pb.NewMonitorService("go.micro.api.monitor", client.DefaultClient)
            ip := getIp(r)
            if len(r.Header["X-Real-Ip"]) > 0 {
                ip = r.Header["X-Real-Ip"][0]
            } else if len(r.Header["X-Forwarded-For"]) > 0{
                ip = r.Header["X-Forwarded-For"][0]
            }

            var host string
            if len(r.Header["X-Host"]) > 0 {
                host = r.Header["X-Host"][0]
            }


            // Create context for RPC request
            md := make(map[string]string)
            if len(r.Header["Request-Id"]) > 0 {
                md["Request-Id"] = r.Header["Request-Id"][0]
            }
            ctx := metadata.NewContext(context.Background(), md)

            req := pb.Activity{
                Ip: ip,
                Route: r.URL.Path,
                Host: host,
            }
            go moniClient.NotifyActivity(ctx, &req)
        }
        next.ServeHTTP(w, r)
    })
}

func HTTPTraceWrapper(serviceName string) (func(http.Handler) http.Handler) {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            moniClient := pb.NewMonitorService("go.micro.api.monitor", client.DefaultClient)
            requestId, err := moniClient.CreateRequest(context.Background(), &pb.CreateRequestMessage{
                Route: r.URL.Path,
                Ip: getIp(r),
            })
            if err != nil {
                log.Println("Error creating request: ", err.Error(), "... serviceName: ", serviceName)
            } else {
                // Overwrite existing Request-Id headers
                log.Println("Received request id ", requestId.RequestId)
                r.Header["Request-Id"] = []string{requestId.RequestId}
            }

            wrappedWriter := &statusWrapper{ResponseWriter: w}
            next.ServeHTTP(wrappedWriter, r)
            if requestId != nil {
                moniClient.FinishRequest(context.Background(), &pb.FinishRequestMessage{
                    RequestId: requestId.RequestId,
                    StatusCode: wrappedWriter.status,
                    ResponseSize: wrappedWriter.length,
                })
            }
        })
    }
}

func RPCTraceWrapper(serviceName string) (func(server.HandlerFunc) server.HandlerFunc) {
    return func(fn server.HandlerFunc) server.HandlerFunc {
        return func(ctx context.Context, req server.Request, resp interface{}) error {
            moniClient := pb.NewMonitorService("go.micro.api.monitor", client.DefaultClient)
            md, ok := metadata.FromContext(ctx)
            if ok {
                pbTrace := pb.Trace{
                    RequestId: md["Request-Id"],
                    ParentId: md["Trace-Id"],   // The new trace's parent will be the current trace
                    Method: md["Micro-Method"],
                    Service: serviceName,
                }
                traceId, err := moniClient.StartTrace(ctx, &pbTrace)
                if err == nil {
                    defer moniClient.FinishTrace(ctx, traceId)

                    // Bump the current trace to the parent and set the current trace to the new trace
                    md["Parent-Trace-Id"] = md["Trace-Id"]
                    md["Trace-Id"] = traceId.TraceId
                    ctx = metadata.NewContext(ctx, md)
                } else {
                    log.Println("Error getting trace id: ", err.Error())
                }
            } else {
                log.Println("Failed to get metadata!!!")
            }
            return fn(ctx, req, resp)
        }
    }
}
