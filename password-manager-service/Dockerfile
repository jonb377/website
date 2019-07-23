FROM golang:1.12.0 as builder

WORKDIR /go/src/github.com/jonb377/website/password-manager-service

COPY . .

ENV GOPATH=/go

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/jonb377/website/password-manager-service/password-manager-service .

ENV MICRO_SERVER_ADDRESS=:9546

CMD ["./password-manager-service"]

