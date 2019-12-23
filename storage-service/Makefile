build:
	protoc -I. --micro_out=. --go_out=plugins=micro:. proto/storage/storage.proto
	docker build -t jonb377/website:storage .

push:
	docker push jonb377/website:storage
