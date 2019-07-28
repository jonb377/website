module github.com/jonb377/website/user-service

go 1.12

replace (
	github.com/jonb377/website/user-service => ./
	github.com/nats-io/nats.go => github.com/nats-io/nats.go v1.8.1
)

require (
	cloud.google.com/go v0.43.0 // indirect
	github.com/golang/protobuf v1.3.2
	github.com/google/pprof v0.0.0-20190723021845-34ac40c74b70 // indirect
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/jinzhu/gorm v1.9.10
	github.com/micro/go-micro v1.7.1-0.20190627135301-d8e998ad85fe
	github.com/micro/go-plugins v1.1.1
	github.com/micro/protoc-gen-micro v0.8.0 // indirect
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4 // indirect
	golang.org/x/exp v0.0.0-20190718202018-cfdd5522f6f6 // indirect
	golang.org/x/image v0.0.0-20190703141733-d6a02ce849c9 // indirect
	golang.org/x/mobile v0.0.0-20190719004257-d2bd2a29d028 // indirect
	golang.org/x/net v0.0.0-20190724013045-ca1201d0de80 // indirect
	golang.org/x/sys v0.0.0-20190726091711-fc99dfbffb4e // indirect
	golang.org/x/tools v0.0.0-20190728063539-fc6e2057e7f6 // indirect
	google.golang.org/grpc v1.22.1 // indirect
)
