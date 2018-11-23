FROM golang
MAINTAINER Maksym Stepanenko <stepanenkomaksi@gmail.com>
ADD ../user-service /go/src/github.com/go-land/user-service
WORKDIR /go/src/github.com/go-land/user-service

# Install dependencies
RUN go get github.com/gorilla/mux

# Install service application
RUN go install github.com/go-land/user-service

ADD run.sh run.sh
ADD loop.sh loop.sh
RUN chmod +x run.sh loop.sh

#CMD ./run.sh
CMD ./loop.sh