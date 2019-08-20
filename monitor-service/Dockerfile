FROM jonb377/website:builder as builder

WORKDIR /go/src/github.com/jonb377/website

RUN rm -r monitor-service

COPY . monitor-service

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o monitor monitor-service/main.go

FROM alpine:latest

RUN mkdir /app
WORKDIR /app
COPY config/ipstack_access_key.txt /var/srv/
COPY --from=builder /go/src/github.com/jonb377/website/monitor .

ENV MICRO_SERVER_ADDRESS=:9546

CMD ["./monitor"]
