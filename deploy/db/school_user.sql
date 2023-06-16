-- MySQL dump 10.13  Distrib 8.0.33, for Linux (x86_64)
--
-- Host: 数据库连接地址    Database: school_user
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
-- Table structure for table `major`
--

DROP TABLE IF EXISTS `major`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `major` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '专业id',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '专业名称',
  `desc` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '专业描述',
  `delete_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户专业统计表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `major`
--

LOCK TABLES `major` WRITE;
/*!40000 ALTER TABLE `major` DISABLE KEYS */;
INSERT INTO `major` VALUES (1,'计算机科学','这是升级版本极客',NULL),(2,'C语言','计算机','2023-04-08 17:47:02'),(3,'Java语言','计算机',NULL),(4,'Go语言','计算机','2023-04-08 18:48:30'),(5,'PHP语言','更新数据',NULL),(6,'Py语言','编 程',NULL),(7,'Rust语言','这是升级版本极客',NULL),(8,'R语言','新语言',NULL),(9,'美术','分类',NULL),(10,'科技','分类',NULL);
/*!40000 ALTER TABLE `major` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `third`
--

DROP TABLE IF EXISTS `third`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `third` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户三方数据自增ID',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `type` int NOT NULL DEFAULT '0' COMMENT '微信0、QQ1、Github2、Gitee3 ',
  `access_token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '默认访问密钥',
  `referesh_token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'not set' COMMENT '刷新token',
  `acctoken_expire` bigint NOT NULL DEFAULT '180' COMMENT '过期时间',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户三方关联数据';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `third`
--

LOCK TABLES `third` WRITE;
/*!40000 ALTER TABLE `third` DISABLE KEYS */;
INSERT INTO `third` VALUES (5,5,0,'dfsgfdsdsfsdfsdfdsfsd','hjfgjhsggfadfdsqedgg',878,'2023-04-08 18:47:50','2023-04-08 18:47:50',NULL),(7,4,1,'gsdfhdjhgjesdfsdw4536yt','gkhgdhrw4465',703,'2023-04-08 11:02:31','2023-04-08 19:02:29',NULL);
/*!40000 ALTER TABLE `third` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `third_data`
--

DROP TABLE IF EXISTS `third_data`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `third_data` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '三方数据自增id',
  `third_id` bigint NOT NULL COMMENT '三方数据id',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '三方用户名称',
  `sign` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '三方数据签名',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_time` timestamp NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `third_data_third_id_uindex` (`third_id`) COMMENT '用户id'
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='第三方用户数据';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `third_data`
--

LOCK TABLES `third_data` WRITE;
/*!40000 ALTER TABLE `third_data` DISABLE KEYS */;
INSERT INTO `third_data` VALUES (3,5,'sgdsddfgf','大伯子不爱吃饭','2023-04-08 18:59:32','2023-04-08 18:59:32',NULL),(4,7,'小花子','大伯子不爱吃饭','2023-04-08 19:03:31','2023-04-08 19:03:31',NULL);
/*!40000 ALTER TABLE `third_data` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `uid` bigint NOT NULL AUTO_INCREMENT COMMENT '数据表自增ID',
  `unique_id` bigint NOT NULL COMMENT '用户唯一数据id',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户名称',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户密码',
  `age` bigint DEFAULT NULL COMMENT '用户年龄',
  `gender` int DEFAULT NULL COMMENT '用户性别 1男2女3未设置',
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户电话',
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '用户邮箱',
  `grade` int DEFAULT NULL COMMENT '用户年纪 （大一、大二、大三、大四） 1，2，3，4',
  `major` int DEFAULT NULL COMMENT '用户专业信息(关联字段)',
  `star` float DEFAULT NULL COMMENT '用户等级(0~5)',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '用户头像链接',
  `sign` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '用户个性签名',
  `class` int DEFAULT NULL COMMENT '用户班级',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`uid`),
  KEY `user_unique_id_index` (`unique_id`) COMMENT '唯一id'
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户数据基本信息表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (20,1653421638231789568,'admin','sBacaca77888as',22,1,'18109273856','byz0825@outlook.com',4,1,0,'https://graduating-project.oss-cn-chengdu.aliyuncs.com/cache/user/oaLAI-1653421638231789568-20-4k5k7.png','我今天很开心，你说对不对押',192,'2023-05-04 12:28:45','2023-05-07 13:55:18',NULL),(21,1653975443474223104,'test1','testAbcdTest12',0,0,'18109273868','',0,0,0,'','',0,'2023-05-04 12:10:50','2023-05-04 12:10:50',NULL),(22,1654101236183470080,'test2','testAbcdTest111',0,0,'18109273861','',0,0,0,'','',0,'2023-05-04 20:30:41','2023-05-04 20:30:41',NULL);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-06-13 23:34:22
