.PHONY: build
build:
	@docker-compose up --build

.PHONY: run
run:
	docker-compose up -d

.PHONY: stop
stop:
	docker-compose down

.PHONY: clean
clean:
    docker rmi go_api,postgres_db