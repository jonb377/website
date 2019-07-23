build:
	protoc -I. --micro_out=. --go_out=plugins=micro:. proto/password-manager/password-manager.proto
#	GOOS=linux GOARCH=amd64 go build
	docker build -t password-manager-service:latest .
	#rm password-manager-service

run:
	docker run -p 9546:9546 password-manager-service:latest
