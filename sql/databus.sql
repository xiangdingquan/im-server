SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for app
-- ----------------------------
DROP TABLE IF EXISTS `app`;
CREATE TABLE `app` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `app_key` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL,
  `app_secret` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of app
-- ----------------------------
INSERT INTO `app` VALUES ('1', '170e302355453683', '3d0e8db7bed0503949e545a469789279', '2022-08-28 15:01:01', '2022-08-28 15:01:01');

-- ----------------------------
-- Table structure for auth
-- ----------------------------
DROP TABLE IF EXISTS `auth`;
CREATE TABLE `auth` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `app_id` int(11) NOT NULL,
  `topic_id` int(11) NOT NULL,
  `group2` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL,
  `operation` tinyint(4) NOT NULL,
  `number` bigint(20) NOT NULL,
  `is_delete` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of auth
-- ----------------------------
INSERT INTO `auth` VALUES ('1', '1', '1', 'Sync-MainCommunity-S', '3', '0', '0', '2022-08-28 15:10:05', '2022-08-28 09:52:47');
INSERT INTO `auth` VALUES ('2', '1', '1', 'Sync-MainCommunity-P', '3', '0', '0', '2022-08-28 15:10:05', '2022-08-28 09:52:47');
INSERT INTO `auth` VALUES ('3', '1', '2', 'Inbox-MainCommunity-S', '3', '0', '0', '2022-08-28 15:10:05', '2022-08-28 09:52:47');
INSERT INTO `auth` VALUES ('4', '1', '2', 'Inbox-MainCommunity-P', '3', '0', '0', '2022-08-28 15:10:05', '2022-08-28 09:52:47');
INSERT INTO `auth` VALUES ('5', '1', '3', 'Bots-MainCommunity-S', '3', '0', '0', '2022-08-28 15:10:05', '2022-08-28 09:52:47');
INSERT INTO `auth` VALUES ('6', '1', '3', 'Bots-MainCommunity-P', '3', '0', '0', '2022-08-28 15:10:05', '2022-08-28 09:52:47');
INSERT INTO `auth` VALUES ('7', '1', '4', 'Push-MainCommunity-S', '3', '0', '0', '2022-08-28 15:10:05', '2022-08-28 09:52:47');
INSERT INTO `auth` VALUES ('8', '1', '4', 'Push-MainCommunity-P', '3', '0', '0', '2022-08-28 15:10:05', '2022-08-28 09:52:47');
INSERT INTO `auth` VALUES ('9', '1', '5', 'AdminLog-MainCommunity-S', '3', '0', '0', '2022-08-28 15:10:05', '2020-03-06 07:35:26');
INSERT INTO `auth` VALUES ('10', '1', '5', 'AdminLog-MainCommunity-P', '3', '0', '0', '2022-08-28 15:10:05', '2020-03-06 07:35:30');
INSERT INTO `auth` VALUES ('11', '1', '6', 'GifBot-MainCommunity-S', '3', '0', '0', '2022-08-28 15:10:05', '2020-03-06 07:35:26');
INSERT INTO `auth` VALUES ('12', '1', '6', 'GifBot-MainCommunity-P', '3', '0', '0', '2022-08-28 15:10:05', '2020-03-12 03:34:58');

-- ----------------------------
-- Table structure for topic
-- ----------------------------
DROP TABLE IF EXISTS `topic`;
CREATE TABLE `topic` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `topic` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `cluster` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of topic
-- ----------------------------
INSERT INTO `topic` VALUES ('1', 'Sync-T', 'databus_kafka_9092-266', '2022-08-28 15:08:57', '2022-08-28 09:52:02');
INSERT INTO `topic` VALUES ('2', 'Inbox-T', 'databus_kafka_9092-266', '2022-08-28 15:08:57', '2022-08-28 09:52:02');
INSERT INTO `topic` VALUES ('3', 'Bots-T', 'databus_kafka_9092-266', '2022-08-28 15:08:57', '2022-08-28 09:52:02');
INSERT INTO `topic` VALUES ('4', 'Push-T', 'databus_kafka_9092-266', '2022-08-28 15:08:57', '2022-08-28 09:52:02');
INSERT INTO `topic` VALUES ('5', 'AdminLog-T', 'databus_kafka_9092-266', '2022-08-28 15:08:57', '2022-08-28 09:52:02');
INSERT INTO `topic` VALUES ('6', 'GifBot-T', 'databus_kafka_9092-266', '2022-08-28 15:08:57', '2022-08-28 09:52:02');
