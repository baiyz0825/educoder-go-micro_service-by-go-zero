-- MySQL dump 10.13  Distrib 8.0.33, for Linux (x86_64)
--
-- Host: 数据库连接地址    Database: school_trade
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
-- Table structure for table `product`
--

DROP TABLE IF EXISTS `product`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '产品id',
  `uuid` bigint NOT NULL COMMENT '产品唯一标识',
  `name` varchar(255) NOT NULL COMMENT '产品名称',
  `product_bind` bigint NOT NULL DEFAULT '0' COMMENT '绑定的资源id',
  `type` bigint NOT NULL COMMENT '产品分类(与资源分类一致)',
  `owner` bigint NOT NULL COMMENT '产品所属发布人',
  `price` decimal(10,4) NOT NULL COMMENT '产品价格',
  `saled` int NOT NULL DEFAULT '0' COMMENT '是否已销售 0 no 1 yes',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_time` timestamp NULL COMMENT '删除时间',
  `product_poster` varchar(255) DEFAULT NULL COMMENT '产品头图',
  PRIMARY KEY (`id`),
  KEY `product_name_uuid_index` (`name`,`uuid`) COMMENT '产品名称 uuid 联合索引'
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='产品信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product`
--

LOCK TABLES `product` WRITE;
/*!40000 ALTER TABLE `product` DISABLE KEYS */;
INSERT INTO `product` VALUES (6,1654146961923641344,'满江红电影海报',18,1,20,1.2600,0,'2023-05-04 23:32:23','2023-05-04 23:32:23',NULL,'https://graduating-project.oss-cn-chengdu.aliyuncs.com/cache/poster/满江红海报.jpg'),(7,1654147934704373760,'马里奥海报',17,1,20,2.2600,0,'2023-05-04 15:46:28','2023-05-04 15:46:28',NULL,'https://graduating-project.oss-cn-chengdu.aliyuncs.com/cache/poster/马里奥大战小马里奥.jpg');
/*!40000 ALTER TABLE `product` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-06-13 23:34:36
