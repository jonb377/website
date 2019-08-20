build:
	protoc -I. --micro_out=. --go_out=plugins=micro:. proto/monitor/monitor.proto
	docker build -t jonb377/website:monitor .

push:
	docker push jonb377/website:monitor
