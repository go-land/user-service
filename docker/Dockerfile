########################################################################################################################
# Build user-service
########################################################################################################################
FROM golang as go-builder

ARG SERVICE_NAME

ADD user-service /go/src/github.com/go-land/$SERVICE_NAME
WORKDIR /go/src/github.com/go-land/$SERVICE_NAME

# Install all dependencies
ADD $SERVICE_NAME/docker/get-dependencies.sh get-dependencies.sh
RUN chmod +x get-dependencies.sh
RUN bash get-dependencies.sh

#RUN go get -u github.com/golang/dep/cmd/dep
#RUN dep init && dep ensure

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o $SERVICE_NAME .

########################################################################################################################
# install user-service
########################################################################################################################
FROM alpine:latest
MAINTAINER Maksym Stepanenko <stepanenkomaksi@gmail.com>

# Security related package, good to have.
RUN apk --no-cache add ca-certificates
RUN apk add bash curl

ARG SERVICE_NAME
ARG VERSION
ARG VCS_REF
ARG BUILD_DATE
ARG SERVICE_NAMESPACE

LABEL com.max.service=$SERVICE_NAME
LABEL com.max.version=$VERSION
LABEL com.max.vcs-ref=$VCS_REF
LABEL com.max.build-date=$BUILD_DATE
LABEL com.max.namespace=$SERVICE_NAMESPACE

RUN mkdir -p /services/$SERVICE_NAME
WORKDIR /services/$SERVICE_NAME

COPY --from=go-builder /go/src/github.com/go-land/$SERVICE_NAME/$SERVICE_NAME .
ADD $SERVICE_NAME/docker/run.sh run.sh
ADD $SERVICE_NAME/docker/loop.sh loop.sh
RUN chmod +x run.sh loop.sh

CMD ./run.sh
