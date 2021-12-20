PKG_LIST = `go list ./... | grep -v /vendor`

.PHONY: all
all: | db dep test build deploy

.PHONY: dep
dep:
	go mod tidy
	go get -u github.com/swaggo/swag/cmd/swag

.PHONY: db
db:
	docker-compose up -d psql
	docker-compose ps
	# need to make sure our db is up and available before we run tests.
	# 2 seconds
	sleep 2

.PHONY: db-down
db-down:
	docker-compose down

.PHONY: build
build: generate
	go build main.go

.PHONY: generate
generate:
	swag init

.PHONY: test
test:
	go test $(PKG_LIST)

.PHONY: deploy
deploy:
	./main

