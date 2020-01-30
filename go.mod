module github.com/jonb377/website

go 1.12

replace (
	github.com/jonb377/website => ./
	github.com/nats-io/nats.go => github.com/nats-io/nats.go v1.8.1
)

require (
	firebase.google.com/go v3.11.0+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.1.1
	github.com/gorilla/mux v1.7.2
	github.com/jinzhu/gorm v1.9.10
	github.com/micro/go-micro v1.7.1-0.20190627135301-d8e998ad85fe
	github.com/micro/go-plugins v1.1.1
	github.com/micro/micro v1.7.1-0.20190627121529-410a2eba67f1
)
