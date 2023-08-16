/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 80100
 Source Host           : 127.0.0.1:3306
 Source Schema         : douyin

 Target Server Type    : MySQL
 Target Server Version : 80100
 File Encoding         : 65001

 Date: 13/08/2023 22:08:44
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `user_id` bigint NOT NULL,
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `play_url` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `cover_url` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `favorite_count` bigint NULL DEFAULT 0,
  `comment_count` bigint NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_user_video`(`user_id` ASC) USING BTREE,
  INDEX `idx_videos_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of videos
-- ----------------------------
INSERT INTO `videos` VALUES (10, '2023-08-13 22:07:55.878', '2023-08-13 22:07:55.878', NULL, 111, 'test_title1', 'byte-camp-video.oss-cn-beijing.aliyuncs.com/play/111/test1.mp4', 'byte-camp-video.oss-cn-beijing.aliyuncs.com/play/111/test1.mp4?x-oss-process=video/snapshot,t_0,f_jpg', 0, 0);
INSERT INTO `videos` VALUES (11, '2023-08-13 22:08:05.806', '2023-08-13 22:08:05.806', NULL, 222, 'test_title2', 'byte-camp-video.oss-cn-beijing.aliyuncs.com/play/222/test2.mp4', 'byte-camp-video.oss-cn-beijing.aliyuncs.com/play/222/test2.mp4?x-oss-process=video/snapshot,t_0,f_jpg', 0, 0);
INSERT INTO `videos` VALUES (12, '2023-08-13 22:08:18.068', '2023-08-13 22:08:18.068', NULL, 333, 'test_title3', 'byte-camp-video.oss-cn-beijing.aliyuncs.com/play/333/test3.mp4', 'byte-camp-video.oss-cn-beijing.aliyuncs.com/play/333/test3.mp4?x-oss-process=video/snapshot,t_0,f_jpg', 0, 0);
INSERT INTO `videos` VALUES (13, '2023-08-13 22:08:27.637', '2023-08-13 22:08:27.637', NULL, 444, 'test_title4', 'byte-camp-video.oss-cn-beijing.aliyuncs.com/play/444/test4.mp4', 'byte-camp-video.oss-cn-beijing.aliyuncs.com/play/444/test4.mp4?x-oss-process=video/snapshot,t_0,f_jpg', 0, 0);

SET FOREIGN_KEY_CHECKS = 1;

-- ==============================

DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos`  (
                           `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
                           `created_at` timestamp NULL DEFAULT NULL,
                           `updated_at` timestamp NULL DEFAULT NULL,
                           `deleted_at` timestamp NULL DEFAULT NULL,
                           `user_id` bigint NOT NULL,
                           `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                           `play_url` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                           `cover_url` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                           `favorite_count` bigint NULL DEFAULT 0,
                           `comment_count` bigint NULL DEFAULT 0,
                           PRIMARY KEY (`id`) USING BTREE,
                           INDEX `fk_user_video`(`user_id` ASC) USING BTREE,
                           INDEX `idx_videos_deleted_at`(`deleted_at` ASC) USING BTREE
)




