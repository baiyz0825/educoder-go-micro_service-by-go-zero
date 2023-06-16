CREATE DATABASE `school_resources` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
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
# 分类测试数据
INSERT INTO school_resources.classification (class_id, class_parent_id, class_name, class_resource_num, create_time, update_time, delete_time) VALUES (1, 0, '计算机', 441, '2023-04-09 13:56:07', '2023-04-09 13:56:07', null);
INSERT INTO school_resources.classification (class_id, class_parent_id, class_name, class_resource_num, create_time, update_time, delete_time) VALUES (2, 0, '艺术', 20, '2023-04-09 16:32:19', '2023-04-09 16:32:19', null);
INSERT INTO school_resources.classification (class_id, class_parent_id, class_name, class_resource_num, create_time, update_time, delete_time) VALUES (3, 0, '科技', 1, '2023-04-09 16:32:26', '2023-04-09 16:32:26', null);
INSERT INTO school_resources.classification (class_id, class_parent_id, class_name, class_resource_num, create_time, update_time, delete_time) VALUES (4, 0, '人文', 1, '2023-04-09 16:32:40', '2023-04-09 16:32:40', null);
INSERT INTO school_resources.classification (class_id, class_parent_id, class_name, class_resource_num, create_time, update_time, delete_time) VALUES (5, 1, '人工智能', 11, '2023-04-09 13:56:07', '2023-04-09 13:56:07', null);
INSERT INTO school_resources.classification (class_id, class_parent_id, class_name, class_resource_num, create_time, update_time, delete_time) VALUES (6, 2, '音乐', 12, '2023-04-09 16:33:30', '2023-04-09 16:33:30', null);
INSERT INTO school_resources.classification (class_id, class_parent_id, class_name, class_resource_num, create_time, update_time, delete_time) VALUES (8, 1, '音乐视频分析', 112, '2023-04-09 13:56:07', '2023-04-09 13:56:07', null);
INSERT INTO school_resources.classification (class_id, class_parent_id, class_name, class_resource_num, create_time, update_time, delete_time) VALUES (9, 1, '成分计算机', 112, '2023-04-09 13:56:07', '2023-04-09 13:56:07', null);
INSERT INTO school_resources.classification (class_id, class_parent_id, class_name, class_resource_num, create_time, update_time, delete_time) VALUES (10, 1, '硬件计算机', 112, '2023-04-09 13:56:07', '2023-04-09 13:56:07', null);
INSERT INTO school_resources.classification (class_id, class_parent_id, class_name, class_resource_num, create_time, update_time, delete_time) VALUES (11, 10, '硬件分析', 112, '2023-04-09 13:56:07', '2023-04-09 13:56:07', null);
INSERT INTO school_resources.classification (class_id, class_parent_id, class_name, class_resource_num, create_time, update_time, delete_time) VALUES (12, 10, '硬件组成', 112, '2023-04-09 13:56:07', '2023-04-09 13:56:07', null);
INSERT INTO school_resources.classification (class_id, class_parent_id, class_name, class_resource_num, create_time, update_time, delete_time) VALUES (13, 10, '硬件调用', 112, '2023-04-09 13:56:07', '2023-04-09 13:56:07', null);
INSERT INTO school_resources.classification (class_id, class_parent_id, class_name, class_resource_num, create_time, update_time, delete_time) VALUES (14, 12, '主板', 112, '2023-04-09 13:56:07', '2023-04-09 13:56:07', null);
INSERT INTO school_resources.classification (class_id, class_parent_id, class_name, class_resource_num, create_time, update_time, delete_time) VALUES (15, 12, 'CPU', 112, '2023-04-09 13:56:07', '2023-04-09 13:56:07', null);
INSERT INTO school_resources.classification (class_id, class_parent_id, class_name, class_resource_num, create_time, update_time, delete_time) VALUES (16, 12, '硬盘', 112, '2023-04-09 13:56:07', '2023-04-09 13:56:07', null);
INSERT INTO school_resources.classification (class_id, class_parent_id, class_name, class_resource_num, create_time, update_time, delete_time) VALUES (17, 0, '经济管理学', 22, '2023-04-16 08:17:51', '2023-05-11 12:58:20', null);

