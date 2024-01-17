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

INSERT INTO `users` (`created_at`, `updated_at`, `uid`, `email`, `password`, `user_id`, `first_name`, `last_name`, `gender`, `birth_day`, `phone_number`, `post_office_number`, `pref`, `city`, `extra`) VALUES 
(NOW(), NOW(), 'uid1', 'user1@example.com', '$2a$10$5iX715a77t55o4754a577..O457457457457457457457457457457457457457457457457457457', 'user1', '山田', '太郎', '男性', '1980-01-01', '090-1234-5678', '123-4567', '東京都', '港区', '住所情報'),
(NOW(), NOW(), 'uid2', 'user2@example.com', '$2a$10$5iX715a77t55o4754a577..O457457457457457457457457457457457457457457457457457457', 'user2', '田中', '花子', '女性', '1990-02-02', '080-2345-6789', '234-5678', '神奈川県', '横浜市', '住所情報'),
(NOW(), NOW(), 'uid3', 'user3@example.com', '$2a$10$5iX715a77t55o4754a577..O457457457457457457457457457457457457457457457457457457', 'user3', '佐藤', '一郎', '男性', '2000-03-03', '070-3456-7890', '345-6789', '埼玉県', 'さいたま市', '住所情報'),
(NOW(), NOW(), 'uid4', 'user4@example.com', '$2a$10$5iX715a77t55o4754a577..O457457457457457457457457457457457457457457457457457457', 'user4', '鈴木', '二郎', '男性', '2010-04-04', '060-4567-8901', '456-7890', '千葉県', '千葉市', '住所情報'),
(NOW(), NOW(), 'uid5', 'user5@example.com', '$2a$10$5iX715a77t55o4754a577..O457457457457457457457457457457457457457457457457457457', 'user5', '伊藤', '三郎', '男性', '2020-05-05', '050-5667-8901', '456-7890', '千葉県', '千葉市', '住所情報');

INSERT INTO `administrators` (`created_at`, `updated_at`, `deleted_at`, `user_id`, `admin`) VALUES 
(NOW(), NOW(), NULL, 'uid1', 1),
(NOW(), NOW(), NULL, 'uid2', 0),
(NOW(), NOW(), NULL, 'uid3', 0),
(NOW(), NOW(), NULL, 'uid4', 1),
(NOW(), NOW(), NULL, 'uid5', 0);
