build:
	docker build -t jonb377/website:builder .
	cd auth-service && make build
	cd user-service && make build
	cd password-manager-service && make build
	cd router-service && make build
	cd web-service && make build
	cd cron && make build
	cd notifications-service && make build

run:
	docker-compose build
	docker-compose up --force-recreate

push:
	docker login
	docker push jonb377/website:auth
	docker push jonb377/website:user
	docker push jonb377/website:password-manager
	docker push jonb377/website:router
	docker push jonb377/website:cron
	docker push jonb377/website:web
	docker push jonb377/website:notifications
