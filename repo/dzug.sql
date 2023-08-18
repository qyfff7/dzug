SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for message
-- ----------------------------

DROP TABLE IF EXISTS `message`;
CREATE TABLE `message` (
                           `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '消息id',
                           `thread_id` varchar(255) NOT NULL COMMENT '对话id',
                           `message_uuid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '消息uuid',
                           `to_user_id` bigint NOT NULL DEFAULT '0' COMMENT '该消息接收者的id',
                           `from_user_id` bigint NOT NULL DEFAULT '0' COMMENT '该消息发送者的id',
                           `contents` varchar(255) NOT NULL DEFAULT '' COMMENT '消息内容',
                           `create_time` bigint unsigned NOT NULL DEFAULT '0' COMMENT '自设创建时间(unix)',
                           `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                           `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                           `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间，软删除',
                           PRIMARY KEY (`id`),
                           KEY `fk_user_message_to` (`to_user_id`),
                           KEY `fk_user_message_from` (`from_user_id`),
                           CONSTRAINT `fk_user_message_from` FOREIGN KEY (`from_user_id`) REFERENCES `user` (`user_id`),
                           CONSTRAINT `fk_user_message_to` FOREIGN KEY (`to_user_id`) REFERENCES `user` (`user_id`),
                           UNIQUE KEY `uni_message_uuid` (`message_uuid`) COMMENT 'uuid需要唯一'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='消息表';

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
                         `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
                         `created_at` datetime(3) NULL DEFAULT NULL,
                         `updated_at` datetime(3) NULL DEFAULT NULL,
                         `deleted_at` datetime(3) NULL DEFAULT NULL,
                         `user_id` bigint(0) NOT NULL COMMENT '用户真正的id',
                         `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户名称',
                         `background_images` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '主页背景图',
                         `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户头像',
                         `signature` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '个人简介',
                         `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '密码，已加密',
                         `follow_count` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '关注人数',
                         `follower_count` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '粉丝人数',
                         `work_count` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '作品数',
                         `favorite_count` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '点赞视频数',
                         `total_favorited` bigint(0) NULL DEFAULT 0 COMMENT '获赞数',
                         PRIMARY KEY (`id`) USING BTREE,
                         UNIQUE INDEX `user_id`(`user_id`) USING BTREE,
                         UNIQUE INDEX `name`(`name`) USING BTREE,
                         INDEX `idx_user_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 30 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '2023-08-13 14:42:40.323', '2023-08-13 14:42:40.323', NULL, 5174765849415680, 'ddd', NULL, NULL, NULL, '64646487a38998227cbbc23dcad51cd7f76ab2', 0, 0, 0, 0, 0);
INSERT INTO `user` VALUES (2, '2023-08-13 15:48:47.399', '2023-08-13 15:48:47.399', NULL, 5191404976345088, 'aa', NULL, NULL, NULL, '66666687a38998227cbbc23dcad51cd7f76ab2', 0, 0, 0, 0, 0);
INSERT INTO `user` VALUES (3, '2023-08-13 22:04:40.873', '2023-08-13 22:04:40.873', NULL, 5286001098362880, 'aaa', 'aaaaaa', 'aaa', 'aaa', '61616187a38998227cbbc23dcad51cd7f76ab2', 2, 2, 2, 2, 2);
INSERT INTO `user` VALUES (5, '2023-08-14 10:46:23.986', '2023-08-14 10:46:23.986', NULL, 5477693852225536, 'ccc', NULL, NULL, NULL, '63636387a38998227cbbc23dcad51cd7f76ab2', 0, 0, 0, 0, 0);
INSERT INTO `user` VALUES (24, '2023-08-15 14:32:36.243', '2023-08-15 14:32:36.243', NULL, 5897007889649664, 'sss', NULL, NULL, NULL, '73737387a38998227cbbc23dcad51cd7f76ab2', 0, 0, 0, 0, 0);
INSERT INTO `user` VALUES (25, '2023-08-15 14:32:49.947', '2023-08-15 14:32:49.947', NULL, 5897065364197376, 'fff', NULL, NULL, NULL, '66666687a38998227cbbc23dcad51cd7f76ab2', 0, 0, 0, 0, 0);
INSERT INTO `user` VALUES (26, '2023-08-15 14:32:55.450', '2023-08-15 14:32:55.450', NULL, 5897088445452288, 'ggg', NULL, NULL, NULL, '67676787a38998227cbbc23dcad51cd7f76ab2', 0, 0, 0, 0, 0);
INSERT INTO `user` VALUES (27, '2023-08-15 14:33:01.879', '2023-08-15 14:33:01.879', NULL, 5897115414827008, 'hhh', NULL, NULL, NULL, '68686887a38998227cbbc23dcad51cd7f76ab2', 0, 0, 0, 0, 0);
INSERT INTO `user` VALUES (28, '2023-08-15 14:34:11.764', '2023-08-15 14:34:11.764', NULL, 5897408529567744, 'jjj', NULL, NULL, NULL, '6a6a6a87a38998227cbbc23dcad51cd7f76ab2', 0, 0, 0, 0, 0);
INSERT INTO `user` VALUES (29, '2023-08-15 14:55:06.414', '2023-08-15 14:55:06.414', NULL, 5902670913081344, 'kkk', NULL, NULL, NULL, '6b6b6b87a38998227cbbc23dcad51cd7f76ab2', 0, 0, 0, 0, 0);

