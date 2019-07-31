build:
	protoc -I. --micro_out=. --go_out=plugins=micro:. proto/auth/auth.proto
	docker build -t website:latest ..
	docker build -t auth-service:latest .

run:
	docker run -p 9546:9546 auth-service:latest
