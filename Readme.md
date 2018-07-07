# Articles API

A simple API with three endpoints.

The first endpoint, POST /articles handles receipt of article data in json format, and stores it within the service.

The second endpoint GET /articles/{id} returns the JSON representation of the article.

The final endpoint, GET /tags/{tagName}/{date} returns the list of articles that have that tag name on the given date and some summary data about that tag for that day.

The source code implements this microservice (A Go code test).

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

1. Linux OS - Ubuntu or Fedora is preferred.
2. git
3. Golang Setup locally on OS
4. Postman for REST API testing
5. Docker & Docker Compose

### Installing

This application can be setup locally in following ways:

#### Option A
```
go get github.com/govinda-attal/articles-api
```

#### Option B (Preferred) :heavy_check_mark:
```
cd $GOPATH/src/github.com/
mkdir govinda-attal
cd govinda-attal
git clone https://github.com/govinda-attal/articles-api.git
```

### Application Development Setup

'Makefile' will be used to setup articles-api quickly on the development workstation.

```
cd $GOPATH/src/github.com/govinda-attal/articles-api
make install # This will go get 'dep' and use it to install vendor dependencies.
```

## Running Tests

### Unit tests

Sample BDD Unit tests are implemented using ginkgo and gomega. These unit tests are written for HTTP handler only.
Three Unit tests are written for this code test.

```
cd $GOPATH/src/github.com/govinda-attal/articles-api
make test # This will run execute unit tests with ginkgo -r command
```

### Integration tests

This microservice achives given requirements with Golang and Postgres Database as backend. To keep this foot-print of this application minimum postgres db will execute within a docker container. Where as following backend microservice can be hosted within a docker container or local OS.

#### Option A: Docker Compose - orchestrate DB and Microservice as docker containers (Preferred) :heavy_check_mark:

```
cd $GOPATH/src/github.com/govinda-attal/articles-api
docker-compose up -d # This will start Postgres DB, Articles Microservice and Swagger-UI which will point to Articles microservice swagger definition.
```

Docker compose will orchestrate containers and they can be accessed from Local OS as below:
1. Postgres DB on localhost:5432
2. Microservice on :earth_asia: http://localhost:9080
3. Swagger-UI on :earth_asia: http://localhost:8080

#### Option B: Postgres as Docker container but Microservice is run locally on your OS

```
cd $GOPATH/src/github.com/govinda-attal/articles-api
make db-local # This will host Postgres DB as docker container and its port 5432 is mapped to 5432 on Host OS.
make db-migrate-up # This will run DB migrations to setup necessary postgres DB artefacts - Schema, Tables, Functions, Triggers, Indexes
make serve # This will build the Microservice and run it locally and the API is exposed on port 9080 
```

### Postman Test Collection

Start Postman and import sample test collection at
``` 
$GOPATH/src/github.com/govinda-attal/articles-api/test/fixtures/articles-api.postman_collection.json
```

### Swagger UI to view and trial Microservice Open API

Post running command *docker-compose up -d* use browser to open :earth_asia: http://localhost:8080

## Cleanup

For containers orchestrated by Docker-compose
```
cd $GOPATH/src/github.com/govinda-attal/articles-api
docker-compose down
docker image rm gattal/articles-api:latest
docker image rm postgres:latest
```

In case when Postgres DB was running as docker container with Microservice running locally
```
docker stop articles-db
docker rm articles-db
docker image rm postgres:latest 
```
Press Ctrl+C on the terminal on which microservice was running.


To delete the source code for this microservice run :skull: cd $GOPATH/src/github.com/ && rm -rf  govinda-attal :skull: 

## Authors

* [Govinda Attal](https://github.com/govinda-attal)

## Acknowledgments

* [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
* [Organizing Database Access](https://www.alexedwards.net/blog/organising-database-access)
* [Standard Package Layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1)
* [Dependecy Injection, Mocking & TDD/BDD in Go](https://www.youtube.com/watch?v=uFXfTXSSt4I)
* [Example Usage of net/http/httptest](https://golang.org/src/net/http/httptest/example_test.go)
* [Unit testing handlers with Gorilla Mux](https://stackoverflow.com/questions/34435185/unit-testing-for-functions-that-use-gorilla-mux-url-parameters)
* [Testing HTTP Handlers](https://blog.questionable.services/article/testing-http-handlers-go/)
* [Ginkgo - A Golang BDD Testing Framework](https://onsi.github.io/ginkgo/)
* [Using Postgres Arrays with Golang](https://www.opsdash.com/blog/postgres-arrays-golang.html)
* [Database Migrations in Golang](https://github.com/golang-migrate/migrate)
* [Gorilla Mux](https://github.com/gorilla/mux)
* [Negroni Middleware](https://github.com/urfave/negroni)
* [Technical Test API Specification](https://ffxblue.github.io/interview-tests/test/article-api/)
