FROM golang
MAINTAINER Maksym Stepanenko <stepanenkomaksi@gmail.com>

# Install dependencies
RUN go get github.com/gorilla/mux

# Install service application
ADD user-service /go/src/github.com/go-land/user-service
WORKDIR /go/src/github.com/go-land/user-service
RUN go build -o /go/bin/user-service .

ADD user-service/docker/run.sh run.sh
ADD user-service/docker/loop.sh loop.sh
RUN chmod +x run.sh loop.sh

CMD ./run.sh
#CMD ./loop.sh