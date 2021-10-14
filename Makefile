GOCMD:=$(shell which go)
GOLINT:=$(shell which golint)
GOIMPORT:=$(shell which goimports)
GOFMT:=$(shell which gofmt)
GOBUILD:=$(GOCMD) build
GOINSTALL:=$(GOCMD) install
GOCLEAN:=$(GOCMD) clean
GOTEST:=$(GOCMD) test
GOGET:=$(GOCMD) get
GOLIST:=$(GOCMD) list
GOVET:=$(GOCMD) vet
GOPATH:=$(shell $(GOCMD) env GOPATH)
u := $(if $(update),-u)

BINARY_NAME:=swag
GOFILES:=$(shell find . -name "*.go" -type f)

build:
	go build -o bin/main main.go

run:
	make docs
	echo "Running API"
	PROFILE=DEV go run main.go

docs: deps
	echo "Generating API docs"
	swag init
	echo "Generated API docs"

tests:
	echo "Running tests"
	go test ./api/v1/store

deps:
	go get -u github.com/swaggo/swag/cmd/swag
	go get -u github.com/swaggo/echo-swagger
