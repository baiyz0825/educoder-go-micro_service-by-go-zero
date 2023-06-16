CREATE DATABASE `school_resources`;
use `school_resources`;
-- classification: table
CREATE TABLE `classification`
(
    `class_id`           bigint                                                        NOT NULL AUTO_INCREMENT COMMENT '资源id',
    `class_parent_id`    bigint                                                        NOT NULL COMMENT '父分类ID',
    `class_name`         varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '未设置' COMMENT '分类名称',
    `class_resource_num` bigint                                                                 DEFAULT NULL COMMENT '分类下资源数量',
    `create_time`        timestamp                                                     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`        timestamp                                                     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time`        timestamp                                                     NULL COMMENT '删除时间',
    PRIMARY KEY (`class_id`),
    KEY `classification_class_parent_id_index` (`class_parent_id`) COMMENT '父分类ID索引'
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT ='资源分类信息';

-- No native definition for element: classification_class_parent_id_index (index)

-- count: table
CREATE TABLE `count`
(
    `id`           bigint    NOT NULL AUTO_INCREMENT COMMENT '用户上传数据记录id',
    `uid`          bigint    NOT NULL COMMENT '用户id',
    `file_num`     bigint    NOT NULL DEFAULT '0' COMMENT '用户存储文件数量',
    `video_num`    bigint    NOT NULL DEFAULT '0' COMMENT '用户存储视频数量',
    `pic_num`      bigint    NOT NULL DEFAULT '0' COMMENT '用户存储图片数量',
    `storage_size` bigint    NOT NULL DEFAULT '0' COMMENT '用户存储空间占用（mb）不足mb按mb计算',
    `create_time`  timestamp NULL     DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`  timestamp NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time`  timestamp NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `count_uid_uindex` (`uid`) COMMENT '用户id索引'
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT ='用户上传资源量统计信息';

-- file: table
CREATE TABLE `file`
(
    `id`             bigint                                                        NOT NULL AUTO_INCREMENT COMMENT '文件自增id',
    `uuid`           bigint                                                        NOT NULL COMMENT '文件uuid唯一标识',
    `name`           varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '文件名称',
    `obfuscate_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '文件混淆名称',
    `size`           bigint                                                        NOT NULL COMMENT '文件占用空间大小（kb）',
    `owner`          bigint                                                        NOT NULL COMMENT '对应用户id',
    `status`         int                                                           NOT NULL DEFAULT '1' COMMENT '0:已删除（云端） 1:（本地存储状态） 2:（云端存储状态，末态） 3:(用户隐藏状态）',
    `type`           int                                                           NOT NULL COMMENT '文件所属类型 文本0、文件1、视频2、图片3',
    `class`          bigint                                                        NOT NULL COMMENT '文件所属分类',
    `suffix`         varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '文件后缀信息',
    `download_allow` int                                                           NOT NULL DEFAULT '1' COMMENT '是否允许查看 0 no 1 yes',
    `link`           varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci          DEFAULT NULL COMMENT '文件云端存储目录',
    `create_time`    timestamp                                                     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`    timestamp                                                     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time`    timestamp                                                     NULL COMMENT '删除时间',
    `file_poster`    varchar(255) COLLATE utf8mb4_general_ci                                DEFAULT NULL COMMENT '文件头图',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT ='文件资源存储表（非文本类型）';

-- online_text: table
CREATE TABLE `online_text`
(
    `id`          bigint    NOT NULL AUTO_INCREMENT COMMENT '在线文本自增id',
    `uuid`        bigint    NOT NULL COMMENT '文本uuid',
    `text_name`   varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文档名称',
    `text_poster` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文本头图',
    `type_suffix` int       NOT NULL COMMENT '文本输入格式（0 markdown）',
    `owner`       bigint    NOT NULL COMMENT '所属用户',
    `content`     mediumblob COMMENT '存储实际内容',
    `class_id`    bigint    NOT NULL COMMENT '所属资源分类id',
    `permission`  int       NOT NULL                      DEFAULT '1' COMMENT '是否允许查看 0 no 1 yes',
    `create_time` timestamp NULL                          DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NULL                          DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time` timestamp NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `online_text_uuid_uindex` (`uuid`) COMMENT 'uuid id；索引'
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT ='在线文本资源信息';

-- res_comment: table
CREATE TABLE `res_comment`
(
    `id`          bigint                                                        NOT NULL AUTO_INCREMENT COMMENT '评论自增id',
    `owner`       bigint                                                        NOT NULL COMMENT '评论所属人信息',
    `resource_id` bigint                                                        NOT NULL COMMENT '资源id',
    `content`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'none set' COMMENT '评论内容',
    `create_time` timestamp                                                     NULL     DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp                                                     NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time` timestamp                                                     NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY `res_comment_owner_resource_id_index` (`owner`, `resource_id`) COMMENT '资源 && 用户 联合索引'
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT ='文件资源评论信息';

-- No native definition for element: res_comment_owner_resource_id_index (index)

