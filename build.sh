#!/bin/bash
codedir='/usr/src/codetest'
echo ">>> Compling code"
#docker run --rm -e GOPATH=/usr/ -v "$PWD":"$codedir" -w "$codedir" golang:1.8 sh -c "env && go get -v ./... && GOOS=linux CGO_ENABLED=0 go build -a -ldflags '-s' -tags netgo -ldflags '-w' -v"
docker run --rm -e GOPATH=/usr/ -v "$PWD":"$codedir" -w "$codedir" golang:1.8 sh -c "env && go get -v ./... && GOOS=linux go build -tags netgo -installsuffix netgo -v"
#docker run --rm -e GOPATH=/usr/ -v "$PWD":"$codedir" -w "$codedir" golang:1.8 sh -c "env && go get -v ./... && go build -v"
echo ">>> Building docker container"
docker build -t codetest .
