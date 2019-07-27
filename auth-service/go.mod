module github.com/jonb377/website/auth-service

go 1.12

replace (
	github.com/jonb377/website/auth-service => ./
	github.com/nats-io/nats.go => github.com/nats-io/nats.go v1.8.1
)

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.1.1
	github.com/jinzhu/gorm v1.9.10
	github.com/micro/go-micro v1.7.1-0.20190627135301-d8e998ad85fe
	github.com/micro/go-plugins v1.1.1
)
