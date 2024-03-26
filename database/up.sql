CREATE USER 'pruebaSaludo'@'localhost' IDENTIFIED BY '123';
GRANT ALL PRIVILEGES ON *.* TO 'pruebaSaludo'@'localhost';

DROP DATABASE IF EXISTS apiSaludos;

CREATE DATABASE apiSaludos;
USE apiSaludos;

DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id VARCHAR(32) PRIMARY KEY,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

DROP TABLE IF EXISTS greets;

CREATE TABLE greets (
    id VARCHAR(32) PRIMARY KEY,
    greet_content VARCHAR(32) NOT NULL,
    lenguage VARCHAR(32) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    user_id VARCHAR(32) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);