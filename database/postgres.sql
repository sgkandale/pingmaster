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
   ping_timeout INT NOT NULL,
   CONSTRAINT fk_creator FOREIGN KEY(creator) REFERENCES users(name) ON DELETE CASCADE
);

CREATE TABLE pings(
   key VARCHAR(1024) NOT NULL,
   timestamp BIGINT NOT NULL,
   duration INT NOT NULL,
   status_code INT,
   error VARCHAR(4096),
   PRIMARY KEY (key, timestamp),
   CONSTRAINT fk_target FOREIGN KEY(key) REFERENCES targets(key) ON DELETE CASCADE
);