#!/usr/bin/env bash

sudo apt-get install -y build-essential
go get -u github.com/golang/dep/cmd/dep
dep ensure

export GOOS=darwin
export GOARCH=amd64
go build -o albionmarket-client -v -x cmd/albionmarket-client/albionmarket-client.go
