CREATE DATABASE pingmaster;

\c pingmaster

CREATE TABLE users(
   name VARCHAR(32) PRIMARY KEY NOT NULL,
   password VARCHAR(2048) NOT NULL,
   created_at BIGINT,
   last_login BIGINT
);