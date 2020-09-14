FROM golang:1.15.1-alpine3.12 AS build

WORKDIR /restapi

COPY ./ ./

RUN set -ox pipefail \
  && apk update \
  && apk add --no-cache build-base mariadb-client mariadb-dev \
  && rm -rf /var/cache/apk/* \
  && curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s \
  && go build -o /go/app/restapi ./cmd/restapi/main.go
