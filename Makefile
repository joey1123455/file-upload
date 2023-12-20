build:
	docker-compose build 

dev:
	docker-compose up 

prod:
	docker-compose up -d

down:
	docker-compose down

logs:
	docker-compose logs

test:
	docker-compose run --rm app go test -v ./...
