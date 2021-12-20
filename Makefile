PKG_LIST = `go list ./... | grep -v /vendor`
.PHONY: all
all: | db dep test build deploy

.PHONY: dep
dep:
	@echo "Download dependencies"
	@echo "GOPATH is:" $(GOPATH)
	go mod tidy
	go get -u github.com/swaggo/swag/cmd/swag

.PHONY: db
db:
	@echo "Running postgresdb"
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
	@echo "Building binary"
	go build main.go

.PHONY: generate
generate:
	@echo "Generating swagger docs"
	$(GOPATH)/bin/swag init

.PHONY: test
test:
	@echo "Running tests"
	go test $(PKG_LIST)

.PHONY: deploy
deploy:
	@echo "Deploying main"
	./main

