-- 创建 member 数据库
CREATE DATABASE IF NOT EXISTS member;

-- 创建 user 用户表
CREATE TABLE IF NOT EXISTS `user` (
    `id` INT UNSIGNED AUTO_INCREMENT,
    `name` VARCHAR(32) NOT NULL,
    `nickname` VARCHAR(32) NOT NULL,
    `email` VARCHAR(32) NOT NULL,
    `phone_number` VARCHAR(32) NOT NULL,
    `bcrypt_pw` VARCHAR(64) NOT NULL,
    `group` VARCHAR(32) NOT NULL,
    `status` INT NOT NULL,
    `create_date` DATE NOT NULL,
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;