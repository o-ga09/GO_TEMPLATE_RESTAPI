CREATE DATABASE IF NOT EXISTS example;
USE example;

CREATE TABLE IF NOT EXISTS user (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `userid` varchar(255) NOT NULL,
    `username` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO user (userid, username) VALUES ('000000000', 'testuser1');
INSERT INTO user (userid, username) VALUES ('000000001', 'testuser2');