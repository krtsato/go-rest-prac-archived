FROM golang:1.15.1-alpine3.12

WORKDIR /restapi

COPY ./go.mod ./

RUN set -ox pipefail \
  && apk update \
  && apk add --no-cache build-base mariadb-client mariadb-dev \
  && rm -rf /var/cache/apk/*

