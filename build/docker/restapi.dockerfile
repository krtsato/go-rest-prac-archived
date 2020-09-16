FROM golang:1.15.1-alpine3.12 AS build

WORKDIR /go-rest-sample

COPY ./ ./

ENV CGO_ENABLED=0

RUN set -ox pipefail \
  && apk update \
  && apk add --no-cache \
  build-base bash vim curl git \
  mariadb-client mariadb-dev \
  && rm -rf /var/cache/apk/* \
  && curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s \
  && go build -v -o /build/tmp/restapi ./cmd/restapi/main.go

FROM alpine:3.12.0

WORKDIR /go-rest-sample

COPY --from=build /bin/restapi ./restapi

CMD ["/go-rest-sample/restapi"]