build:
	protoc -I. --micro_out=. --go_out=plugins=micro:. proto/auth/auth.proto
	docker build -t jonb377/website:auth .

run:
	docker run -p 9546:9546 auth-service:latest
