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
* [Postgres Array Functions](https://postgres.cz/wiki/Array_based_functions)
* [Database Migrations in Golang](https://github.com/golang-migrate/migrate)
* [Gorilla Mux](https://github.com/gorilla/mux)
* [Negroni Middleware](https://github.com/urfave/negroni)
* [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
* [Technical Test API Specification](https://ffxblue.github.io/interview-tests/test/article-api/)

## Assumptions & Description of the Solution

1. API Spec mandated a Persistence layer. Of three API functions, getting summary for a given tag and date was a bit tricky. It made me think that possibly it will be best to use a DB that will easily support Arrays. So to realize this aspect of Spec brought in Postgres. It has support for Arrays and GIN for Array index implementation. 
    - [Using Postgres Arrays with Golang](https://www.opsdash.com/blog/postgres-arrays-golang.html)
    - [Postgres Array Functions](https://postgres.cz/wiki/Array_based_functions)

2. Database access can be achieved in multiple ways and to in my opinion at high level these can be classified into ORM style or low level style with database/sql. Personally prefer much more powerful style of using go native database/sql.

3. With Database required, it was important to have a simple yet powerful method to DB migration frameworks. These framework help to setup database so that same can used effectively by application service. So chose a simple framework golang-migrate/migrate.
    - [Database Migrations in Golang](https://github.com/golang-migrate/migrate)

4. The code test clearly laid out importance to organisation of codebase. There a few school of thoughts on how go code base is setup. To my knowledge, one as adivsed by Ben Johnson on [Standard Package Layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1) and a little more enhanced practice followed as per [golang-standards/project-layout](https://github.com/golang-standards/project-layout). Source code layout of Kubernetes, Istio, Hyperledger, etc have their own preferred code organisation philosophy that is very similar to [golang-standards/project-layout](https://github.com/golang-standards/project-layout). So had to pick up a hybird strategy and somewhat similar strategy we follow at our current workplace.

5. Microservice should implement certain route specific validations like only integers are allowed to find a given article ``/articles/{id:[0-9]+}``. So it is wise to use battle tested Routing framework like [Gorilla Mux](https://github.com/gorilla/mux). The simplicity of handlers that can be used along with gorilla mux follow native net/http handler func specification.

6. Microservice should allow easy reportation on NFRs like response times. It is easy to use [Negroni Middleware](https://github.com/urfave/negroni) when this functionality is provided out of the box. Of course Negroni middleware framework makes it easy to plugin custom middlewares for custom requirements too.

7. Microservice implementation must have unit tests. As this micro-service being an application service with different behaviours and say more to be added in future, it nice to follow BDD style testing. Somtimes it is difficult to follow GoDoc style documentation for application services. But BDD will naturally help to document behaviours of the application micro service. So use of [Ginkgo - A Golang BDD Testing Framework](https://onsi.github.io/ginkgo/)

8. Simple style of Mocking Dependencies when making code unit testable [Dependecy Injection, Mocking & TDD/BDD in Go](https://www.youtube.com/watch?v=uFXfTXSSt4I). Many Go purists will shy away from BDD and this style of Mocking Dependencies, but it has helped me and our current team.

9. Comments, Comments, Comments. Go Code must be well commented as per standard guidelines - specially for exported Methods, Attributes, Types, Packages.

10. Need to have custom but simple strategy for error handling. So error structure will include code and message when returned by API. Intentionally error code are mapped to relevant HTTP Status codes for simplicity and easy comprehension, but there is also an option to have more application specific code.

11. Dockerize API : hence Docker File

12. API Configuration (like DB endpoint information) managed externally within configuration file. Hence use [Viper - Configuration with Fangs](https://github.com/spf13/viper)

13. Microservice to follow CLI style dictated by popular framework like [Cobra](https://github.com/spf13/cobra). To be nice, microservice is a cli with subcommands
    - start - starts the server to listen to HTTP requests
    - migrate up|down - to manage provider database migrations

