module github.com/jonb377/website/password-manager-service

go 1.12

replace (
	github.com/jonb377/website/password-manager-service => ./
	github.com/nats-io/nats.go => github.com/nats-io/nats.go v1.8.1
)

require (
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/gorm v1.9.10
	github.com/micro/go-micro v1.7.1-0.20190627135301-d8e998ad85fe
	github.com/micro/go-plugins v1.1.1
)
