FROM golang:1.15.1-alpine3.12

WORKDIR /restapi

COPY ./go.mod ./

RUN set -ox pipefail \
  && apk update \
  && apk add --no-cache bash git \
  mariadb-client mariadb-dev \
  && rm -rf /var/cache/apk/*

# Todo: マルチステージビルド