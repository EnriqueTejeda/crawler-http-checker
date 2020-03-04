# Helpers
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=crawler-http-checker
DOCKER_NAME=etejeda/crawler-http-checker:latest

.PHONY: build test clean build-docker
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
run: 
	./$(BINARY_NAME)
test:
	$(GOTEST) -v
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
build-docker:
	docker build --compress . -t ${DOCKER_NAME}
