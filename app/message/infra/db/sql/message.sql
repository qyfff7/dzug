DROP TABLE IF EXISTS `message`;
CREATE TABLE `message` (
                           `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '消息id',
                           `thread_id` bigint unsigned NOT NULL COMMENT '对话id',
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