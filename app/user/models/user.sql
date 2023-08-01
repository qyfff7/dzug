-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
                        `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
                        `name` varchar(32) NOT NULL DEFAULT '' COMMENT '用户名称',
                        `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码，已加密',
                        `follow_count` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '关注人数',
                        `follower_count` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '粉丝人数',
                        `work_count` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '作品数',
                        `favorite_count` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '点赞视频数',
                        `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                        `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                        `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间，软删除',
                        PRIMARY KEY (`id`),
--                      UNIQUE KEY `uni_name` (`name`) COMMENT '用户名称需要唯一'
--                      给用户名和id创建唯一索引
                        UNIQUE KEY `uni_name` (`name`) USING BTREE COMMENT '用户名唯一',
                        UNIQUE KEY `nui_id` (`id`) USING BTREE COMMENT 'id唯一'
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='用户表';