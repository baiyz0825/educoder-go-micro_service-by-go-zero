CREATE DATABASE `user`;
use `user`;
-- user: table
CREATE TABLE `user`
(
    `id`        bigint NOT NULL AUTO_INCREMENT COMMENT '用户id',
    `user_name` varchar(255) DEFAULT NULL COMMENT '用户名',
    `passwd`    varchar(255) DEFAULT NULL COMMENT '密码',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

INSERT INTO user.user (user_name, passwd) VALUES ('bigwhite', 'abcd1234');
