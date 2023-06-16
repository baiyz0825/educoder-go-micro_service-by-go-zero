-- MySQL dump 10.13  Distrib 8.0.33, for Linux (x86_64)
--
-- Host: 数据库连接地址    Database: school_order
-- ------------------------------------------------------
-- Server version	8.0.33-0ubuntu0.22.04.2

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `order`
--

DROP TABLE IF EXISTS `order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `order` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '订单id',
  `uuid` bigint NOT NULL COMMENT '唯一订单流水号',
  `product_id` bigint NOT NULL COMMENT '商品中的商品id',
  `sys_model` bigint NOT NULL DEFAULT '0' COMMENT '订单来源模块（0 商品模块）',
  `status` int NOT NULL DEFAULT '0' COMMENT '订单状态（0:创建    1:付款中   2:付款成功    3:已发货   4:用户已确认）',
  `user_id` bigint NOT NULL COMMENT '下订单用户id',
  `pay_price` decimal(10,4) NOT NULL DEFAULT '0.0000' COMMENT '实际订单生成金额',
  `pay_path` int NOT NULL DEFAULT '0' COMMENT '支付渠道（0:微信  1:支付宝 ）        ',
  `pay_path_order_num` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '支付渠道流水号',
  `pay_code_url` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '支付渠道二维码链接',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update_time',
  `delete_time` timestamp NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `order_user_id_uuid_index` (`user_id`,`uuid`) COMMENT '用户id和订单uuid '
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='订单表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order`
--

LOCK TABLES `order` WRITE;
/*!40000 ALTER TABLE `order` DISABLE KEYS */;
INSERT INTO `order` VALUES (26,1654420328278921216,7,0,2,20,2.2600,1,'1654420328278921216',NULL,'2023-05-05 15:51:39','2023-05-06 03:33:25',NULL),(27,1654423042362707968,7,0,2,20,2.2600,1,'1654423042362707968',NULL,'2023-05-05 15:51:39','2023-05-06 03:32:51',NULL),(28,1659187327165009920,7,0,1,20,2.2600,1,'1659187327165009920','https://qr.alipay.com/bax09712lwyad1sgds9b0096','2023-05-18 21:21:00','2023-05-18 21:21:01',NULL);
/*!40000 ALTER TABLE `order` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_earn`
--

DROP TABLE IF EXISTS `user_earn`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_earn` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '统计表id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `earn_num` decimal(10,4) NOT NULL COMMENT '用户入账',
  `pay_num` decimal(10,4) NOT NULL COMMENT '用户支出价格',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_time` timestamp NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_earn_pk` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户收入支出统计';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_earn`
--

LOCK TABLES `user_earn` WRITE;
/*!40000 ALTER TABLE `user_earn` DISABLE KEYS */;
INSERT INTO `user_earn` VALUES (4,20,98.9400,7.9100,'2023-05-05 17:21:25','2023-05-05 17:21:25',NULL);
/*!40000 ALTER TABLE `user_earn` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-06-13 23:33:56
