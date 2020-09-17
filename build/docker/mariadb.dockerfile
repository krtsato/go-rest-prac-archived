FROM mariadb:10.5.5-focal

RUN set -ex \
&& apt-get update \
&& apt-get install -y --no-install-recommends tzdata \
&& rm /etc/localtime \
&& ln -s /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
&& apt-get --purge -y remove tzdata \
&& apt-get clean \
&& rm -rf /var/libi/apt/*
