# Base builder for each service
FROM golang:1.12.0

WORKDIR /go/src/github.com/jonb377/website

ENV GOPATH=/go
ENV GO111MODULE=on

COPY go.mod .

RUN go mod download

COPY . .
