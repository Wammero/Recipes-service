all: docker-start-service-debug

docker-start-service-debug: docker-compose-up

docker-compose-up: docker-compose-build
	cd docker && docker-compose up

docker-compose-build:
	cd docker && docker-compose build

docker-clean-containers:
	cd docker && docker-compose down

docker-clean-data:
	cd docker && docker-compose down --rmi all
