FROM golang:1.15.1-alpine3.12

WORKDIR /restapi

ENV GO111MODULE on

RUN set -ox pipefail \
  && apk update \
  && apk add --no-cache bash vim \
  mariadb-client mariadb-dev \
  && rm -rf /var/cache/apk/* \
  && go get -u gorm.io/gorm \
  && go get -u gorm.io/driver/mysql \
  && go get -u github.com/gorilla/mux

COPY . .

# Todo: マルチステージビルドは一旦お預け