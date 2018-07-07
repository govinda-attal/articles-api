.PHONY: init install test build db-local db-migrate-up db-migrate-down serve clean pack deploy ship run

include .env
export $(shell sed 's/=.*//' .env)

init:
	go get -u github.com/golang/dep/cmd/dep
	go get github.com/onsi/ginkgo/ginkgo
	go get github.com/onsi/gomega/...

install: init
	rm -rf ./vendor
	dep ensure

test: init
	ginkgo -r

build: 
	rm -rf ./dist
	mkdir dist
	mkdir dist/config
	GOOS=linux GOARCH=amd64 go build -o ./dist/$(APP_NAME) .
	cp ./test/fixtures/app-config-local.yaml ./dist/config/app-config.yaml

db-local:
	docker run --name $(APP_NAME)-db -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -p 5432:5432 -d postgres
	sleep 5s && 

db-migrate-up: build
	cd dist && cp -r ../migrations . && ./$(APP_NAME) migrate up

db-migrate-down: build
	cd dist && cp -r ../migrations . && ./$(APP_NAME) migrate down

serve: build
	cd dist && ./$(APP_NAME) start

clean:
	rm ./dist/ -rf

pack:
	docker build --build-arg APP_NAME=$(APP_NAME) -t gattal/$(APP_NAME):$(TAG) .

upload:
	docker push gattal/$(APP_NAME):$(TAG)	

run:
	docker run --name articles-api -d -p $(HOST_PORT):9080 gattal/$(APP_NAME):$(TAG)

ship: init test pack upload clean