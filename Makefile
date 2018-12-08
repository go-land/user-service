########################################################################################################################
# Constants
########################################################################################################################

SERVICE_NAME=user-service

# Get the version number from VERSION file.
CODE_VERSION = $(strip $(shell cat VERSION))

ifndef CODE_VERSION
$(error You need to create a VERSION file to build a release)
endif

# Get the latest commit.
GIT_COMMIT = $(strip $(shell git rev-parse --short HEAD))

# Image and binary can be overidden with env vars.
DOCKER_IMAGE ?= go-land/$(SERVICE_NAME)
DOCKER_TAG = $(CODE_VERSION)

########################################################################################################################
# Commands
########################################################################################################################

.PHONY: default build proto install deploy docker_build docker_push run output

default: build

proto:
	protoc --proto_path=${GOPATH}/src:. --go_out=plugins=grpc:. proto/*.proto

install:
	go get -t -v ./...

# Build docker image
build: proto docker_build output

# Build and push Docker image
deploy: proto docker_build docker_push output

docker_build:
	@docker build \
		--build-arg VCS_REF=$(GIT_COMMIT) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		--build-arg VERSION=$(CODE_VERSION) \
		--build-arg SERVICE_NAME=$(SERVICE_NAME) \
		-f docker/Dockerfile -t $(DOCKER_IMAGE):$(DOCKER_TAG) ../

docker_push:
	# Tag image as latest
	docker tag $(DOCKER_IMAGE):$(DOCKER_TAG) $(DOCKER_IMAGE):latest

	# Push to DockerHub
	docker push $(DOCKER_IMAGE):$(DOCKER_TAG)
	docker push $(DOCKER_IMAGE):latest

run:
	docker-compose -f docker-compose.yml -p dev_go up

output:
	@echo Docker Image: $(DOCKER_IMAGE):$(DOCKER_TAG)

