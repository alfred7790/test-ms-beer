.PHONY: all
all: | db build deploy

.PHONY: db
db:
	docker-compose up -d psql
	docker-compose ps

.PHONY: db-down
db-down:
	docker-compose down

.PHONY: build
build: generate
	go build main.go

.PHONY: generate
generate:
	swag init

.PHONY: deploy
deploy:
	./main

