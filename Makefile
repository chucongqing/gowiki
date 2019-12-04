.PHONY: build run test
GOCMD=go
BIN_PATH=./bin/
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run
BINARY_NAME=mybinary
BINARY_UNIX=$(BINARY_NAME)_unix


default: run


build: 
	$(GOBUILD) -o $(BIN_PATH)$(BINARY_UNIX) 
run:
	$(GORUN) wiki.go
test:
	$(GOTEST) ./...
