build:
	protoc -I. --micro_out=. --go_out=plugins=micro:. proto/password-manager/password-manager.proto
	docker build -t website:latest ..
	docker build -t password-manager-service:latest .

run:
	docker run -p 9546:9546 password-manager-service:latest
