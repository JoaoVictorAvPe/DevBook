CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

CREATE TABLE user (
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(32) not null,
    created_in timestamp default current_timestamp()
 ) ENGINE=INNODB;