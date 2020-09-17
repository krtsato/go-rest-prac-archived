#!/bin/bash

until mysql -h $MYSQL_ROOT_HOST -u $MYSQL_USER -p$MYSQL_PASSWORD -e 'quit'; do
  echo >&2 'Mariadb is sleeping now.'
  sleep 1
done

echo >&2 'Mariadb is up.'
