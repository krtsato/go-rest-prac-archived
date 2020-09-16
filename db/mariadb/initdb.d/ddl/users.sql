create database go_rest_sample;
use go_rest_sample;

create table users(
    id int primary key auto_increment, 
    first_name varchar(100) not null,
    last_name varchar(100) not null,
);