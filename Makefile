build:
	docker build -t jonb377/website:builder .
	cd auth-service && make build
	cd user-service && make build
	cd password-manager-service && make build
	cd router-service && make build

run:
	docker-compose build
	docker-compose up --force-recreate

push:
	docker login
	#docker push jonb377/website:builder
	docker push jonb377/website:auth
	docker push jonb377/website:user
	docker push jonb377/website:password-manager
	docker push jonb377/website:router
