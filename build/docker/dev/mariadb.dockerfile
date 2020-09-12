FROM mariadb:10.5.5-focal

RUN set -ex \
  && apt-get update \
  && apt-get install -y --no-install-recommends bash vim \
  && apt-get clean \
  && rm -rf /var/lib/apt/*
