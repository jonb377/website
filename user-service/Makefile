build:
	protoc -I. --go_out=plugins=micro:. proto/user/user.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t user-service:latest .
	rm user-service

run:
	docker run -p 9546:9546 user-service:latest
