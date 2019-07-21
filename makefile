build:
	docker-compose build

up: build
	docker-compose up

dev:
	docker-compose up postgres

willson: build
	docker-compose up wclient