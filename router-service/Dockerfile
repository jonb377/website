FROM jonb377/website:builder as builder

WORKDIR /go/src/github.com/jonb377/website

RUN rm -r router-service

COPY . router-service

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o router router-service/main.go

FROM alpine:latest

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/jonb377/website/router .

CMD ["./router"]
