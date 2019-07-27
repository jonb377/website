FROM golang:1.12.0 as builder

WORKDIR /go/src/github.com/jonb377/website/user-service

ENV GOPATH=/go

COPY go.mod .

RUN GO111MODULE=on go mod download

COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/jonb377/website/user-service/user-service .

ENV MICRO_SERVER_ADDRESS=:9546

CMD ["./user-service"]
