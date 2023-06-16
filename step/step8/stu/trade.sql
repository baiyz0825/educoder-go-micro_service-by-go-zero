# ---------------------------------------------------------------不要修改这两句------------------------------------------------------
CREATE DATABASE `school_trade`;
use `school_trade`;
# ----------------------------------------------------------在下面编写创建 product 表的Sql语句------------------------------------------------------


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

