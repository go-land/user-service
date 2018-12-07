.PHONY: proto data build

proto:
	protoc --proto_path=${GOPATH}/src:. --go_out=plugins=grpc:. proto/*.proto

build:
	docker build -f docker/Dockerfile -t "go-land/user-service:1.0.0" ../

run:
	docker-compose -f docker-compose.yml -p dev_go up

docker-clean:
	docker stop $(docker ps -a | grep dev_go_* | awk '{print $1}')
	docker rm $(docker ps -a | grep dev_go_* | awk '{print $1}')

