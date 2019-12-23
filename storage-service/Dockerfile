FROM jonb377/website:builder as builder

WORKDIR /go/src/github.com/jonb377/website

RUN rm -r storage-service

COPY . storage-service

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o storage storage-service/main.go

FROM alpine:latest

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/jonb377/website/storage .

ENV MICRO_SERVER_ADDRESS=:9546

CMD ["./storage"]

