# Test API REST
Autor: I.S.C. Edgar Alfred Rodriguez Robles
E-mail: alfred.7790@gmail.com

# Table of Contents
[_TOC_]

# Overview
This project is an example of API REST.

# Requirements
To run this project it is necessary to have installed:
- docker
- docker-compose
- go

# Quick start - Using Makefile
> IMPORTANT! If you have a linux distribution, you will be able to perform this procedure, otherwise I suggest you try another way to run the service.
1. Clone the repo:
```shell
$ git clone git@bitbucket.org:alfred7790/test.git
```
2. Open the project
```shell
$ cd test
```
3. Get dependencies:
```shell
$ make dep
```
4. Build and run the service:
```shell
$ make
```
5. If every thing is ok, you should see something like this:
```shell
...
[GIN-debug] Listening and serving HTTP on :8080
connected to 'test' database
```
6. Go to [swagger docs](http:localhost:8080/v1/swagger/index.html) and have fun.

# Quick start - Running Binary
1. Clone the repo:
```shell
$ git clone git@bitbucket.org:alfred7790/test.git
```
2. Open the project
```shell
$ cd test
```
3. Running the DB service:
```shell
$ docker-compose up -d psql
```
4. Get dependencies:
```shell
$ go mod tidy && go get -u github.com/swaggo/swag/cmd/swag
```
5. Build the service:
```shell
$ go build main.go
```
6. Running the service:
```shell
$ ./main
```
7. If every thing is ok, you should see something like this:
```shell
...
[GIN-debug] Listening and serving HTTP on :8080
connected to 'test' database
```
8. Go to [swagger docs](http:localhost:8080/v1/swagger/index.html) and have fun.

# Custom Config
> If you need to change the default values of the configuration.
1. Open the project.
```shell
$ cd test
```
2. Copy the template of the configuration to `config.yml`.
```shell
$ cp ./app/config/config.template.yml config.yml
```
3. Edit it with your own values.
> WARNING! Make sure that if you edit the values about the `DB service`, also you should modify the `docker-compose.yml` file.
4. Restart the service `using Makefile` or `running Binary`.
5. Go to [swagger docs](http:localhost:8080/v1/swagger/index.html) and have fun.

# Unit Testing
1. Open the project.
```shell
$ cd test
```
2. Run test
```shell
$ go test ./...
```
3. You can use the `Makefile`.
```shell
$ make test
```
# Up and Down DB service
- Using `Makefile`.
```shell
  $ make db
```
```shell
  $ make db-down
```
- Using `docker-compose`
```shell
$ docker-compose up -d psl
```
```shell
  $ docker-compose down
```

# Dependencies
If you get an error with swagger or with another package, try this:
```shell
$ make dep
```
Or
```shell
$ go mod tidy && go get -u github.com/swaggo/swag/cmd/swag
```

# Comments:
- The Database is not persistent.
- I didn't cover the whole project with tests.
- I include default values in the config structure to run the service more easily,
  but it is not good practice including database credentials in public files, 
  these values should be placed in an .env file.
- I didn't include comments about functions because I didn't have enough time.