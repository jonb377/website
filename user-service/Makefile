build:
	protoc -I. --micro_out=. --go_out=plugins=micro:. proto/user/user.proto
	docker build -t jonb377/website:user .

push:
	docker push jonb377/website:user
