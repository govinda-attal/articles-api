.PHONY: init install test build serve clean pack deploy ship run

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
serve: build
	./dist/$(APP_NAME) start

clean:
	rm ./dist/ -rf

pack:
	docker build --build-arg APP_NAME=$(APP_NAME) -t gattal/$(APP_NAME):$(TAG) .

upload:
	docker push gattal/$(APP_NAME):$(TAG)	

run:
	docker run --name articles-api -d -p $(HOST_PORT):9080 gattal/$(APP_NAME):$(TAG)

ship: init test pack upload clean