include .env

.PHONY: build
build:
	docker-compose build --no-cache

.PHONY: run
run:
	docker-compose up

.PHONY: stop
stop:
	docker-compose down

.PHONY: psql
psql: 
	PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -d $(DB_NAME) -x