#!/bin/sh -eux

cd `dirname $0`

packages=`go list ./src/api/...`

# Apply tools
export PATH=$(pwd)/build-cmd:$PATH
which goimports golint jwg
goimports -w ./src/api/
go tool vet ./src/api/
golint $packages
go generate $packages

goapp test $packages -p 1 $@
