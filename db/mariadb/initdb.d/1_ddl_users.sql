create database if not exists restapi_db;
use restapi_db;

create table restapi_db.users(
    id int primary key auto_increment, 
    first_name varchar(100) not null,
    last_name varchar(100) not null
);