# 用户数据统计测试数据
INSERT INTO school_resources.count (id, uid, file_num, video_num, pic_num, storage_size, create_time, update_time, delete_time) VALUES (3, 4, 5, 33, 774, 762, '2023-04-16 16:24:32', '2023-04-16 16:24:32', null);
INSERT INTO school_resources.count (id, uid, file_num, video_num, pic_num, storage_size, create_time, update_time, delete_time) VALUES (5, 5, 22, 33, 774, 762, '2023-04-16 08:25:54', '2023-04-16 08:25:54', '2023-04-16 16:25:53');
INSERT INTO school_resources.count (id, uid, file_num, video_num, pic_num, storage_size, create_time, update_time, delete_time) VALUES (6, 20, 22, 11, 22, 889, '2023-05-04 04:34:40', '2023-05-04 04:34:40', null);

# 用户文件资源测试数据
INSERT INTO school_resources.file (id, uuid, name, obfuscate_name, size, owner, status, type, class, suffix, download_allow, link, create_time, update_time, delete_time, file_poster) VALUES (1, 1645067063393259520, 'java专家变升级update', 'GnFViu', 135, 6, 1, 0, 5, 'pdf', 1, 'https://testnew', '2023-04-16 08:46:11', '2023-04-16 08:46:11', null, null);
INSERT INTO school_resources.file (id, uuid, name, obfuscate_name, size, owner, status, type, class, suffix, download_allow, link, create_time, update_time, delete_time, file_poster) VALUES (2, 1645067370915434496, 'C专家变成', 'qCqS41', 122, 6, 2, 0, 5, 'png', 1, 'https://test.com/', '2023-04-09 22:13:20', '2023-04-09 22:13:20', null, null);
INSERT INTO school_resources.file (id, uuid, name, obfuscate_name, size, owner, status, type, class, suffix, download_allow, link, create_time, update_time, delete_time, file_poster) VALUES (3, 1645067410450944000, 'Go专家变成', 'XLDeGl', 122, 6, 2, 0, 8, 'png', 1, 'https://test.com/', '2023-04-09 22:13:29', '2023-04-09 22:13:29', null, null);
INSERT INTO school_resources.file (id, uuid, name, obfuscate_name, size, owner, status, type, class, suffix, download_allow, link, create_time, update_time, delete_time, file_poster) VALUES (4, 1645067439534247936, 'PHP专家变成', 'fGRSzh', 122, 6, 2, 0, 8, 'png', 1, 'https://test.com/', '2023-04-09 22:13:36', '2023-04-09 22:13:36', null, null);
INSERT INTO school_resources.file (id, uuid, name, obfuscate_name, size, owner, status, type, class, suffix, download_allow, link, create_time, update_time, delete_time, file_poster) VALUES (5, 1645067498225143808, 'Cpp专家变成', 'm4QJ2L', 122, 4, 2, 0, 9, 'png', 1, 'https://test.com/', '2023-04-09 22:13:50', '2023-04-09 22:13:50', null, null);
INSERT INTO school_resources.file (id, uuid, name, obfuscate_name, size, owner, status, type, class, suffix, download_allow, link, create_time, update_time, delete_time, file_poster) VALUES (6, 1645067533461491712, 'docker专家变成', 'LjLT05', 122, 4, 2, 0, 9, 'png', 1, 'https://test.com/', '2023-04-09 22:13:59', '2023-04-09 22:13:59', null, null);
INSERT INTO school_resources.file (id, uuid, name, obfuscate_name, size, owner, status, type, class, suffix, download_allow, link, create_time, update_time, delete_time, file_poster) VALUES (7, 1645067576461496320, 'Devops专家变成', 'aAAS5i', 22, 4, 2, 0, 10, 'png', 1, 'https://test.com/', '2023-04-09 22:14:09', '2023-04-09 22:14:09', null, null);
INSERT INTO school_resources.file (id, uuid, name, obfuscate_name, size, owner, status, type, class, suffix, download_allow, link, create_time, update_time, delete_time, file_poster) VALUES (8, 1645067760520138752, '照相艺术', 'Lc1RyE', 22, 4, 2, 0, 2, 'png', 1, 'https://test.com/', '2023-04-09 22:14:53', '2023-04-09 22:14:53', null, null);
INSERT INTO school_resources.file (id, uuid, name, obfuscate_name, size, owner, status, type, class, suffix, download_allow, link, create_time, update_time, delete_time, file_poster) VALUES (9, 1645067790270337024, '摄影艺术', 'c9VuXT', 22, 4, 2, 0, 2, 'png', 1, 'https://test.com/', '2023-04-09 22:15:00', '2023-04-09 22:15:00', null, null);
INSERT INTO school_resources.file (id, uuid, name, obfuscate_name, size, owner, status, type, class, suffix, download_allow, link, create_time, update_time, delete_time, file_poster) VALUES (10, 1645067814932844544, '电影艺术', 'sFPrcV', 22, 4, 2, 0, 2, 'png', 1, 'https://test.com/', '2023-04-09 14:45:22', '2023-04-09 14:45:22', '2023-04-09 22:45:20', null);
INSERT INTO school_resources.file (id, uuid, name, obfuscate_name, size, owner, status, type, class, suffix, download_allow, link, create_time, update_time, delete_time, file_poster) VALUES (11, 1645068017433841664, '周杰伦', 'ADhqyj', 22, 4, 2, 0, 6, 'png', 1, 'https://test.com/', '2023-05-04 11:34:55', '2023-05-04 11:34:55', '2023-05-04 19:34:53', null);
INSERT INTO school_resources.file (id, uuid, name, obfuscate_name, size, owner, status, type, class, suffix, download_allow, link, create_time, update_time, delete_time, file_poster) VALUES (12, 1645068039906922496, '张信哲', 'waXN9R', 22, 4, 2, 0, 6, 'png', 1, 'https://test.com/', '2023-05-04 11:35:07', '2023-05-04 11:35:07', '2023-05-04 19:35:05', null);
INSERT INTO school_resources.file (id, uuid, name, obfuscate_name, size, owner, status, type, class, suffix, download_allow, link, create_time, update_time, delete_time, file_poster) VALUES (13, 1645068060492566528, '刘德华', 'cz5PHM', 22, 4, 2, 0, 6, 'png', 1, 'https://test.com/', '2023-04-16 08:44:36', '2023-04-16 08:44:36', '2023-04-16 16:44:35', null);
INSERT INTO school_resources.file (id, uuid, name, obfuscate_name, size, owner, status, type, class, suffix, download_allow, link, create_time, update_time, delete_time, file_poster) VALUES (14, 1647519157228408832, '艺术体操', 'kKKveZ', 22, 4, 1, 0, 2, 'txt', 1, 'https://test-update', '2023-04-16 08:44:25', '2023-04-16 16:44:24', null, null);
INSERT INTO school_resources.file (id, uuid, name, obfuscate_name, size, owner, status, type, class, suffix, download_allow, link, create_time, update_time, delete_time, file_poster) VALUES (15, 1654042585544527872, '满江红', 'BFDmrZ', 10957, 20, 2, 3, 1, '.jpg', 1, 'resources/files/fanart.jpg', '2023-05-04 10:54:55', '2023-05-04 10:54:55', '2023-05-04 18:54:50', null);
INSERT INTO school_resources.file (id, uuid, name, obfuscate_name, size, owner, status, type, class, suffix, download_allow, link, create_time, update_time, delete_time, file_poster) VALUES (16, 1654043087757905920, '马里奥', '6JjqIl', 101695, 20, 2, 3, 1, '.jpg', 1, 'resources/files/8wkrSXXdgREVjK0eIXa9LwMIICg.jpg', '2023-05-04 10:54:55', '2023-05-04 10:54:55', '2023-05-04 18:54:48', null);
INSERT INTO school_resources.file (id, uuid, name, obfuscate_name, size, owner, status, type, class, suffix, download_allow, link, create_time, update_time, delete_time, file_poster) VALUES (17, 1654077249529319424, '马里奥大战小马里奥', 'm0eeBN', 101695, 20, 2, 3, 1, '.jpg', 1, 'resources/files/马里奥大战小马里奥.jpg', '2023-05-04 18:55:22', '2023-05-04 18:55:22', null, 'https://graduating-project.oss-cn-chengdu.aliyuncs.com/cache/poster/马里奥大战小马里奥.jpg');
INSERT INTO school_resources.file (id, uuid, name, obfuscate_name, size, owner, status, type, class, suffix, download_allow, link, create_time, update_time, delete_time, file_poster) VALUES (18, 1654077474629226496, '满江红海报', '8GVgsZ', 20870, 20, 2, 3, 1, '.jpg', 1, 'resources/files/满江红海报.jpg', '2023-05-04 18:56:16', '2023-05-04 18:56:16', null, 'https://graduating-project.oss-cn-chengdu.aliyuncs.com/cache/poster/满江红海报.jpg');