SET FOREIGN_KEY_CHECKS = 1;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
                           `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
                           `comment_uuid` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '评论uuid',
                           `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '评论作者id',
                           `video_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '评论视频id',
                           `contents` varchar(255) NOT NULL DEFAULT '' COMMENT '评论内容',
                           `create_time` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '自设创建时间(unix)',
                           `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                           `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                           `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间，软删除',
                           PRIMARY KEY (`id`),
                           KEY `fk_user_comment` (`user_id`),
                           KEY `fk_video_comment` (`video_id`),
                           KEY `fk_uuid_comment` (`comment_uuid`),
                           CONSTRAINT `fk_user_comment` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`),
                           CONSTRAINT `fk_video_comment` FOREIGN KEY (`video_id`) REFERENCES `video` (`id`),
                           UNIQUE KEY `uni_comment_uuid` (`comment_uuid`) COMMENT 'uuid需要唯一'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='评论表';

-- ----------------------------
-- Table structure for favorite
-- ----------------------------
DROP TABLE IF EXISTS `favorite`;
CREATE TABLE `favorite` (
                            `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
                            `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户id',
                            `video_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '视频id',
                            `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                            `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                            `deleted_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '软删除',
                            PRIMARY KEY (`id`) USING BTREE,
                            UNIQUE KEY `uk_favorite` (`user_id`,`video_id`,`deleted_at`),
                            KEY `fk_user_favorite` (`user_id`),
                            KEY `fk_video_favorite` (`video_id`),
                            CONSTRAINT `fk_user_favorite` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`),
                            CONSTRAINT `fk_video_favorite` FOREIGN KEY (`video_id`) REFERENCES `video` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COMMENT='点赞表';

-- ----------------------------
-- Table structure for relation
-- ----------------------------
DROP TABLE IF EXISTS `relation`;
CREATE TABLE `relation` (
                            `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
                            `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户id',
                            `to_user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '关注目标的用户id',
                            `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                            `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                            `deleted_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
                            PRIMARY KEY (`id`) USING BTREE,
                            UNIQUE KEY `uk_relation` (`user_id`,`to_user_id`,`deleted_at`),
                            KEY `fk_user_relation` (`user_id`),
                            KEY `fk_user_relation_to` (`to_user_id`),
                            CONSTRAINT `fk_user_relation` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`),
                            CONSTRAINT `fk_user_relation_to` FOREIGN KEY (`to_user_id`) REFERENCES `user` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COMMENT='关注表';

-- ----------------------------
-- Table structure for video
-- ----------------------------
DROP TABLE IF EXISTS `video`;
CREATE TABLE `video` (
                         `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
                         `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'user表主键',
                         `title` varchar(128) NOT NULL DEFAULT '' COMMENT '视频标题',
                         `play_url` varchar(128) NOT NULL DEFAULT '' COMMENT '视频地址',
                         `cover_url` varchar(128) NOT NULL DEFAULT '' COMMENT '封面地址',
                         `favorite_count` int(15) unsigned NOT NULL DEFAULT '0' COMMENT '获赞数量',
                         `comment_count` int(15) unsigned NOT NULL DEFAULT '0' COMMENT '评论数量',
                         `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                         `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                         `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间，软删除',
                         PRIMARY KEY (`id`),
                         KEY `fk_user_video` (`user_id`),
                         CONSTRAINT `fk_user_video` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='视频表';

SET FOREIGN_KEY_CHECKS = 1;


