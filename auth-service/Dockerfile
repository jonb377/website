FROM jonb377/website:builder as builder

WORKDIR /go/src/github.com/jonb377/website

RUN rm -r auth-service

COPY . auth-service

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o auth auth-service/main.go

FROM alpine:latest

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/jonb377/website/auth .

ENV MICRO_SERVER_ADDRESS=:9546

CMD ["./auth"]