# 用户在线文本资源测试数据
INSERT INTO school_resources.online_text (id, uuid, text_name, text_poster, type_suffix, owner, content, class_id, permission, create_time, update_time, delete_time) VALUES (18, 1654079973830430720, 'IPTV搭建教程', 'https://graduating-project.oss-cn-chengdu.aliyuncs.com/cache/poster/IPTV搭建教程.png', 0, 21, 0x69707476E69599E7A88B7878787878787878, 1, 1, '2023-05-04 19:06:12', '2023-05-07 11:52:21', null);
INSERT INTO school_resources.online_text (id, uuid, text_name, text_poster, type_suffix, owner, content, class_id, permission, create_time, update_time, delete_time) VALUES (19, 1654080364240441344, 'PVE加载磁盘虚拟', 'https://graduating-project.oss-cn-chengdu.aliyuncs.com/cache/poster/PVE加载磁盘虚拟.png', 0, 20, 0xE9A696E58588E99C80E8A681E8A681E59CA86077656260E9858DE7BDAEE9A1B5E99DA2E4B8ADEFBC8CE59CA8E2809CE98089E9A1B9E2809DE6A08FE4B8ADE68A8A6042494F5360E79A84E580BCE694B9E68890604F564D4628554546492960EFBC8CE5868DE4BB8EE2809CE7A1ACE4BBB6E2809DE6A08FE7BB99E8AFA5E8999AE68B9FE69CBAE58AA0E4B88AE4B880E4B8AA60454649E7A381E79B9860EFBC8CE8AFA5E7A381E79B98E79A84E4BD9CE794A8E8B79FE794B5E88491E4B8BBE69DBFE4B88AE79A84604E5652414D60E5B7AEE4B88DE5A49AEFBC8CE5B0B1E698AFE794A8E69DA5E5AD98E582A86045464960E79A84E9858DE7BDAEE4BFA1E681AFEFBC8CE4BE8BE5A682E590AFE58AA8E9A1B9E58897E8A1A8E38082E5A682E69E9CE6B2A1E69C89E8BF99E4B8AAE7A381E79B98EFBC8CE6AF8FE6ACA1E9858DE7BDAEE5A5BDE590AFE58AA8E9A1B9E4B98BE5908EEFBC8CE58FAAE8A681E8999AE68B9FE69CBAE4B880E585B3EFBC8CE9858DE7BDAEE4BFA1E681AFE5B0B1E4BC9AE6B688E5A4B1E38082E784B6E5908EE59CA8E8999AE68B9FE69CBAE590AFE58AA8E79A84E697B6E58099E68C89E4B88B6045534360E994AEE8BF9BE585A5E68980E8B093E79A846042494F5360E9858DE7BDAEE7958CE99DA2EFBC8CE4BE9DE6ACA1E98089E68BA960426F6F746020604D61696E74656E616E6365204D616E61676572602D3E60426F6F74204F7074696F6E73602D3E6041646420426F6F74204F7074696F6E60EFBC8CE68EA5E79D80E4BC9AE587BAE69DA5E88BA5E5B9B2E4B8AAE58C85E590ABE4BA866045464960E58886E58CBAE79A84E7A1ACE79B98EFBC88E4B880E888ACE698AF31E4B8AAEFBC89EFBC8CE59B9EE8BDA6E994AEE98089E4B8ADE8AFA5E7A1ACE79B98EFBC8CE4BE9DE6ACA1E98089E68BA9E79BAEE5BD9560454649602D3E60726564686174602D3E60677275622E65666960EFBC8CE8BF99E697B6E58099E4BC9AE587BAE69DA5E4B880E4B8AAE5A1ABE58699E590AFE58AA8E9A1B9E4BFA1E681AFE79A84E7958CE99DA2EFBC8CE68891E59CA860496E70757420746865206465736372697074696F6E60E4B8ADE5A1ABE58699E4BA866063656E746F73362E3760EFBC8CE784B6E5908EE98089E4B8AD60436F6D6D6974204368616E67657320616E64204578697460E38082E8BF99E4B8AAE697B6E58099E79BB4E68EA5E8BF94E59B9EE4BA8660426F6F74204F7074696F6E7360E7958CE99DA2EFBC8CE98089E4B8ADE88F9CE58D95604368616E676520426F6F74204F7264657260E8BF9BE8A18CE590AFE58AA8E9A1B9E9A1BAE5BA8FE79A84E8B083E695B4EFBC8CE68A8AE4B98BE5898DE696B0E6B7BBE58AA0E79A846063656E746F73362E3560E8B083E588B0E69C80E4B88AE99DA2E58DB3E58FAFE38082E784B6E5908EE98089E68BA960436F6D6D69742060604368616E67657320616E64204578697460E8BF94E59B9EE5889AE6898DE79A84E7958CE99DA2EFBC8CE68EA5E79D80E4B880E79BB4E68C896045534360E587BAE58EBBE588B0E69C80E5A496E99DA2E79A84E7958CE99DA2EFBC8CE98089E68BA960436F6E74696E756560E5B0B1E4BC9AE68890E58A9FE587BAE78EB06063656E746F7360E79A84E590AFE58AA8E88F9CE58D95E4BA86E38082, 10, 1, '2023-05-04 19:07:45', '2023-05-04 19:07:45', null);
INSERT INTO school_resources.online_text (id, uuid, text_name, text_poster, type_suffix, owner, content, class_id, permission, create_time, update_time, delete_time) VALUES (20, 1654081830137106432, '马里奥大电影后感', 'https://graduating-project.oss-cn-chengdu.aliyuncs.com/cache/poster/马里奥大电影后感.jpg', 0, 20, 0xE9A9ACE9878CE5A5A5E698AFE4B880E6ACBEE794B1E697A5E69CACE4BBBBE5A4A9E5A082E585ACE58FB8E5BC80E58F91E79A84E4B880E7B3BBE58897E794B5E5AD90E6B8B8E6888FE79A84E4B8BBE8A792E38082E5889BE980A0E8AFA5E8A792E889B2E79A84E8AEBEE8AEA1E5B888E698AFE5AEABE69CACE88C82EFBC8CE69C80E5889DE4BA8E31393831E5B9B4E8AEA9E4BB96E587BAE78EB0E59CA8E8A197E69CBAE6B8B8E6888FE3808AE5A4A7E98791E5889AE3808BE4B8ADE38082E99A8FE5908EEFBC8CE9A9ACE9878CE5A5A5E59CA8E5A49AE4B8AAE6B8B8E6888FE4B8ADE68890E4B8BAE4BA86E4BBBBE5A4A9E5A082E69C80E58F97E6ACA2E8BF8EE79A84E5BDA2E8B1A1E4B98BE4B880EFBC8CE58C85E68BACE3808AE9A9ACE9878CE5A5A5E58584E5BC9FE3808BE38081E3808AE8B685E7BAA7E9A9ACE9878CE5A5A5E58584E5BC9FE3808BE7AD89E7AD89E380820A0AE9A9ACE9878CE5A5A5E698AFE4B880E4B8AAE79FAEE5B08FE38081E59C86E6B6A6E79A84E6848FE5A4A7E588A9E4BABAEFBC8CE9809AE5B8B8E7A9BFE79D80E7BAA2E889B2E88D89E5B8BDE5928CE8838CE5B8A6E8A3A4E38082E4BB96E59CA8E6B8B8E6888FE4B8ADE79A84E4BBBBE58AA1E9809AE5B8B8E698AFE68BAFE69591E4BB96E79A84E5A5B3E69C8BE58F8BE6A183E88AB1E585ACE4B8BBE8A2ABE982AAE681B6E79A84E9BE9FE5928CE89891E88F87E680AAE789A9E7BB91E69EB6E38082E6B8B8E6888FE4B8ADE79A84E4BBBBE58AA1E58C85E68BACE8B7B3E8B783E38081E694B6E99B86E98791E5B881E38081E68993E8B4A5E6958CE4BABAE5928CE8A7A3E586B3E8B09CE9A298E380820A0AE9A9ACE9878CE5A5A5E7B3BBE58897E6B8B8E6888FE4BABAE6B094E69E81E9AB98EFBC8CE4B88DE4BB85E6B7B1E58F97E4BC97E5A49AE78EA9E5AEB6E79A84E5969CE788B1EFBC8CE8BF98E58F97E588B0E4BA86E6B8B8E6888FE4B89AE7958CE79A84E5B9BFE6B39BE8B59EE8AA89E38082E79BAEE5898DE8AFA5E7B3BBE58897E6B8B8E6888FE99480E9878FE5B7B2E7BB8FE8B685E8BF8733E4BABFE4BBBDE38082, 2, 1, '2023-05-04 11:28:32', '2023-05-04 11:28:32', null);
INSERT INTO school_resources.online_text (id, uuid, text_name, text_poster, type_suffix, owner, content, class_id, permission, create_time, update_time, delete_time) VALUES (21, 1654083046904696832, '复仇者联盟后感', 'https://graduating-project.oss-cn-chengdu.aliyuncs.com/cache/poster/复仇者联盟后感.jpg', 0, 20, 0xE5A48DE4BB87E88085E88194E79B9FE698AFE4B880E983A8E794B1E6BCABE5A881E6BCABE794BBE587BAE78988E585ACE58FB8E68980E58F91E8A18CE79A84E8B685E7BAA7E88BB1E99B84E794B5E5BDB1E7B3BBE58897EFBC8CE8AEB2E8BFB0E4BA86E4B880E7BEA4E8B685E7BAA7E88BB1E99B84E88194E6898BE4BF9DE68AA4E59CB0E79083E5858DE58F97E59084E7A78DE5A881E88381E38082E58C85E68BACE992A2E99381E4BEA0E38081E7BE8EE59BBDE9989FE995BFE38081E99BB7E7A59EE38081E9BB91E5AFA1E5A687E38081E7BBBFE5B7A8E4BABAE38081E9B9B0E79CBCE5928CE5A587E5BC82E58D9AE5A3ABE7AD89E7BB8FE585B8E8A792E889B2E38082, 2, 1, '2023-05-04 19:18:24', '2023-05-04 19:18:24', null);

