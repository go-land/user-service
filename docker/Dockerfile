FROM golang
MAINTAINER Maksym Stepanenko <stepanenkomaksi@gmail.com>
ADD . /go/src/github.com/algorithms
WORKDIR /go/src/github.com/algorithms

# Install dependencies
RUN go get github.com/gorilla/mux

# Install service application
RUN go install github.com/algorithms

ADD run.sh run.sh
ADD loop.sh loop.sh
RUN chmod +x run.sh loop.sh

#CMD ./run.sh
CMD ./loop.sh

#ENTRYPOINT /go/bin/algorithms
#EXPOSE 7070