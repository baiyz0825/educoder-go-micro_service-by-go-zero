CREATE DATABASE `school_order`;
use `school_order`;
-- order: table
CREATE TABLE `order`
(
    `id`                 bigint                                                        NOT NULL AUTO_INCREMENT COMMENT '订单id',
    `uuid`               bigint                                                        NOT NULL COMMENT '唯一订单流水号',
    `product_id`         bigint                                                        NOT NULL COMMENT '商品中的商品id',
    `sys_model`          bigint                                                        NOT NULL DEFAULT '0' COMMENT '订单来源模块（0 商品模块）',
    `status`             int                                                           NOT NULL DEFAULT '0' COMMENT '订单状态（0:创建    1:付款中   2:付款成功    3:已发货   4:用户已确认）',
    `user_id`            bigint                                                        NOT NULL COMMENT '下订单用户id',
    `pay_price`          decimal(10, 4)                                                NOT NULL DEFAULT '0.0000' COMMENT '实际订单生成金额',
    `pay_path`           int                                                           NOT NULL DEFAULT '0' COMMENT '支付渠道（0:微信  1:支付宝 ）        ',
    `pay_path_order_num` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '支付渠道流水号',
    `pay_code_url`       varchar(255) COLLATE utf8mb4_general_ci                                DEFAULT NULL COMMENT '支付渠道二维码链接',
    `create_time`        timestamp                                                     NULL     DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`        timestamp                                                     NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update_time',
    `delete_time`        timestamp                                                     NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY `order_user_id_uuid_index` (`user_id`, `uuid`) COMMENT '用户id和订单uuid '
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT ='订单表';

-- No native definition for element: order_user_id_uuid_index (index)

-- user_earn: table
CREATE TABLE `user_earn`
(
    `id`          bigint         NOT NULL AUTO_INCREMENT COMMENT '统计表id',
    `user_id`     bigint         NOT NULL COMMENT '用户id',
    `earn_num`    decimal(10, 4) NOT NULL COMMENT '用户入账',
    `pay_num`     decimal(10, 4) NOT NULL COMMENT '用户支出价格',
    `create_time` timestamp      NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp      NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time` timestamp      NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `user_earn_pk` (`user_id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT ='用户收入支出统计';

