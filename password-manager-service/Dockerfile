FROM website:latest as builder

WORKDIR /go/src/github.com/jonb377/website

RUN rm -r password-manager-service

COPY . password-manager-service

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o password-manager password-manager-service/main.go

FROM alpine:latest

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/jonb377/website/password-manager .

ENV MICRO_SERVER_ADDRESS=:9546

CMD ["./password-manager"]

