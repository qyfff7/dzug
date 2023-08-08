

## user表

| 属性名           | 描述                                             | 类型 |
| :--------------- | ------------------------------------------------ | ---- |
| id               | user表的自增id（主键）                           |      |
| user_id          | 用户真正的id，用雪花算法生成（高效、安全、唯一） |      |
| name             | 用户名                                           |      |
| password         | 密码（将采用加密算法进行加密）                   |      |
| follow_count     | 关注总数                                         |      |
| follower_count   | 粉丝总数                                         |      |
| avatar           | 用户头像                                         |      |
| background_image | 用户个人页顶部大图                               |      |
| signature        | 个人简介                                         |      |
| total_favorited  | 获赞数量                                         |      |
| work_count       | 作品数量                                         |      |
| favorite_count   | 点赞数量                                         |      |





```sql
-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id` bigint(20) NOT NULL COMMENT '用户真正的ID',
    `name` varchar(32) NOT NULL DEFAULT '' COMMENT '用户名称',
    `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码，已加密',
    `follow_count` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '关注人数',
    `follower_count` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '粉丝人数',
    `work_count` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '作品数',
    `favorite_count` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '点赞视频数',
	`background_images` varchar(255) DEFAULT '' COMMENT '主页背景图',
    `avatar` varchar(255) DEFAULT '' COMMENT '用户头像',
    `signature` varchar(255) DEFAULT '' COMMENT '个人简介',
    `total_favorited` bigint(20) DEFAULT '0' COMMENT '获赞数',
     PRIMARY KEY (`id`),
--  给用户名和id创建唯一索引
    UNIQUE KEY `uni_name` (`name`) USING BTREE COMMENT '用户名唯一',
    UNIQUE KEY `uni_id` (`user_id`) USING BTREE COMMENT 'id唯一'
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='用户表';
```

