########################################################################################################################
# install user-cli
########################################################################################################################
FROM alpine:latest
MAINTAINER Maksym Stepanenko <stepanenkomaksi@gmail.com>
RUN apk --no-cache add ca-certificates bash curl
COPY user-cli .

ADD $cli/run.sh run.sh
RUN chmod +x run.sh

LABEL com.max.service=user-cli
LABEL com.max.namespace=go-land

CMD ./run.sh


