GOPATH ?= $(HOME)/go
GOBIN ?= $(GOPATH)/bin
GOCMD := go
GOBUILD := $(GOCMD) build
GOVET := ${GOCMD} vet
GOLINT := $(GOBIN)/golint
GOTEST := $(GOCMD) test
all: build lint test
build:
	$(GOBUILD) ./...
lint:
	${GOVET} ./...
	${GOLINT} -set_exit_status ./...
test:
	$(GOTEST) -race ./...
