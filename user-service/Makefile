build:
	protoc -I. --micro_out=. --go_out=plugins=micro:. proto/user/user.proto
	docker build -t user-service:latest .

run:
	docker run -p 9546:9546 user-service:latest
