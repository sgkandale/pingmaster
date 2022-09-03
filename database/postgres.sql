CREATE DATABASE pingmaster;

\c pingmaster

CREATE TABLE users(
   name VARCHAR(32) PRIMARY KEY NOT NULL,
   password VARCHAR(2048) NOT NULL,
   created_at BIGINT,
   last_login BIGINT
);

CREATE TABLE targets(
   key VARCHAR(1024) PRIMARY KEY NOT NULL,
   name VARCHAR(1024) NOT NULL,
   type  VARCHAR(32) NOT NULL,
   creator VARCHAR(32) NOT NULL,
   protocol VARCHAR(32) NOT NULL,
   host_address VARCHAR(1024) NOT NULL,
   port INT NOT NULL,
   ping_interval INT NOT NULL,
   CONSTRAINT fk_creator FOREIGN KEY(creator) REFERENCES users(name) ON DELETE CASCADE
);