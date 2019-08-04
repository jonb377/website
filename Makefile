build:
	docker build -t website:latest .

run:
	docker-compose build
	docker-compose up --force-recreate
