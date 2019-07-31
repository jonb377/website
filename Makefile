build:
	docker build -t website:latest .
	docker-compose build

run:
	docker-compose up
