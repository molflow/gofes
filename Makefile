GOPATH ?= $(HOME)/go
GOBIN ?= $(GOPATH)/bin
GOCMD := go
GOBUILD := $(GOCMD) build
GOINSTALL := $(GOCMD) install
GOVET := $(GOCMD) vet
GOLINT := $(GOBIN)/golint
GOTEST := $(GOCMD) test
all: build lint test
build:
	$(GOBUILD) ./...
install:
	$(GOINSTALL) -i ./...
lint:
	$(GOVET) ./...
	$(GOLINT) -set_exit_status ./...
test:
	$(GOTEST) -race ./...
