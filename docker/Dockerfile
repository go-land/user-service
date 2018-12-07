########################################################################################################################
# Build user-service
########################################################################################################################
FROM golang as go-builder

###################################
# Get hello-service dependencies
###################################
RUN go get -u google.golang.org/grpc

ADD user-service /go/src/github.com/go-land/user-service
WORKDIR /go/src/github.com/go-land/user-service
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o user-service .

########################################################################################################################
# install hello-service
########################################################################################################################
FROM alpine:latest
RUN apk --no-cache add ca-certificates bash curl
RUN mkdir -p /go
WORKDIR /go
COPY --from=go-builder /go/src/github.com/go-land/user-service .
ADD user-service/docker/run.sh run.sh
ADD user-service/docker/loop.sh loop.sh
RUN chmod +x run.sh loop.sh
CMD ./run.sh