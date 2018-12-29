ORG := jeremyhahn
PACKAGE := testapp
TARGET_OS := linux

ARCH := $(shell go env GOARCH)
OS := $(shell go env GOOS)
APPNAME=testapp

export PATH := $(GOPATH)/bin:$(PATH)

.PHONY: build clean scp

default: build

build:
	go build -o $(APPNAME)

clean:
	rm $(APPNAME)

run:
	./$(APPNAME)