# 资源评论数据
INSERT INTO school_resources.res_comment (id, owner, resource_id, content, create_time, update_time, delete_time) VALUES (1, 5, 6, '这个docker专家凑合吧', '2023-05-04 11:30:39', '2023-05-04 11:30:39', '2023-05-04 19:30:37');
INSERT INTO school_resources.res_comment (id, owner, resource_id, content, create_time, update_time, delete_time) VALUES (2, 4, 6, '这个docker专家不好用', '2023-04-16 10:00:05', '2023-04-16 10:00:05', '2023-04-16 18:00:04');
INSERT INTO school_resources.res_comment (id, owner, resource_id, content, create_time, update_time, delete_time) VALUES (3, 20, 1, 'java专家变升级update这个资源写的不错', '2023-05-04 04:30:20', '2023-05-04 04:30:20', '2023-05-04 12:30:19');
INSERT INTO school_resources.res_comment (id, owner, resource_id, content, create_time, update_time, delete_time) VALUES (4, 21, 1, 'java专家变升级update这个资源写的不错写的还行把，不是特别好把', '2023-05-04 12:16:59', '2023-05-04 12:16:59', null);
INSERT INTO school_resources.res_comment (id, owner, resource_id, content, create_time, update_time, delete_time) VALUES (5, 20, 1, 'java专家变升级update这个资源写的不错写的还行把，凑合可以看', '2023-05-04 12:27:38', '2023-05-04 12:27:38', null);
