CREATE DATABASE `school_user` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
use `school_user`;
-- major: table
CREATE TABLE `major`
(
    `id`          bigint                                                       NOT NULL AUTO_INCREMENT COMMENT '专业id',
    `name`        varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '专业名称',
    `desc`        mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '专业描述',
    `delete_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户专业统计表';

-- third: table
CREATE TABLE `third`
(
    `id`              bigint                                                        NOT NULL AUTO_INCREMENT COMMENT '用户三方数据自增ID',
    `user_id`         bigint                                                        NOT NULL COMMENT '用户id',
    `type`            int                                                           NOT NULL DEFAULT '0' COMMENT '微信0、QQ1、Github2、Gitee3 ',
    `access_token`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '默认访问密钥',
    `referesh_token`  varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'not set' COMMENT '刷新token',
    `acctoken_expire` bigint                                                        NOT NULL DEFAULT '180' COMMENT '过期时间',
    `create_time`     timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`     timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time`     timestamp NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户三方关联数据';

-- third_data: table
CREATE TABLE `third_data`
(
    `id`          bigint NOT NULL AUTO_INCREMENT COMMENT '三方数据自增id',
    `third_id`    bigint NOT NULL COMMENT '三方数据id',
    `name`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '三方用户名称',
    `sign`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '三方数据签名',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time` timestamp NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `third_data_third_id_uindex` (`third_id`) COMMENT '用户id'
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='第三方用户数据';

-- user: table
CREATE TABLE `user`
(
    `uid`         bigint                                                        NOT NULL AUTO_INCREMENT COMMENT '数据表自增ID',
    `unique_id`   bigint                                                        NOT NULL COMMENT '用户唯一数据id',
    `name`        varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '用户名称',
    `password`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户密码',
    `age`         bigint                                                        DEFAULT NULL COMMENT '用户年龄',
    `gender`      int                                                           DEFAULT NULL COMMENT '用户性别 1男2女3未设置',
    `phone`       varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '用户电话',
    `email`       varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  DEFAULT NULL COMMENT '用户邮箱',
    `grade`       int                                                           DEFAULT NULL COMMENT '用户年纪 （大一、大二、大三、大四） 1，2，3，4',
    `major`       int                                                           DEFAULT NULL COMMENT '用户专业信息(关联字段)',
    `star`        float                                                         DEFAULT NULL COMMENT '用户等级(0~5)',
    `avatar`      varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '用户头像链接',
    `sign`        mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '用户个性签名',
    `class`       int                                                           DEFAULT NULL COMMENT '用户班级',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`uid`),
    KEY           `user_unique_id_index` (`unique_id`) COMMENT '唯一id'
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户数据基本信息表';

-- No native definition for element: user_unique_id_index (index)
# 主修表测试数据
INSERT INTO school_user.major (id, name, `desc`, delete_time) VALUES (1, '计算机科学', '这是升级版本极客', null);
INSERT INTO school_user.major (id, name, `desc`, delete_time) VALUES (2, 'C语言', '计算机', '2023-04-08 17:47:02');
INSERT INTO school_user.major (id, name, `desc`, delete_time) VALUES (3, 'Java语言', '计算机', null);
INSERT INTO school_user.major (id, name, `desc`, delete_time) VALUES (4, 'Go语言', '计算机', '2023-04-08 18:48:30');
INSERT INTO school_user.major (id, name, `desc`, delete_time) VALUES (5, 'PHP语言', '更新数据', null);
INSERT INTO school_user.major (id, name, `desc`, delete_time) VALUES (6, 'Py语言', '编 程', null);
INSERT INTO school_user.major (id, name, `desc`, delete_time) VALUES (7, 'Rust语言', '这是升级版本极客', null);
INSERT INTO school_user.major (id, name, `desc`, delete_time) VALUES (8, 'R语言', '新语言', null);
INSERT INTO school_user.major (id, name, `desc`, delete_time) VALUES (9, '美术', '分类', null);
INSERT INTO school_user.major (id, name, `desc`, delete_time) VALUES (10, '科技', '分类', null);

# 三方绑定用户数据
INSERT INTO school_user.third_data (id, third_id, name, sign, create_time, update_time, delete_time) VALUES (1, 1, 'Test1QQ', '这是测试数据1的qq', '2023-05-11 10:37:48', '2023-05-11 10:37:48', null);
INSERT INTO school_user.third_data (id, third_id, name, sign, create_time, update_time, delete_time) VALUES (2, 2, 'Test2Gitee', '这是测试数据2的Gitee', '2023-05-11 10:37:48', '2023-05-11 10:37:48', null);

# 三方账户数据
INSERT INTO school_user.third (id, user_id, type, access_token, referesh_token, acctoken_expire, create_time, update_time, delete_time) VALUES (1, 1, 1, 'sdfsdfsdssagsdgas', 'dadsdsaffgsdgssafas', 123456, '2023-05-11 10:34:07', '2023-05-11 10:34:30', null);
INSERT INTO school_user.third (id, user_id, type, access_token, referesh_token, acctoken_expire, create_time, update_time, delete_time) VALUES (2, 2, 3, 'dasfsdafdfs', 'dsfsdfdfsddsdfs', 12345452, '2023-05-11 10:34:30', '2023-05-11 10:34:30', null);

# 用户数据
INSERT INTO school_user.user (uid, unique_id, name, password, age, gender, phone, email, grade, major, star, avatar, sign, class, create_time, update_time, delete_time) VALUES (1, 1653421638231789568, 'test1', 'test1234', 22, 1, '18178114924', 'ovgtmcffm@163.com', 1, 1, 2, 'https://abcd.asas/img', '这是测试数据一', 192, '2023-05-11 10:31:47', '2023-05-11 10:33:17', null);
INSERT INTO school_user.user (uid, unique_id, name, password, age, gender, phone, email, grade, major, star, avatar, sign, class, create_time, update_time, delete_time) VALUES (2, 1653975443474223104, 'test2', 'test12345', 23, 2, '16627447244', 'oaamcffm@163.com', 2, 5, 5, 'https://abcd.asasdff/img', '这是测试数据2', 193, '2023-05-11 10:33:17', '2023-05-11 10:33:17', null);

