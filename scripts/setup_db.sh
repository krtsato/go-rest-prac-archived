#!/bin/bash

until mysql -h $MYSQL_HOST -u $MYSQL_USER -p$MYSQL_PASSWORD -e 'quit'; do
  echo >&2 'MySQL is sleeping now.'
  sleep 1
done

echo >&2 'MySQL is up.'

mysql -h $MYSQL_HOST -u root -p$MYSQL_ROOT_PASSWORD -e "grant all on *.* to 'mysql_user'@'%'"

echo 'MYSQL_USER get ALL PRIVILEGES.'

mysql -h $MYSQL_HOST -u root -p$MYSQL_ROOT_PASSWORD -e "create database if not exists restapi_db;"

mysql -h $MYSQL_HOST -u root -p$MYSQL_ROOT_PASSWORD -e "create table if not exists restapi_db.users (
  id int primary key AUTO_INCREMENT,
  first_name varchar(20) NOT NULL, 
  last_name VARCHAR(20) NOT NULL
);"

mysql -h $MYSQL_HOST -u root -p$MYSQL_ROOT_PASSWORD -e "insert into restapi_db.users (
  first_name, 
  last_name
)
values (
  'Taro',
  'Yamada'
),
(
  'Jiro',
  'Sato'
);"

echo 'MySQL has survey_db and its tables'
