#!/bin/sh -eux

cd `dirname $0`

# gcloud components install --quiet app-engine-go

go get -u github.com/golang/dep/cmd/dep
cd src/
dep ensure
cd -

# Commands for pre-building
rm -rf build-cmd/
mkdir build-cmd/
go build -o build-cmd/goimports ./src/vendor/golang.org/x/tools/cmd/goimports
go build -o build-cmd/golint    ./src/vendor/github.com/golang/lint/golint
go build -o build-cmd/jwg       ./src/vendor/github.com/favclip/jwg
