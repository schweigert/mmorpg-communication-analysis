build:
	docker-compose build

up: build
	docker-compose up

dev:
	docker-compose up postgres