-- MySQL dump 10.13  Distrib 8.0.33, for Linux (x86_64)
--
-- Host: 数据库连接地址    Database: school_resources
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
-- Table structure for table `classification`
--

DROP TABLE IF EXISTS `classification`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `classification` (
  `class_id` bigint NOT NULL AUTO_INCREMENT COMMENT '资源id',
  `class_parent_id` bigint NOT NULL COMMENT '父分类ID',
  `class_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '未设置' COMMENT '分类名称',
  `class_resource_num` bigint DEFAULT NULL COMMENT '分类下资源数量',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_time` timestamp NULL COMMENT '删除时间',
  PRIMARY KEY (`class_id`),
  KEY `classification_class_parent_id_index` (`class_parent_id`) COMMENT '父分类ID索引'
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='资源分类信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `classification`
--

LOCK TABLES `classification` WRITE;
/*!40000 ALTER TABLE `classification` DISABLE KEYS */;
INSERT INTO `classification` VALUES (1,0,'计算机',441,'2023-04-09 13:56:07','2023-04-09 13:56:07',NULL),(2,0,'艺术',20,'2023-04-09 16:32:19','2023-04-09 16:32:19',NULL),(3,0,'科技',1,'2023-04-09 16:32:26','2023-04-09 16:32:26',NULL),(4,0,'人文',1,'2023-04-09 16:32:40','2023-04-09 16:32:40',NULL),(5,1,'人工智能',11,'2023-04-09 13:56:07','2023-04-09 13:56:07',NULL),(6,2,'音乐',12,'2023-04-09 16:33:30','2023-04-09 16:33:30',NULL),(8,1,'音乐视频分析',112,'2023-04-09 13:56:07','2023-04-09 13:56:07',NULL),(9,1,'成分计算机',112,'2023-04-09 13:56:07','2023-04-09 13:56:07',NULL),(10,1,'硬件计算机',112,'2023-04-09 13:56:07','2023-04-09 13:56:07',NULL),(11,10,'硬件分析',112,'2023-04-09 13:56:07','2023-04-09 13:56:07',NULL),(12,10,'硬件组成',112,'2023-04-09 13:56:07','2023-04-09 13:56:07',NULL),(13,10,'硬件调用',112,'2023-04-09 13:56:07','2023-04-09 13:56:07',NULL),(14,12,'主板',112,'2023-04-09 13:56:07','2023-04-09 13:56:07',NULL),(15,12,'CPU',112,'2023-04-09 13:56:07','2023-04-09 13:56:07',NULL),(16,12,'硬盘',112,'2023-04-09 13:56:07','2023-04-09 13:56:07',NULL),(17,0,'经济管理学',22,'2023-04-16 08:17:51','2023-05-11 12:58:20',NULL);
/*!40000 ALTER TABLE `classification` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `count`
--

DROP TABLE IF EXISTS `count`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `count` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户上传数据记录id',
  `uid` bigint NOT NULL COMMENT '用户id',
  `file_num` bigint NOT NULL DEFAULT '0' COMMENT '用户存储文件数量',
  `video_num` bigint NOT NULL DEFAULT '0' COMMENT '用户存储视频数量',
  `pic_num` bigint NOT NULL DEFAULT '0' COMMENT '用户存储图片数量',
  `storage_size` bigint NOT NULL DEFAULT '0' COMMENT '用户存储空间占用（mb）不足mb按mb计算',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_time` timestamp NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `count_uid_uindex` (`uid`) COMMENT '用户id索引'
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户上传资源量统计信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `count`
--

LOCK TABLES `count` WRITE;
/*!40000 ALTER TABLE `count` DISABLE KEYS */;
INSERT INTO `count` VALUES (3,4,5,33,774,762,'2023-04-16 16:24:32','2023-04-16 16:24:32',NULL),(5,5,22,33,774,762,'2023-04-16 08:25:54','2023-04-16 08:25:54','2023-04-16 16:25:53'),(6,20,22,11,22,889,'2023-05-04 04:34:40','2023-05-04 04:34:40',NULL);
/*!40000 ALTER TABLE `count` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `file`
--

DROP TABLE IF EXISTS `file`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `file` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '文件自增id',
  `uuid` bigint NOT NULL COMMENT '文件uuid唯一标识',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '文件名称',
  `obfuscate_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '文件混淆名称',
  `size` bigint NOT NULL COMMENT '文件占用空间大小（kb）',
  `owner` bigint NOT NULL COMMENT '对应用户id',
  `status` int NOT NULL DEFAULT '1' COMMENT '0:已删除（云端） 1:（本地存储状态） 2:（云端存储状态，末态） 3:(用户隐藏状态）',
  `type` int NOT NULL COMMENT '文件所属类型 文本0、文件1、视频2、图片3',
  `class` bigint NOT NULL COMMENT '文件所属分类',
  `suffix` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '文件后缀信息',
  `download_allow` int NOT NULL DEFAULT '1' COMMENT '是否允许查看 0 no 1 yes',
  `link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '文件云端存储目录',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_time` timestamp NULL COMMENT '删除时间',
  `file_poster` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文件头图',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='文件资源存储表（非文本类型）';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `file`
--

LOCK TABLES `file` WRITE;
/*!40000 ALTER TABLE `file` DISABLE KEYS */;
INSERT INTO `file` VALUES (1,1645067063393259520,'java专家变升级update','GnFViu',135,6,1,0,5,'pdf',1,'https://testnew','2023-04-16 08:46:11','2023-04-16 08:46:11',NULL,NULL),(2,1645067370915434496,'C专家变成','qCqS41',122,6,2,0,5,'png',1,'https://test.com/','2023-04-09 22:13:20','2023-04-09 22:13:20',NULL,NULL),(3,1645067410450944000,'Go专家变成','XLDeGl',122,6,2,0,8,'png',1,'https://test.com/','2023-04-09 22:13:29','2023-04-09 22:13:29',NULL,NULL),(4,1645067439534247936,'PHP专家变成','fGRSzh',122,6,2,0,8,'png',1,'https://test.com/','2023-04-09 22:13:36','2023-04-09 22:13:36',NULL,NULL),(5,1645067498225143808,'Cpp专家变成','m4QJ2L',122,4,2,0,9,'png',1,'https://test.com/','2023-04-09 22:13:50','2023-04-09 22:13:50',NULL,NULL),(6,1645067533461491712,'docker专家变成','LjLT05',122,4,2,0,9,'png',1,'https://test.com/','2023-04-09 22:13:59','2023-04-09 22:13:59',NULL,NULL),(7,1645067576461496320,'Devops专家变成','aAAS5i',22,4,2,0,10,'png',1,'https://test.com/','2023-04-09 22:14:09','2023-04-09 22:14:09',NULL,NULL),(8,1645067760520138752,'照相艺术','Lc1RyE',22,4,2,0,2,'png',1,'https://test.com/','2023-04-09 22:14:53','2023-04-09 22:14:53',NULL,NULL),(9,1645067790270337024,'摄影艺术','c9VuXT',22,4,2,0,2,'png',1,'https://test.com/','2023-04-09 22:15:00','2023-04-09 22:15:00',NULL,NULL),(10,1645067814932844544,'电影艺术','sFPrcV',22,4,2,0,2,'png',1,'https://test.com/','2023-04-09 14:45:22','2023-04-09 14:45:22','2023-04-09 22:45:20',NULL),(11,1645068017433841664,'周杰伦','ADhqyj',22,4,2,0,6,'png',1,'https://test.com/','2023-05-04 11:34:55','2023-05-04 11:34:55','2023-05-04 19:34:53',NULL),(12,1645068039906922496,'张信哲','waXN9R',22,4,2,0,6,'png',1,'https://test.com/','2023-05-04 11:35:07','2023-05-04 11:35:07','2023-05-04 19:35:05',NULL),(13,1645068060492566528,'刘德华','cz5PHM',22,4,2,0,6,'png',1,'https://test.com/','2023-04-16 08:44:36','2023-04-16 08:44:36','2023-04-16 16:44:35',NULL),(14,1647519157228408832,'艺术体操','kKKveZ',22,4,1,0,2,'txt',1,'https://test-update','2023-04-16 08:44:25','2023-04-16 16:44:24',NULL,NULL),(15,1654042585544527872,'满江红','BFDmrZ',10957,20,2,3,1,'.jpg',1,'resources/files/fanart.jpg','2023-05-04 10:54:55','2023-05-04 10:54:55','2023-05-04 18:54:50',NULL),(16,1654043087757905920,'马里奥','6JjqIl',101695,20,2,3,1,'.jpg',1,'resources/files/8wkrSXXdgREVjK0eIXa9LwMIICg.jpg','2023-05-04 10:54:55','2023-05-04 10:54:55','2023-05-04 18:54:48',NULL),(17,1654077249529319424,'马里奥大战小马里奥','m0eeBN',101695,20,2,3,1,'.jpg',1,'resources/files/马里奥大战小马里奥.jpg','2023-05-04 18:55:22','2023-05-04 18:55:22',NULL,'https://graduating-project.oss-cn-chengdu.aliyuncs.com/cache/poster/马里奥大战小马里奥.jpg'),(18,1654077474629226496,'满江红海报','8GVgsZ',20870,20,2,3,1,'.jpg',1,'resources/files/满江红海报.jpg','2023-05-04 18:56:16','2023-05-04 18:56:16',NULL,'https://graduating-project.oss-cn-chengdu.aliyuncs.com/cache/poster/满江红海报.jpg');
/*!40000 ALTER TABLE `file` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `online_text`
--

DROP TABLE IF EXISTS `online_text`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `online_text` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '在线文本自增id',
  `uuid` bigint NOT NULL COMMENT '文本uuid',
  `text_name` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文档名称',
  `text_poster` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文本头图',
  `type_suffix` int NOT NULL COMMENT '文本输入格式（0 markdown）',
  `owner` bigint NOT NULL COMMENT '所属用户',
  `content` mediumblob COMMENT '存储实际内容',
  `class_id` bigint NOT NULL COMMENT '所属资源分类id',
  `permission` int NOT NULL DEFAULT '1' COMMENT '是否允许查看 0 no 1 yes',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_time` timestamp NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `online_text_uuid_uindex` (`uuid`) COMMENT 'uuid id；索引'
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='在线文本资源信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `online_text`
--

LOCK TABLES `online_text` WRITE;
/*!40000 ALTER TABLE `online_text` DISABLE KEYS */;
INSERT INTO `online_text` VALUES (18,1654079973830430720,'IPTV搭建教程','https://graduating-project.oss-cn-chengdu.aliyuncs.com/cache/poster/IPTV搭建教程.png',0,21,_binary 'iptv教程xxxxxxxx',1,1,'2023-05-04 19:06:12','2023-05-07 11:52:21',NULL),(19,1654080364240441344,'PVE加载磁盘虚拟','https://graduating-project.oss-cn-chengdu.aliyuncs.com/cache/poster/PVE加载磁盘虚拟.png',0,20,_binary '首先需要要在`web`配置页面中，在“选项”栏中把`BIOS`的值改成`OVMF(UEFI)`，再从“硬件”栏给该虚拟机加上一个`EFI磁盘`，该磁盘的作用跟电脑主板上的`NVRAM`差不多，就是用来存储`EFI`的配置信息，例如启动项列表。如果没有这个磁盘，每次配置好启动项之后，只要虚拟机一关，配置信息就会消失。然后在虚拟机启动的时候按下`ESC`键进入所谓的`BIOS`配置界面，依次选择`Boot` `Maintenance Manager`->`Boot Options`->`Add Boot Option`，接着会出来若干个包含了`EFI`分区的硬盘（一般是1个），回车键选中该硬盘，依次选择目录`EFI`->`redhat`->`grub.efi`，这时候会出来一个填写启动项信息的界面，我在`Input the description`中填写了`centos6.7`，然后选中`Commit Changes and Exit`。这个时候直接返回了`Boot Options`界面，选中菜单`Change Boot Order`进行启动项顺序的调整，把之前新添加的`centos6.5`调到最上面即可。然后选择`Commit ``Changes and Exit`返回刚才的界面，接着一直按`ESC`出去到最外面的界面，选择`Continue`就会成功出现`centos`的启动菜单了。',10,1,'2023-05-04 19:07:45','2023-05-04 19:07:45',NULL),(20,1654081830137106432,'马里奥大电影后感','https://graduating-project.oss-cn-chengdu.aliyuncs.com/cache/poster/马里奥大电影后感.jpg',0,20,_binary '马里奥是一款由日本任天堂公司开发的一系列电子游戏的主角。创造该角色的设计师是宫本茂，最初于1981年让他出现在街机游戏《大金刚》中。随后，马里奥在多个游戏中成为了任天堂最受欢迎的形象之一，包括《马里奥兄弟》、《超级马里奥兄弟》等等。\n\n马里奥是一个矮小、圆润的意大利人，通常穿着红色草帽和背带裤。他在游戏中的任务通常是拯救他的女朋友桃花公主被邪恶的龟和蘑菇怪物绑架。游戏中的任务包括跳跃、收集金币、打败敌人和解决谜题。\n\n马里奥系列游戏人气极高，不仅深受众多玩家的喜爱，还受到了游戏业界的广泛赞誉。目前该系列游戏销量已经超过3亿份。',2,1,'2023-05-04 11:28:32','2023-05-04 11:28:32',NULL),(21,1654083046904696832,'复仇者联盟后感','https://graduating-project.oss-cn-chengdu.aliyuncs.com/cache/poster/复仇者联盟后感.jpg',0,20,_binary '复仇者联盟是一部由漫威漫画出版公司所发行的超级英雄电影系列，讲述了一群超级英雄联手保护地球免受各种威胁。包括钢铁侠、美国队长、雷神、黑寡妇、绿巨人、鹰眼和奇异博士等经典角色。',2,1,'2023-05-04 19:18:24','2023-05-04 19:18:24',NULL);
/*!40000 ALTER TABLE `online_text` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `res_comment`
--

DROP TABLE IF EXISTS `res_comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `res_comment` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '评论自增id',
  `owner` bigint NOT NULL COMMENT '评论所属人信息',
  `resource_id` bigint NOT NULL COMMENT '资源id',
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'none set' COMMENT '评论内容',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_time` timestamp NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `res_comment_owner_resource_id_index` (`owner`,`resource_id`) COMMENT '资源 && 用户 联合索引'
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='文件资源评论信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `res_comment`
--

LOCK TABLES `res_comment` WRITE;
/*!40000 ALTER TABLE `res_comment` DISABLE KEYS */;
INSERT INTO `res_comment` VALUES (1,5,6,'这个docker专家凑合吧','2023-05-04 11:30:39','2023-05-04 11:30:39','2023-05-04 19:30:37'),(2,4,6,'这个docker专家不好用','2023-04-16 10:00:05','2023-04-16 10:00:05','2023-04-16 18:00:04'),(3,20,1,'java专家变升级update这个资源写的不错','2023-05-04 04:30:20','2023-05-04 04:30:20','2023-05-04 12:30:19'),(4,21,1,'java专家变升级update这个资源写的不错写的还行把，不是特别好把','2023-05-04 12:16:59','2023-05-04 12:16:59',NULL),(5,20,1,'java专家变升级update这个资源写的不错写的还行把，凑合可以看','2023-05-04 12:27:38','2023-05-04 12:27:38',NULL);
/*!40000 ALTER TABLE `res_comment` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-06-13 23:35:01
