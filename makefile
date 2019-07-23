build: dependencies
	docker-compose build

up: build
	docker-compose up

dev:
	docker-compose up postgres

willson: build
	docker-compose up wclient


dependencies: wwebservice_dependencies.png wclient_dependencies.png

wclient_dependencies.png:
	godepgraph github.com/schweigert/mmorpg-communication-analysis/clients/wclient | dot -Tpng -o wclient_dependencies.png

wwebservice_dependencies.png:
	godepgraph github.com/schweigert/mmorpg-communication-analysis/infrastructure/services/wwebservice | dot -Tpng -o wwebservice_dependencies.png