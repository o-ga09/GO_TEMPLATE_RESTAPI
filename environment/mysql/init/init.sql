CREATE TABLE `users` (
    `id` INT AUTO_INCREMENT,
    `created_at` TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP NOT NULL,
    `deleted_at` TIMESTAMP,
    `uid` VARCHAR(255),
    `email` VARCHAR(255),
    `password` VARCHAR(255),
    `user_id` VARCHAR(50),
    `first_name` VARCHAR(50),
    `last_name` VARCHAR(50),
    `gender` VARCHAR(255),
    `birth_day` DATE,
    `phone_number` VARCHAR(20),
    `post_office_number` VARCHAR(20),
    `pref` VARCHAR(20),
    `city` VARCHAR(50),
    `extra` VARCHAR(255),
    PRIMARY KEY (`id`,`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `administrators` (
    `id` INT AUTO_INCREMENT,
    `created_at` TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP NOT NULL,
    `deleted_at` TIMESTAMP,
    `user_id` VARCHAR(255),
    `admin` INT,
    PRIMARY KEY (`id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

GRANT ALL PRIVILEGES ON devdb.* TO 'api'@'%';
GRANT ALL PRIVILEGES ON stgdb.* TO 'api'@'%';
GRANT ALL PRIVILEGES ON proddb.* TO 'api'@'%';

CREATE DATABASE `proddb`;
CREATE DATABASE `stgdb`;
CREATE DATABASE `devdb`;