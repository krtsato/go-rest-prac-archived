FROM mysql:8.0.21

RUN set -x \
  && apt-get update \
  && apt-get install -y --no-install-recommends \
  apt-utils bash vim \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/