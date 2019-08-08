build:
	protoc -I. --micro_out=. --go_out=plugins=micro:. proto/password-manager/password-manager.proto
	docker build -t jonb377/website:password-manager .

run:
	docker run -p 9546:9546 password-manager-service:latest
