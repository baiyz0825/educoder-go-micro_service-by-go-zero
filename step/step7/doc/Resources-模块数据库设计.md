[TOC]

---
### 任务描述
本数据库主要存储，用户上传的文件信息、资源信息、文件分类、评论、用户上传文件统计，因此分为以下5个数据表
- file表: 文件表，主要存储用户上传的文件资料，表中包含文件的基本属性和基本链接地址，如用户所输入，文件所属的分类id，文件的下载地址，文件的权限...
- online_text表:在线文本表，主要存储了用户上传的在线文本信息，支持用户进行在线修改在线展示，包含文本的基本属性和所属的分类信息
- comment表:用户资源评价表(商品评价表(这里的商品本质上其实就是指的file表中的一行数据，只是用户上传的文件不一定需要上架售卖))
- classification表:分类表，统一划分资源(文件、在线文本、商品)的基本分类属性，其每一行数据包含一个父类id，父类id为0的时候为一级分类。数据结构事实上是一个多叉树结构。
- count表:用户资源信息统计表，这个表主要进行统计用户上传的资源信息以及用户的文件信息所占的用户空间以及文件数量。
### 相关知识
1. [mysql的基本语法](https://www.runoob.com/mysql/mysql-tutorial.html)
2. 数据库表结构设计

#### 数据库表划分
下面给出这几张表的依赖关系ER图:
![](/step/step7/doc/resources_datbase.drawio.png)
在上述图中，可以看到:
- 文件、在线文本均关联于用户的ID
- 评论详情关联用户ID && 文件ID
- 用户数据统计关联用户ID
- 文件、在线文本关联分类信息ID
### 编程要求
请按照上述图示编写，这四张表的sql语句，其中在sql需要设置对应的主键，在图中每个图最下面的属性是该表的主键，sql语句中不需要体现表与表之间的逻辑关系
修改的文件是: `step/step7/stu/resources.sql`
### 测试说明
平台判定说明:
1. 将执行你的Sql文件 -> 从数据库读出实际创建的表结构
2. 执行标准答案的Sql文件 -> 从数据库读出实际创建的表结构
3. 从数据查询两次创建表结构结果 -> 一致则通过，不一致则不通过