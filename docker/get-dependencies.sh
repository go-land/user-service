#!/bin/bash
echo "Downloading all dependencies for user-service"

# go get -v ./...
#go get google.golang.org/grpc

go get -u github.com/micro/go-micro
go get -u github.com/go-land/job-service/proto