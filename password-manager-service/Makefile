build:
	protoc -I. --micro_out=. --go_out=plugins=micro:. proto/password-manager/password-manager.proto
	docker build -t jonb377/website:password-manager .

push:
	docker push jonb377/website:password-manager
