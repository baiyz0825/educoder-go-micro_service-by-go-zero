CREATE DATABASE `school_trade` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
use `school_trade`;
-- product: table
CREATE TABLE `product`
(
    `id`             bigint         NOT NULL AUTO_INCREMENT COMMENT '产品id',
    `uuid`           bigint         NOT NULL COMMENT '产品唯一标识',
    `name`           varchar(255)   NOT NULL COMMENT '产品名称',
    `product_bind`   bigint         NOT NULL DEFAULT '0' COMMENT '绑定的资源id',
    `type`           bigint         NOT NULL COMMENT '产品分类(与资源分类一致)',
    `owner`          bigint         NOT NULL COMMENT '产品所属发布人',
    `price`          decimal(10, 4) NOT NULL COMMENT '产品价格',
    `saled`          int            NOT NULL DEFAULT '0' COMMENT '是否已销售 0 no 1 yes',
    `create_time`    timestamp      NULL     DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`    timestamp      NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time`    timestamp      NULL COMMENT '删除时间',
    `product_poster` varchar(255)            DEFAULT NULL COMMENT '产品头图',
    PRIMARY KEY (`id`),
    KEY `product_name_uuid_index` (`name`, `uuid`) COMMENT '产品名称 uuid 联合索引'
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='产品信息';

-- No native definition for element: product_name_uuid_index (index)
# 产品测试数据
INSERT INTO school_trade.product (id, uuid, name, product_bind, type, owner, price, saled, create_time, update_time, delete_time, product_poster) VALUES (6, 1654146961923641344, '满江红电影海报', 18, 1, 20, 1.2600, 0, '2023-05-04 23:32:23', '2023-05-04 23:32:23', null, 'https://graduating-project.oss-cn-chengdu.aliyuncs.com/cache/poster/满江红海报.jpg');
INSERT INTO school_trade.product (id, uuid, name, product_bind, type, owner, price, saled, create_time, update_time, delete_time, product_poster) VALUES (7, 1654147934704373760, '马里奥海报', 17, 1, 20, 2.2600, 0, '2023-05-04 15:46:28', '2023-05-04 15:46:28', null, 'https://graduating-project.oss-cn-chengdu.aliyuncs.com/cache/poster/马里奥大战小马里奥.jpg');

