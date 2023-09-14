SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for app_configs
-- ----------------------------
DROP TABLE IF EXISTS `app_configs`;
CREATE TABLE `app_configs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `client_type` tinyint(4) NOT NULL DEFAULT '0',
  `key2` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `type2` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'string',
  `value2` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL,
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of app_configs
-- ----------------------------
INSERT INTO `app_configs` VALUES ('1', '0', 'emojies_animated_zoom', 'number', '0.625', '0', '2022-08-30 09:58:15', '2022-08-30 11:20:50');
INSERT INTO `app_configs` VALUES ('2', '0', 'youtube_pip', 'string', 'disabled', '0', '2022-08-30 10:00:39', '2022-08-30 10:00:39');
INSERT INTO `app_configs` VALUES ('3', '0', 'background_connection', 'bool', 'true', '0', '2022-08-30 10:01:16', '2022-08-30 20:45:21');
INSERT INTO `app_configs` VALUES ('4', '0', 'keep_alive_service', 'bool', 'true', '0', '2022-08-30 10:01:53', '2022-08-30 20:46:48');
INSERT INTO `app_configs` VALUES ('5', '0', 'qr_login_camera', 'bool', 'true', '0', '2022-08-30 10:02:43', '2022-08-30 19:31:56');
INSERT INTO `app_configs` VALUES ('6', '0', 'qr_login_code', 'string', 'primary', '0', '2022-08-30 10:04:17', '2022-08-30 19:34:15');
INSERT INTO `app_configs` VALUES ('7', '0', 'dialog_filters_enabled', 'bool', 'true', '0', '2022-08-30 10:05:13', '2022-08-30 10:57:17');
INSERT INTO `app_configs` VALUES ('8', '0', 'dialog_filters_tooltip', 'bool', 'true', '0', '2022-08-30 10:05:13', '2022-08-30 10:57:10');
INSERT INTO `app_configs` VALUES ('17', '0', 'emojies_send_dice', 'array', '[\"?\",\"?\"]', '0', '2022-08-30 10:26:17', '2022-08-30 10:26:17');
INSERT INTO `app_configs` VALUES ('18', '0', 'emojies_send_dice_success', 'object', '{\"?\":{\"frame_start\":62,\"value\":6}}}', '0', '2022-08-30 10:26:59', '2022-08-30 10:26:59');

-- ----------------------------
-- Table structure for app_languages
-- ----------------------------
DROP TABLE IF EXISTS `app_languages`;
CREATE TABLE `app_languages` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `app` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `lang_code` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `version` int(11) NOT NULL,
  `strings_count` int(11) NOT NULL DEFAULT '0',
  `translated_count` int(11) NOT NULL DEFAULT '0',
  `state` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `app` (`app`,`lang_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of app_languages
-- ----------------------------

INSERT INTO `app_languages` (`id`, `app`, `lang_code`, `version`, `strings_count`, `translated_count`, `state`, `created_at`, `updated_at`) VALUES
(1, 'android', 'ar-raw', 1039455, 3140, 3140, 0, '2022-08-30 15:58:14', '2022-08-30 15:58:14'),
(2, 'android', 'ca', 1035769, 3140, 3140, 0, '2022-08-30 15:58:14', '2022-08-30 15:58:14'),
(3, 'android', 'classic-zh-cn', 1052642, 3140, 3140, 0, '2022-08-30 15:58:14', '2022-08-30 15:58:14'),
(4, 'android', 'de', 1041075, 3140, 3140, 0, '2022-08-30 15:58:14', '2022-08-30 15:58:14'),
(5, 'android', 'en', 1028848, 3140, 3140, 0, '2022-08-30 15:58:14', '2022-08-30 15:58:14'),
(6, 'android', 'es', 1049315, 3140, 3140, 0, '2022-08-30 15:58:14', '2022-08-30 15:58:14'),
(7, 'android', 'fa-raw', 1028885, 3140, 3140, 0, '2022-08-30 15:58:14', '2022-08-30 15:58:14'),
(8, 'android', 'fr', 1046686, 3140, 3140, 0, '2022-08-30 15:58:14', '2022-08-30 15:58:14'),
(9, 'android', 'it', 1045968, 3140, 3140, 0, '2022-08-30 15:58:14', '2022-08-30 15:58:14'),
(10, 'android', 'ko-raw', 1042026, 3140, 3140, 0, '2022-08-30 15:58:14', '2022-08-30 15:58:14'),
(11, 'android', 'nl', 1035125, 3140, 3140, 0, '2022-08-30 15:58:14', '2022-08-30 15:58:14'),
(12, 'android', 'pt-br', 1049471, 3140, 3140, 0, '2022-08-30 15:58:14', '2022-08-30 15:58:14'),
(13, 'android', 'ru', 1043805, 3140, 3140, 0, '2022-08-30 15:58:14', '2022-08-30 15:58:14'),
(14, 'android', 'tr', 1040158, 3140, 3140, 0, '2022-08-30 15:58:14', '2022-08-30 15:58:14'),
(15, 'android', 'uk', 1049844, 3140, 3140, 0, '2022-08-30 15:58:14', '2022-08-30 15:58:14'),
(16, 'android', 'zh-hans-raw', 1048594, 3140, 3140, 0, '2022-08-30 15:58:14', '2022-08-30 15:58:14'),
(17, 'android', 'zh-hant-raw', 1045457, 3140, 3140, 0, '2022-08-30 15:58:14', '2022-08-30 15:58:14'),
(18, 'ios', 'ar-raw', 748380, 3838, 3838, 0, '2022-08-30 15:58:42', '2022-08-30 15:58:42'),
(19, 'ios', 'ca', 746233, 3838, 3838, 0, '2022-08-30 15:58:42', '2022-08-30 15:58:42'),
(20, 'ios', 'classic-zh-cn', 762800, 3838, 3832, 0, '2022-08-30 15:58:42', '2022-08-30 15:58:42'),
(21, 'ios', 'de', 749602, 3838, 3837, 0, '2022-08-30 15:58:42', '2022-08-30 15:58:42'),
(22, 'ios', 'en', 738421, 3838, 3838, 0, '2022-08-30 15:58:42', '2022-08-30 15:58:42'),
(23, 'ios', 'es', 757651, 3838, 3837, 0, '2022-08-30 15:58:42', '2022-08-30 15:58:42'),
(24, 'ios', 'fa-raw', 743043, 3838, 3838, 0, '2022-08-30 15:58:42', '2022-08-30 15:58:42'),
(25, 'ios', 'fr', 738459, 3838, 3837, 0, '2022-08-30 15:58:42', '2022-08-30 15:58:42'),
(26, 'ios', 'it', 752331, 3838, 3838, 0, '2022-08-30 15:58:42', '2022-08-30 15:58:42'),
(27, 'ios', 'ko-raw', 750821, 3838, 3729, 0, '2022-08-30 15:58:42', '2022-08-30 15:58:42'),
(28, 'ios', 'nl', 745918, 3838, 3837, 0, '2022-08-30 15:58:42', '2022-08-30 15:58:42'),
(29, 'ios', 'pt-br', 753162, 3838, 3838, 0, '2022-08-30 15:58:42', '2022-08-30 15:58:42'),
(30, 'ios', 'ru', 742822, 3838, 3838, 0, '2022-08-30 15:58:42', '2022-08-30 15:58:42'),
(31, 'ios', 'tr', 742567, 3838, 3838, 0, '2022-08-30 15:58:42', '2022-08-30 15:58:42'),
(32, 'ios', 'uk', 758358, 3838, 3837, 0, '2022-08-30 15:58:42', '2022-08-30 15:58:42'),
(33, 'ios', 'zh-hans-raw', 743495, 3838, 2711, 0, '2022-08-30 15:58:42', '2022-08-30 15:58:42'),
(34, 'ios', 'zh-hant-raw', 755350, 3838, 3838, 0, '2022-08-30 15:58:42', '2022-08-30 15:58:42'),
(35, 'macos', 'ar-raw', 173445, 2566, 2538, 0, '2022-08-30 15:58:59', '2022-08-30 15:58:59'),
(36, 'macos', 'ca', 173423, 2566, 2464, 0, '2022-08-30 15:58:59', '2022-08-30 15:58:59'),
(37, 'macos', 'classic-zh-cn', 175538, 2566, 2331, 0, '2022-08-30 15:58:59', '2022-08-30 15:58:59'),
(38, 'macos', 'de', 173600, 2566, 2566, 0, '2022-08-30 15:58:59', '2022-08-30 15:58:59'),
(39, 'macos', 'en', 172709, 2566, 2566, 0, '2022-08-30 15:58:59', '2022-08-30 15:58:59'),
(40, 'macos', 'es', 174932, 2566, 2566, 0, '2022-08-30 15:58:59', '2022-08-30 15:58:59'),
(41, 'macos', 'fa-raw', 172871, 2566, 2566, 0, '2022-08-30 15:58:59', '2022-08-30 15:58:59'),
(42, 'macos', 'fr', 174387, 2566, 2566, 0, '2022-08-30 15:58:59', '2022-08-30 15:58:59'),
(43, 'macos', 'it', 173894, 2566, 2566, 0, '2022-08-30 15:58:59', '2022-08-30 15:58:59'),
(44, 'macos', 'ko-raw', 173663, 2566, 2164, 0, '2022-08-30 15:58:59', '2022-08-30 15:58:59'),
(45, 'macos', 'nl', 173053, 2566, 2566, 0, '2022-08-30 15:58:59', '2022-08-30 15:58:59'),
(46, 'macos', 'pt-br', 174075, 2566, 2566, 0, '2022-08-30 15:58:59', '2022-08-30 15:58:59'),
(47, 'macos', 'ru', 172710, 2566, 2566, 0, '2022-08-30 15:58:59', '2022-08-30 15:58:59'),
(48, 'macos', 'tr', 172780, 2566, 2566, 0, '2022-08-30 15:58:59', '2022-08-30 15:58:59'),
(49, 'macos', 'uk', 173049, 2566, 2566, 0, '2022-08-30 15:58:59', '2022-08-30 15:58:59'),
(50, 'macos', 'zh-hans-raw', 175267, 2566, 1795, 0, '2022-08-30 15:58:59', '2022-08-30 15:58:59'),
(51, 'macos', 'zh-hant-raw', 174281, 2566, 2566, 0, '2022-08-30 15:58:59', '2022-08-30 15:58:59'),
(52, 'tdesktop', 'ar-raw', 307237, 2217, 2217, 0, '2022-08-30 15:59:16', '2022-08-30 15:59:16'),
(53, 'tdesktop', 'ca', 309343, 2217, 2217, 0, '2022-08-30 15:59:16', '2022-08-30 15:59:16'),
(54, 'tdesktop', 'classic-zh-cn', 309310, 2217, 2214, 0, '2022-08-30 15:59:16', '2022-08-30 15:59:16'),
(55, 'tdesktop', 'de', 309298, 2217, 2217, 0, '2022-08-30 15:59:16', '2022-08-30 15:59:16'),
(56, 'tdesktop', 'en', 305939, 2217, 2217, 0, '2022-08-30 15:59:16', '2022-08-30 15:59:16'),
(57, 'tdesktop', 'es', 306406, 2217, 2217, 0, '2022-08-30 15:59:16', '2022-08-30 15:59:16'),
(58, 'tdesktop', 'fa-raw', 307279, 2217, 2217, 0, '2022-08-30 15:59:16', '2022-08-30 15:59:16'),
(59, 'tdesktop', 'fr', 307050, 2217, 2217, 0, '2022-08-30 15:59:16', '2022-08-30 15:59:16'),
(60, 'tdesktop', 'it', 306520, 2217, 2217, 0, '2022-08-30 15:59:16', '2022-08-30 15:59:16'),
(61, 'tdesktop', 'ko-raw', 304263, 2217, 2042, 0, '2022-08-30 15:59:16', '2022-08-30 15:59:16'),
(62, 'tdesktop', 'nl', 306629, 2217, 2217, 0, '2022-08-30 15:59:16', '2022-08-30 15:59:16'),
(63, 'tdesktop', 'pt-br', 306647, 2217, 2217, 0, '2022-08-30 15:59:16', '2022-08-30 15:59:16'),
(64, 'tdesktop', 'ru', 306387, 2217, 2217, 0, '2022-08-30 15:59:16', '2022-08-30 15:59:16'),
(65, 'tdesktop', 'tr', 306347, 2217, 2217, 0, '2022-08-30 15:59:16', '2022-08-30 15:59:16'),
(66, 'tdesktop', 'uk', 309818, 2217, 2217, 0, '2022-08-30 15:59:16', '2022-08-30 15:59:16'),
(67, 'tdesktop', 'zh-hans-raw', 306049, 2217, 1803, 0, '2022-08-30 15:59:16', '2022-08-30 15:59:16'),
(68, 'tdesktop', 'zh-hant-raw', 306896, 2217, 2217, 0, '2022-08-30 15:59:16', '2022-08-30 15:59:16');

-- ----------------------------
-- Table structure for auths
-- ----------------------------
DROP TABLE IF EXISTS `auths`;
CREATE TABLE `auths` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `auth_key_id` bigint(20) NOT NULL,
  `layer` int(11) NOT NULL DEFAULT '0',
  `api_id` int(11) NOT NULL,
  `device_model` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `system_version` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `app_version` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `system_lang_code` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `lang_pack` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `lang_code` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `system_code` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `proxy` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `params` varchar(4096) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `client_ip` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `auth_key_id` (`auth_key_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of auths
-- ----------------------------

-- ----------------------------
-- Table structure for auth_keys
-- ----------------------------
DROP TABLE IF EXISTS `auth_keys`;
CREATE TABLE `auth_keys` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `auth_key_id` bigint(20) NOT NULL COMMENT 'auth_id',
  `auth_key` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `body` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'auth_key，原始数据为256的二进制数据，存储时转换成base64格式',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `auth_key_id` (`auth_key_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of auth_keys
-- ----------------------------

-- ----------------------------
-- Table structure for auth_op_logs
-- ----------------------------
DROP TABLE IF EXISTS `auth_op_logs`;
CREATE TABLE `auth_op_logs` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `auth_key_id` bigint(20) NOT NULL,
  `ip` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `op_type` int(11) NOT NULL DEFAULT '1',
  `log_text` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of auth_op_logs
-- ----------------------------

-- ----------------------------
-- Table structure for auth_seq_updates
-- ----------------------------
DROP TABLE IF EXISTS `auth_seq_updates`;
CREATE TABLE `auth_seq_updates` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `auth_id` bigint(20) NOT NULL,
  `user_id` int(11) NOT NULL,
  `seq` int(11) NOT NULL DEFAULT '0',
  `update_type` int(11) NOT NULL DEFAULT '0',
  `update_data` blob NOT NULL,
  `date2` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `auth_id` (`auth_id`,`user_id`,`seq`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of auth_seq_updates
-- ----------------------------

-- ----------------------------
-- Table structure for auth_users
-- ----------------------------
DROP TABLE IF EXISTS `auth_users`;
CREATE TABLE `auth_users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `auth_key_id` bigint(20) NOT NULL,
  `user_id` int(11) NOT NULL DEFAULT '0',
  `hash` bigint(20) NOT NULL DEFAULT '0',
  `layer` int(11) NOT NULL DEFAULT '0',
  `device_model` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `platform` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `system_version` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `api_id` int(11) NOT NULL DEFAULT '0',
  `app_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `app_version` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `date_created` int(11) NOT NULL DEFAULT '0',
  `date_actived` int(11) NOT NULL DEFAULT '0',
  `ip` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `country` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `region` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `auth_key_id` (`auth_key_id`,`user_id`),
  KEY `auth_key_id_2` (`auth_key_id`,`user_id`,`deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of auth_users
-- ----------------------------

-- ----------------------------
-- Table structure for banned
-- ----------------------------
DROP TABLE IF EXISTS `banned`;
CREATE TABLE `banned` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `phone` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `banned_time` bigint(20) NOT NULL,
  `expires` bigint(20) NOT NULL DEFAULT '0',
  `banned_reason` varchar(1024) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `log` varchar(1024) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `state` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of banned
-- ----------------------------

-- ----------------------------
-- Table structure for bots
-- ----------------------------
DROP TABLE IF EXISTS `bots`;
CREATE TABLE `bots` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `bot_id` int(11) NOT NULL,
  `bot_type` tinyint(4) NOT NULL DEFAULT '0',
  `creator_user_id` int(11) NOT NULL DEFAULT '0',
  `token` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `description` varchar(10240) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `bot_chat_history` tinyint(4) NOT NULL DEFAULT '0',
  `bot_nochats` tinyint(4) NOT NULL DEFAULT '1',
  `verified` tinyint(4) NOT NULL DEFAULT '0',
  `bot_inline_geo` tinyint(4) NOT NULL DEFAULT '0',
  `bot_info_version` int(11) NOT NULL DEFAULT '1',
  `bot_inline_placeholder` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `bot_id` (`bot_id`),
  UNIQUE KEY `token` (`token`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of bots
-- ----------------------------
INSERT INTO `bots` VALUES ('1', '6', '0', '0', 'aZYUYx2PO0vCQKS6eUY6u4IcpvoImZMtHr4', 'BotFather is the one bot to rule them all. Use it to create new bot accounts and manage your existing bots.\r\n\r\nAbout Telegram bots:\r\nhttps://core.telegram.org/bots\r\nBot API manual:\r\nhttps://core.telegram.org/bots/api\r\n\r\nContact @BotSupport if you have questions about the Bot API.', '0', '1', '1', '0', '1', '', '2022-08-30 15:17:52', '2022-08-30 09:50:10');
INSERT INTO `bots` VALUES ('6', '136817688', '0', '0', 'Ai8BUyhcOWburm0ODhwLQ9BldcatNWY8Laa', '', '0', '0', '0', '0', '1', '', '2022-08-30 10:14:26', '2022-08-30 22:50:45');
INSERT INTO `bots` VALUES ('7', '101', '0', '0', '101:YPT3TNGsBdGPyZk8cJP7CTQ6OHcvFh59lgm', 'This GIF search bot automatically works in all your chats and groups, no need to add it anywhere. Simply type @gif in any chat, then type your query (without hitting \'send\'). This will open a panel with GIF suggestions. Tap on a GIF to send it to your chat partner right away.', '0', '1', '0', '0', '1', 'Search GIFs', '2022-08-30 21:07:33', '2022-08-30 23:26:56');
INSERT INTO `bots` VALUES ('12', '103', '0', '0', 'GLKpl5YaCX86x1iIh1JVowYf5on7iDL4npd', 'This bot can help you find and share images. It works automatically, no need to add it anywhere. Simply open any of your chats and type @pic something in the message field. Then tap on a result to send.\\n\\nFor example, try typing @pic funny cat here.', '0', '1', '0', '0', '1', 'Search Images', '2022-08-30 21:07:33', '2022-08-30 09:14:06');

-- ----------------------------
-- Table structure for bot_commands
-- ----------------------------
DROP TABLE IF EXISTS `bot_commands`;
CREATE TABLE `bot_commands` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `bot_id` int(11) NOT NULL,
  `command` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `description` varchar(10240) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `bot_id` (`bot_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of bot_commands
-- ----------------------------

-- ----------------------------
-- Table structure for bot_updates
-- ----------------------------
DROP TABLE IF EXISTS `bot_updates`;
CREATE TABLE `bot_updates` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `bot_id` int(11) NOT NULL,
  `update_id` int(11) NOT NULL,
  `update_type` tinyint(11) NOT NULL,
  `update_data` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `date2` bigint(20) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `bot_id` (`bot_id`,`update_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of bot_updates
-- ----------------------------

-- ----------------------------
-- Table structure for channels
-- ----------------------------
DROP TABLE IF EXISTS `channels`;
CREATE TABLE `channels` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `creator_user_id` int(11) NOT NULL,
  `access_hash` bigint(20) NOT NULL,
  `secret_key_id` bigint(20) NOT NULL DEFAULT '0',
  `random_id` bigint(20) NOT NULL,
  `top_message` int(11) NOT NULL DEFAULT '0',
  `pinned_msg_id` int(11) NOT NULL DEFAULT '0',
  `read_outbox_max_id` int(11) NOT NULL DEFAULT '0',
  `date2` int(11) NOT NULL DEFAULT '0',
  `pts` int(11) NOT NULL DEFAULT '0',
  `participants_count` int(11) NOT NULL DEFAULT '0',
  `admins_count` int(11) NOT NULL DEFAULT '0',
  `kicked_count` int(11) NOT NULL DEFAULT '0',
  `banned_count` int(11) NOT NULL DEFAULT '0',
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `about` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `photo` varchar(5000) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `photo_id` bigint(20) NOT NULL DEFAULT '0',
  `public` tinyint(4) NOT NULL DEFAULT '0',
  `username` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `link` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `broadcast` tinyint(4) NOT NULL DEFAULT '0',
  `verified` tinyint(4) NOT NULL DEFAULT '0',
  `megagroup` tinyint(4) NOT NULL DEFAULT '0',
  `democracy` tinyint(4) NOT NULL DEFAULT '0',
  `signatures` tinyint(4) NOT NULL DEFAULT '0',
  `admins_enabled` tinyint(4) NOT NULL DEFAULT '0',
  `default_banned_rights` int(11) NOT NULL DEFAULT '0',
  `migrated_from_chat_id` int(11) NOT NULL DEFAULT '0',
  `pre_history_hidden` tinyint(4) NOT NULL DEFAULT '0',
  `has_link` tinyint(4) NOT NULL DEFAULT '0',
  `linked_chat_id` int(11) NOT NULL DEFAULT '0',
  `deactivated` tinyint(4) NOT NULL DEFAULT '0',
  `version` int(11) NOT NULL DEFAULT '1',
  `date` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `creator_user_id_3` (`creator_user_id`,`access_hash`)
) ENGINE=InnoDB AUTO_INCREMENT=1073741824 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of channels
-- ----------------------------

-- ----------------------------
-- Table structure for channel_admin_logs
-- ----------------------------
DROP TABLE IF EXISTS `channel_admin_logs`;
CREATE TABLE `channel_admin_logs` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `channel_id` int(11) NOT NULL,
  `event` int(11) NOT NULL,
  `event_data` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `query` varchar(5096) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `date2` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `channel_id` (`channel_id`),
  KEY `user_id` (`user_id`,`channel_id`,`event`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of channel_admin_logs
-- ----------------------------

-- ----------------------------
-- Table structure for channel_messages
-- ----------------------------
DROP TABLE IF EXISTS `channel_messages`;
CREATE TABLE `channel_messages` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `channel_id` int(11) NOT NULL,
  `channel_message_id` int(11) NOT NULL,
  `sender_user_id` int(11) NOT NULL,
  `random_id` bigint(20) NOT NULL,
  `pts` int(11) NOT NULL DEFAULT '0',
  `message_data_id` bigint(20) NOT NULL,
  `message_type` tinyint(4) NOT NULL DEFAULT '0',
  `message_data` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `message` varchar(4048) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `media_type` tinyint(4) NOT NULL DEFAULT '-1',
  `media_unread` tinyint(4) NOT NULL DEFAULT '0',
  `has_media_unread` tinyint(4) NOT NULL DEFAULT '0',
  `edit_message` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `edit_date` int(11) NOT NULL DEFAULT '0',
  `views` int(11) NOT NULL DEFAULT '1',
  `from_scheduled` int(11) NOT NULL DEFAULT '0',
  `date` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `sender_user_id` (`sender_user_id`,`random_id`),
  UNIQUE KEY `channel_id` (`channel_id`,`channel_message_id`),
  UNIQUE KEY `message_data_id` (`message_data_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of channel_messages
-- ----------------------------

-- ----------------------------
-- Table structure for channel_participants
-- ----------------------------
DROP TABLE IF EXISTS `channel_participants`;
CREATE TABLE `channel_participants` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `channel_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `is_creator` tinyint(4) NOT NULL DEFAULT '0',
  `is_pinned` tinyint(4) NOT NULL DEFAULT '0',
  `order_pinned` bigint(20) NOT NULL DEFAULT '0',
  `read_inbox_max_id` int(11) NOT NULL DEFAULT '0',
  `unread_count` int(11) NOT NULL DEFAULT '0',
  `unread_mentions_count` int(11) NOT NULL DEFAULT '0',
  `unread_mark` tinyint(4) NOT NULL DEFAULT '0',
  `draft_type` tinyint(4) NOT NULL DEFAULT '0',
  `draft_message_data` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `folder_id` int(11) NOT NULL DEFAULT '0',
  `folder_pinned` int(11) NOT NULL DEFAULT '0',
  `folder_order_pinned` bigint(20) NOT NULL DEFAULT '0',
  `inviter_user_id` int(11) NOT NULL DEFAULT '0',
  `promoted_by` int(11) NOT NULL DEFAULT '0',
  `admin_rights` int(11) NOT NULL DEFAULT '0',
  `hidden_prehistory` tinyint(4) NOT NULL DEFAULT '0',
  `hidden_prehistory_message_id` int(11) NOT NULL DEFAULT '0',
  `kicked_by` int(11) NOT NULL DEFAULT '0',
  `banned_rights` int(11) NOT NULL DEFAULT '0',
  `banned_until_date` int(11) NOT NULL DEFAULT '0',
  `migrated_from_max_id` int(11) NOT NULL DEFAULT '0',
  `available_min_id` int(11) NOT NULL DEFAULT '0',
  `available_min_pts` int(11) NOT NULL DEFAULT '0',
  `rank` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `has_scheduled` int(11) NOT NULL DEFAULT '0',
  `state` tinyint(4) NOT NULL DEFAULT '0',
  `date2` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `channel_id` (`channel_id`,`user_id`),
  KEY `chat_id` (`channel_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of channel_participants
-- ----------------------------

-- ----------------------------
-- Table structure for channel_pts_updates
-- ----------------------------
DROP TABLE IF EXISTS `channel_pts_updates`;
CREATE TABLE `channel_pts_updates` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `channel_id` int(11) NOT NULL,
  `pts` int(11) NOT NULL,
  `pts_count` int(11) NOT NULL,
  `update_type` tinyint(4) NOT NULL DEFAULT '0',
  `new_message_id` int(11) NOT NULL DEFAULT '0',
  `update_data` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `date2` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of channel_pts_updates
-- ----------------------------

-- ----------------------------
-- Table structure for channel_unread_mentions
-- ----------------------------
DROP TABLE IF EXISTS `channel_unread_mentions`;
CREATE TABLE `channel_unread_mentions` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `channel_id` int(11) NOT NULL,
  `mentioned_message_id` int(11) NOT NULL,
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of channel_unread_mentions
-- ----------------------------

-- ----------------------------
-- Table structure for channel_updates_state
-- ----------------------------
DROP TABLE IF EXISTS `channel_updates_state`;
CREATE TABLE `channel_updates_state` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `channel_id` int(11) NOT NULL,
  `pts` int(11) NOT NULL DEFAULT '0',
  `pts2` int(11) NOT NULL DEFAULT '0',
  `date` int(11) NOT NULL,
  `date2` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`,`channel_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of channel_updates_state
-- ----------------------------

-- ----------------------------
-- Table structure for chats
-- ----------------------------
DROP TABLE IF EXISTS `chats`;
CREATE TABLE `chats` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `creator_user_id` int(11) NOT NULL,
  `access_hash` bigint(20) NOT NULL,
  `random_id` bigint(20) NOT NULL,
  `participant_count` int(11) NOT NULL,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `about` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `link` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `photo` varchar(5000) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `chat_photo` varchar(4096) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `photo_id` bigint(20) NOT NULL DEFAULT '0',
  `admins_enabled` tinyint(4) NOT NULL DEFAULT '0',
  `default_banned_rights` int(11) NOT NULL DEFAULT '0',
  `migrated_to_id` int(11) NOT NULL DEFAULT '0',
  `migrated_to_access_hash` bigint(20) NOT NULL DEFAULT '0',
  `deactivated` tinyint(4) NOT NULL DEFAULT '0',
  `version` int(11) NOT NULL DEFAULT '1',
  `date` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of chats
-- ----------------------------

-- ----------------------------
-- Table structure for chat_participants
-- ----------------------------
DROP TABLE IF EXISTS `chat_participants`;
CREATE TABLE `chat_participants` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `chat_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `participant_type` tinyint(4) DEFAULT '0',
  `is_pinned` tinyint(4) NOT NULL DEFAULT '0',
  `order_pinned` bigint(20) NOT NULL DEFAULT '0',
  `top_message` int(11) NOT NULL DEFAULT '0',
  `pinned_msg_id` int(11) NOT NULL DEFAULT '0',
  `read_inbox_max_id` int(11) NOT NULL DEFAULT '0',
  `read_outbox_max_id` int(11) NOT NULL DEFAULT '0',
  `unread_count` int(11) NOT NULL DEFAULT '0',
  `unread_mentions_count` int(11) NOT NULL DEFAULT '0',
  `unread_mark` tinyint(4) NOT NULL DEFAULT '0',
  `draft_type` tinyint(4) NOT NULL DEFAULT '0',
  `draft_message_data` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `folder_id` int(11) NOT NULL DEFAULT '0',
  `folder_pinned` int(11) NOT NULL DEFAULT '0',
  `folder_order_pinned` bigint(20) NOT NULL DEFAULT '0',
  `inviter_user_id` int(11) NOT NULL DEFAULT '0',
  `invited_at` int(11) NOT NULL DEFAULT '0',
  `kicked_at` int(11) NOT NULL DEFAULT '0',
  `left_at` int(11) NOT NULL DEFAULT '0',
  `has_scheduled` int(11) NOT NULL DEFAULT '0',
  `state` tinyint(4) NOT NULL DEFAULT '0',
  `date2` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `chat_id_2` (`chat_id`,`user_id`),
  KEY `chat_id` (`chat_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of chat_participants
-- ----------------------------

-- ----------------------------
-- Table structure for conversations
-- ----------------------------
DROP TABLE IF EXISTS `conversations`;
CREATE TABLE `conversations` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `peer_id` int(11) NOT NULL,
  `is_pinned` tinyint(4) NOT NULL DEFAULT '0',
  `order_pinned` bigint(20) NOT NULL DEFAULT '0',
  `top_message` int(11) NOT NULL DEFAULT '0',
  `pinned_msg_id` int(11) NOT NULL DEFAULT '0',
  `read_inbox_max_id` int(11) NOT NULL DEFAULT '0',
  `read_outbox_max_id` int(11) NOT NULL DEFAULT '0',
  `unread_count` int(11) NOT NULL DEFAULT '0',
  `unread_mark` tinyint(4) NOT NULL DEFAULT '0',
  `draft_type` tinyint(4) NOT NULL DEFAULT '0',
  `draft_message_data` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `folder_id` int(11) NOT NULL DEFAULT '0',
  `folder_pinned` int(11) NOT NULL DEFAULT '0',
  `folder_order_pinned` bigint(20) NOT NULL DEFAULT '0',
  `has_scheduled` int(11) NOT NULL DEFAULT '0',
  `date2` int(11) NOT NULL,
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`,`peer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of conversations
-- ----------------------------

-- ----------------------------
-- Table structure for devices
-- ----------------------------
DROP TABLE IF EXISTS `devices`;
CREATE TABLE `devices` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `auth_key_id` bigint(20) NOT NULL,
  `user_id` int(11) NOT NULL,
  `token_type` tinyint(4) NOT NULL,
  `token` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL,
  `no_muted` tinyint(4) NOT NULL DEFAULT '0',
  `locked_period` int(11) NOT NULL DEFAULT '0',
  `app_sandbox` tinyint(4) NOT NULL DEFAULT '0',
  `secret` varchar(1024) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `other_uids` varchar(1024) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `state` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `auth_key_id` (`auth_key_id`,`user_id`,`token_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of devices
-- ----------------------------

-- ----------------------------
-- Table structure for dialog_filters
-- ----------------------------
DROP TABLE IF EXISTS `dialog_filters`;
CREATE TABLE `dialog_filters` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `dialog_filter_id` int(11) NOT NULL,
  `dialog_filter` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `order_value` bigint(20) NOT NULL DEFAULT '0',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of dialog_filters
-- ----------------------------

-- ----------------------------
-- Table structure for documents
-- ----------------------------
DROP TABLE IF EXISTS `documents`;
CREATE TABLE `documents` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `document_id` bigint(20) NOT NULL,
  `access_hash` bigint(20) NOT NULL,
  `dc_id` int(11) NOT NULL,
  `file_path` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `file_size` int(11) NOT NULL,
  `uploaded_file_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `ext` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `mime_type` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `thumb_id` bigint(20) NOT NULL DEFAULT '0',
  `version` int(11) NOT NULL DEFAULT '0',
  `attributes` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `document_id` (`document_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of documents
-- ----------------------------

-- ----------------------------
-- Table structure for encrypted_files
-- ----------------------------
DROP TABLE IF EXISTS `encrypted_files`;
CREATE TABLE `encrypted_files` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `encrypted_file_id` bigint(20) NOT NULL,
  `access_hash` bigint(20) NOT NULL,
  `dc_id` int(11) NOT NULL,
  `file_size` int(11) NOT NULL,
  `key_fingerprint` int(11) NOT NULL,
  `md5_checksum` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `file_path` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of encrypted_files
-- ----------------------------

-- ----------------------------
-- Table structure for files
-- ----------------------------
DROP TABLE IF EXISTS `files`;
CREATE TABLE `files` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `file_id` bigint(20) NOT NULL,
  `access_hash` bigint(20) NOT NULL,
  `creator_id` bigint(20) NOT NULL DEFAULT '0',
  `creator_user_id` int(11) NOT NULL DEFAULT '0',
  `file_part_id` bigint(20) NOT NULL DEFAULT '0',
  `file_parts` int(11) NOT NULL DEFAULT '0',
  `file_size` bigint(20) NOT NULL,
  `file_path` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `ext` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `is_big_file` tinyint(4) NOT NULL DEFAULT '0',
  `md5_checksum` char(33) COLLATE utf8mb4_unicode_ci NOT NULL,
  `upload_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of files
-- ----------------------------

-- ----------------------------
-- Table structure for file_parts
-- ----------------------------
DROP TABLE IF EXISTS `file_parts`;
CREATE TABLE `file_parts` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `creator_id` bigint(20) NOT NULL DEFAULT '0',
  `creator_user_id` int(11) NOT NULL DEFAULT '0',
  `file_id` bigint(20) NOT NULL DEFAULT '0',
  `file_part_id` bigint(20) NOT NULL,
  `file_part` int(11) NOT NULL DEFAULT '0',
  `is_big_file` tinyint(4) NOT NULL DEFAULT '0',
  `file_total_parts` int(11) NOT NULL DEFAULT '0',
  `file_path` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `file_size` bigint(20) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of file_parts
-- ----------------------------

-- ----------------------------
-- Table structure for folders
-- ----------------------------
DROP TABLE IF EXISTS `folders`;
CREATE TABLE `folders` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `photo` varchar(4096) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `autofill_new_broadcasts` tinyint(4) NOT NULL DEFAULT '0',
  `autofill_public_groups` tinyint(4) NOT NULL DEFAULT '0',
  `autofill_new_correspondents` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of folders
-- ----------------------------

-- ----------------------------
-- Table structure for giphy_datas
-- ----------------------------
DROP TABLE IF EXISTS `giphy_datas`;
CREATE TABLE `giphy_datas` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `giphy_id` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL,
  `document_id` bigint(20) NOT NULL,
  `photo_id` bigint(20) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `giphy_id` (`giphy_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of giphy_datas
-- ----------------------------

-- ----------------------------
-- Table structure for imported_contacts
-- ----------------------------
DROP TABLE IF EXISTS `imported_contacts`;
CREATE TABLE `imported_contacts` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `imported_user_id` int(11) NOT NULL,
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`),
  UNIQUE KEY `user_id_2` (`user_id`,`imported_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of imported_contacts
-- ----------------------------

-- ----------------------------
-- Table structure for languages
-- ----------------------------
DROP TABLE IF EXISTS `languages`;
CREATE TABLE `languages` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `lang_code` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `base_lang_code` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `link` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `official` tinyint(4) NOT NULL DEFAULT '0',
  `rtl` tinyint(4) NOT NULL DEFAULT '0',
  `beta` tinyint(4) NOT NULL DEFAULT '0',
  `name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `native_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `plural_code` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `translations_url` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `state` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `lang_code` (`lang_code`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of languages
-- ----------------------------
INSERT INTO `languages` VALUES ('1', 'en', '', 'en', '1', '0', '0', 'English', 'English', 'en', 'https://translations.telegram.org/en/', '0', '2022-08-30 10:22:08', '2022-08-30 12:03:59');
INSERT INTO `languages` VALUES ('2', 'ca', '', 'ca', '1', '0', '0', 'Catalan', 'Català', 'ca', 'https://translations.telegram.org/ca/', '0', '2022-08-30 10:22:08', '2022-08-30 20:59:23');
INSERT INTO `languages` VALUES ('3', 'nl', '', 'nl', '1', '0', '0', 'Dutch', 'Nederlands', 'nl', 'https://translations.telegram.org/nl/', '0', '2022-08-30 10:22:08', '2022-08-30 20:59:25');
INSERT INTO `languages` VALUES ('4', 'fr', '', 'fr', '1', '0', '0', 'French', 'Français', 'fr', 'https://translations.telegram.org/fr/', '0', '2022-08-30 10:22:08', '2022-08-30 20:59:26');
INSERT INTO `languages` VALUES ('5', 'de', '', 'de', '1', '0', '0', 'German', 'Deutsch', 'de', 'https://translations.telegram.org/de/', '0', '2022-08-30 10:22:08', '2022-08-30 20:59:27');
INSERT INTO `languages` VALUES ('6', 'it', '', 'it', '1', '0', '0', 'Italian', 'Italiano', 'it', 'https://translations.telegram.org/it/', '0', '2022-08-30 10:22:08', '2022-08-30 20:59:28');
INSERT INTO `languages` VALUES ('7', 'ko-raw', '', 'ko-beta', '0', '0', '1', 'Korean', '한국어', 'ko', 'https://translations.telegram.org/ko/', '0', '2022-08-30 10:22:08', '2022-08-30 20:59:30');
INSERT INTO `languages` VALUES ('8', 'ms', '', 'ms', '1', '0', '0', 'Malay', 'Bahasa Melayu', 'ms', 'https://translations.telegram.org/ms/', '0', '2022-08-30 10:22:08', '2022-08-30 20:59:31');
INSERT INTO `languages` VALUES ('9', 'pt-br', '', 'pt-br', '1', '0', '0', 'Portuguese (Brazil)', 'Português (Brasil)', 'pt', 'https://translations.telegram.org/pt-br/', '0', '2022-08-30 10:22:08', '2022-08-30 20:59:32');
INSERT INTO `languages` VALUES ('10', 'ru', '', 'ru', '1', '0', '0', 'Russian', 'Русский', 'ru', 'https://translations.telegram.org/ru/', '0', '2022-08-30 10:22:08', '2022-08-30 12:29:28');
INSERT INTO `languages` VALUES ('11', 'es', '', 'es', '1', '0', '0', 'Spanish', 'Español', 'es', 'https://translations.telegram.org/es/', '0', '2022-08-30 10:22:08', '2022-08-30 20:59:33');
INSERT INTO `languages` VALUES ('12', 'tr', '', 'tr', '1', '0', '0', 'Turkish', 'Türkçe', 'tr', 'https://translations.telegram.org/tr/', '0', '2022-08-30 10:22:08', '2022-08-30 20:59:34');
INSERT INTO `languages` VALUES ('13', 'ur', '', 'ur', '1', '0', '0', 'Ukrainian', 'Українська', 'ur', 'https://translations.telegram.org/ur/', '0', '2022-08-30 10:22:08', '2022-08-30 20:59:35');
INSERT INTO `languages` VALUES ('14', 'zh-hans-raw', '', 'zh-hans-beta', '0', '0', '1', 'Chinese (Simplified)', '简体中文 (beta)', 'zh', 'https://translations.telegram.org/zh-hans/', '2', '2022-08-30 10:22:08', '2022-08-30 23:57:32');
INSERT INTO `languages` VALUES ('15', 'classic-zh-cn', 'zh-hans-raw', 'classic-zh-cn', '0', '0', '0', 'Chinese (Simplified, @zh_CN)', '简体中文 (@zh_CN 版)', 'zh', 'https://translations.telegram.org/classic-zh-cn/', '0', '2022-08-30 10:22:08', '2022-08-30 12:01:43');
INSERT INTO `languages` VALUES ('17', 'ar-raw', '', 'ar-beta', '0', '1', '1', 'Arabic', 'العربية (beta)', 'ar', 'https://translations.telegram.org/ar/', '0', '2022-08-30 10:22:08', '2022-08-30 12:26:19');
INSERT INTO `languages` VALUES ('18', 'zh-hant-raw', '', 'zh-hant-beta', '0', '0', '1', 'Chinese (Traditional)', '繁體中文 (beta)', 'zh', 'https://translations.telegram.org/zh-hant/', '0', '2022-08-30 10:22:08', '2022-08-30 12:19:58');
INSERT INTO `languages` VALUES ('19', 'fa-raw', '', 'fa-beta', '0', '1', '1', 'Persian', 'فارسی (beta)', 'fa', 'https://translations.telegram.org/fa/', '0', '2022-08-30 10:22:08', '2022-08-30 12:24:48');

-- ----------------------------
-- Table structure for lang_pack_languages
-- ----------------------------
DROP TABLE IF EXISTS `lang_pack_languages`;
CREATE TABLE `lang_pack_languages` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `lang_pack` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `lang_code` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `version` int(11) NOT NULL,
  `base_lang_code` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `official` tinyint(4) NOT NULL DEFAULT '0',
  `rtl` tinyint(4) NOT NULL DEFAULT '0',
  `beta` tinyint(4) NOT NULL DEFAULT '0',
  `name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `native_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `plural_code` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `strings_count` int(11) NOT NULL DEFAULT '0',
  `translated_count` int(11) NOT NULL DEFAULT '0',
  `translations_url` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `state` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of lang_pack_languages
-- ----------------------------

-- ----------------------------
-- Table structure for lang_pack_strings
-- ----------------------------
DROP TABLE IF EXISTS `lang_pack_strings`;
CREATE TABLE `lang_pack_strings` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `lang_pack` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `lang_code` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `version` int(11) NOT NULL,
  `key2` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL,
  `pluralized` tinyint(4) NOT NULL DEFAULT '0',
  `value` varchar(1024) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `zero_value` varchar(1024) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `one_value` varchar(1024) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `two_value` varchar(1024) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `few_value` varchar(1024) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `many_value` varchar(1024) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `other_value` varchar(1024) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `lang_pack` (`lang_pack`,`lang_code`,`key2`),
  KEY `lang_pack_2` (`lang_pack`,`lang_code`),
  KEY `lang_pack_3` (`lang_pack`,`lang_code`,`version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of lang_pack_strings
-- ----------------------------

-- ----------------------------
-- Table structure for messages
-- ----------------------------
DROP TABLE IF EXISTS `messages`;
CREATE TABLE `messages` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `user_message_box_id` int(11) NOT NULL,
  `dialog_id` bigint(20) NOT NULL DEFAULT '0',
  `dialog_message_id` int(11) NOT NULL,
  `sender_user_id` int(11) NOT NULL,
  `peer_type` tinyint(4) NOT NULL,
  `peer_id` int(11) NOT NULL,
  `random_id` bigint(20) NOT NULL DEFAULT '0',
  `message_type` tinyint(4) NOT NULL DEFAULT '0',
  `message_data` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `message_data_id` bigint(20) NOT NULL,
  `message_data_type` int(11) NOT NULL DEFAULT '0',
  `message` varchar(6000) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `pts` int(11) NOT NULL DEFAULT '0',
  `pts_count` int(11) NOT NULL DEFAULT '1',
  `message_box_type` tinyint(4) NOT NULL,
  `reply_to_msg_id` int(11) NOT NULL DEFAULT '0',
  `mentioned` tinyint(4) NOT NULL DEFAULT '0',
  `media_unread` tinyint(4) NOT NULL DEFAULT '0',
  `has_media_unread` tinyint(4) NOT NULL DEFAULT '0',
  `from_scheduled` int(11) NOT NULL DEFAULT '0',
  `date2` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`,`message_data_id`),
  UNIQUE KEY `user_id_2` (`user_id`,`random_id`),
  KEY `user_id_3` (`user_id`,`user_message_box_id`),
  KEY `sender_user_id` (`sender_user_id`,`random_id`,`deleted`),
  KEY `message_data_id` (`message_data_id`,`deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of messages
-- ----------------------------

-- ----------------------------
-- Table structure for phones
-- ----------------------------
DROP TABLE IF EXISTS `phones`;
CREATE TABLE `phones` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `phone` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `region` varchar(8) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'CN',
  `region_code` varchar(8) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '86',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of phones
-- ----------------------------

-- ----------------------------
-- Table structure for phone_books
-- ----------------------------
DROP TABLE IF EXISTS `phone_books`;
CREATE TABLE `phone_books` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0',
  `auth_key_id` bigint(20) NOT NULL,
  `client_id` bigint(20) NOT NULL,
  `phone` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `first_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `last_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `auth_key_id` (`auth_key_id`,`client_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of phone_books
-- ----------------------------

-- ----------------------------
-- Table structure for phone_call_debugs
-- ----------------------------
DROP TABLE IF EXISTS `phone_call_debugs`;
CREATE TABLE `phone_call_debugs` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `call_id` bigint(20) NOT NULL,
  `participant_id` int(11) NOT NULL,
  `participant_auth_key_id` bigint(20) NOT NULL DEFAULT '0',
  `debug_data` varchar(4096) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `call_id` (`call_id`,`participant_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of phone_call_debugs
-- ----------------------------

-- ----------------------------
-- Table structure for phone_call_ratings
-- ----------------------------
DROP TABLE IF EXISTS `phone_call_ratings`;
CREATE TABLE `phone_call_ratings` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `call_id` bigint(20) NOT NULL,
  `participant_id` int(11) NOT NULL,
  `participant_auth_key_id` bigint(20) NOT NULL DEFAULT '0',
  `user_initiative` tinyint(4) NOT NULL DEFAULT '0',
  `rating` int(11) NOT NULL DEFAULT '0',
  `comment` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `call_id` (`call_id`,`participant_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of phone_call_ratings
-- ----------------------------

-- ----------------------------
-- Table structure for phone_call_sessions
-- ----------------------------
DROP TABLE IF EXISTS `phone_call_sessions`;
CREATE TABLE `phone_call_sessions` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `access_hash` bigint(20) NOT NULL,
  `admin_id` int(11) NOT NULL,
  `participant_id` int(11) NOT NULL,
  `admin_auth_key_id` bigint(20) NOT NULL DEFAULT '0',
  `participant_auth_key_id` bigint(20) NOT NULL DEFAULT '0',
  `random_id` bigint(20) NOT NULL,
  `admin_protocol` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `participant_protocol` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `g_a_hash` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `g_a` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `g_b` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `key_fingerprint` bigint(20) NOT NULL DEFAULT '0',
  `connections` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `admin_debug_data` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `participant_debug_data` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `admin_rating` int(11) NOT NULL DEFAULT '0',
  `admin_comment` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `participant_rating` int(11) NOT NULL DEFAULT '0',
  `participant_comment` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `date` int(11) NOT NULL,
  `state` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_auth_key_id` (`admin_auth_key_id`,`random_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of phone_call_sessions
-- ----------------------------

-- ----------------------------
-- Table structure for photos
-- ----------------------------
DROP TABLE IF EXISTS `photos`;
CREATE TABLE `photos` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `photo_id` int(11) NOT NULL,
  `has_stickers` int(11) NOT NULL DEFAULT '0',
  `access_hash` int(11) NOT NULL,
  `date` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of photos
-- ----------------------------

-- ----------------------------
-- Table structure for photo_datas
-- ----------------------------
DROP TABLE IF EXISTS `photo_datas`;
CREATE TABLE `photo_datas` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `photo_id` bigint(20) NOT NULL,
  `photo_type` tinyint(4) NOT NULL,
  `dc_id` int(11) NOT NULL,
  `volume_id` bigint(20) NOT NULL,
  `local_id` int(11) NOT NULL,
  `access_hash` bigint(20) NOT NULL,
  `width` int(11) NOT NULL,
  `height` int(11) NOT NULL,
  `file_size` int(11) NOT NULL DEFAULT '0',
  `file_path` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `ext` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `photo_id` (`photo_id`),
  KEY `dc_id` (`dc_id`,`volume_id`,`local_id`),
  KEY `photo_id_2` (`photo_id`,`local_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of photo_datas
-- ----------------------------

-- ----------------------------
-- Table structure for polls
-- ----------------------------
DROP TABLE IF EXISTS `polls`;
CREATE TABLE `polls` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `poll_id` bigint(20) NOT NULL DEFAULT '0',
  `creator` int(11) NOT NULL,
  `question` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `close_period` int(11) DEFAULT '0',
  `close_date` int(11) NOT NULL DEFAULT '0',
  `closed` tinyint(4) NOT NULL DEFAULT '0',
  `public_voters` tinyint(4) NOT NULL DEFAULT '0',
  `multiple_choice` tinyint(4) NOT NULL DEFAULT '0',
  `quiz` tinyint(4) NOT NULL DEFAULT '0',
  `correct_answer0` tinyint(4) NOT NULL DEFAULT '0',
  `correct_answer1` tinyint(4) NOT NULL DEFAULT '0',
  `correct_answer2` tinyint(4) NOT NULL DEFAULT '0',
  `correct_answer3` tinyint(4) NOT NULL DEFAULT '0',
  `correct_answer4` tinyint(4) NOT NULL DEFAULT '0',
  `correct_answer5` tinyint(4) NOT NULL DEFAULT '0',
  `correct_answer6` tinyint(4) NOT NULL DEFAULT '0',
  `correct_answer7` tinyint(4) NOT NULL DEFAULT '0',
  `correct_answer8` tinyint(4) NOT NULL DEFAULT '0',
  `correct_answer9` tinyint(4) NOT NULL DEFAULT '0',
  `text0` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL,
  `option0` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `voters0` int(11) NOT NULL DEFAULT '0',
  `text1` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `option1` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `voters1` int(11) NOT NULL DEFAULT '0',
  `text2` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `option2` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `voters2` int(11) NOT NULL DEFAULT '0',
  `text3` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `option3` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `voters3` int(11) NOT NULL DEFAULT '0',
  `text4` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `option4` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `voters4` int(11) NOT NULL DEFAULT '0',
  `text5` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `option5` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `voters5` int(11) NOT NULL DEFAULT '0',
  `text6` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `option6` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `voters6` int(11) NOT NULL DEFAULT '0',
  `text7` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `option7` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `voters7` int(11) NOT NULL DEFAULT '0',
  `text8` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `option8` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `voters8` int(11) NOT NULL DEFAULT '0',
  `text9` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `option9` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `voters9` int(11) NOT NULL DEFAULT '0',
  `solution` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `solution_entities` varchar(4096) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `date2` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of polls
-- ----------------------------

-- ----------------------------
-- Table structure for poll_answer_voters
-- ----------------------------
DROP TABLE IF EXISTS `poll_answer_voters`;
CREATE TABLE `poll_answer_voters` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `poll_id` bigint(20) NOT NULL,
  `vote_user_id` int(11) NOT NULL,
  `options` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `option0` tinyint(4) NOT NULL DEFAULT '0',
  `option1` tinyint(4) NOT NULL DEFAULT '0',
  `option2` tinyint(4) NOT NULL DEFAULT '0',
  `option3` tinyint(4) NOT NULL DEFAULT '0',
  `option4` tinyint(4) NOT NULL DEFAULT '0',
  `option5` tinyint(4) NOT NULL DEFAULT '0',
  `option6` tinyint(4) NOT NULL DEFAULT '0',
  `option7` tinyint(4) NOT NULL DEFAULT '0',
  `option8` tinyint(4) NOT NULL DEFAULT '0',
  `option9` tinyint(4) NOT NULL DEFAULT '0',
  `date2` bigint(20) NOT NULL DEFAULT '0',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `poll_id` (`poll_id`,`vote_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of poll_answer_voters
-- ----------------------------

-- ----------------------------
-- Table structure for popular_contacts
-- ----------------------------
DROP TABLE IF EXISTS `popular_contacts`;
CREATE TABLE `popular_contacts` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `phone` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `importers` int(11) NOT NULL DEFAULT '1',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of popular_contacts
-- ----------------------------

-- ----------------------------
-- Table structure for predefined_users
-- ----------------------------
DROP TABLE IF EXISTS `predefined_users`;
CREATE TABLE `predefined_users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `phone` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `first_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `last_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `username` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `code` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `verified` tinyint(4) NOT NULL DEFAULT '0',
  `registered_user_id` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of predefined_users
-- ----------------------------

-- ----------------------------
-- Table structure for reports
-- ----------------------------
DROP TABLE IF EXISTS `reports`;
CREATE TABLE `reports` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `report_type` tinyint(4) NOT NULL DEFAULT '0',
  `peer_type` tinyint(4) NOT NULL DEFAULT '0',
  `peer_id` int(11) NOT NULL DEFAULT '0',
  `message_sender_user_id` int(11) NOT NULL DEFAULT '0',
  `message_id` int(11) NOT NULL DEFAULT '0',
  `reason` tinyint(4) NOT NULL DEFAULT '0',
  `text` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of reports
-- ----------------------------

-- ----------------------------
-- Table structure for saved_gifs
-- ----------------------------
DROP TABLE IF EXISTS `saved_gifs`;
CREATE TABLE `saved_gifs` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `gif_id` bigint(20) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of saved_gifs
-- ----------------------------

-- ----------------------------
-- Table structure for scheduled_messages
-- ----------------------------
DROP TABLE IF EXISTS `scheduled_messages`;
CREATE TABLE `scheduled_messages` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `user_message_box_id` int(11) NOT NULL,
  `peer_type` tinyint(4) NOT NULL,
  `peer_id` int(11) NOT NULL,
  `dialog_id` bigint(20) NOT NULL,
  `random_id` bigint(20) NOT NULL DEFAULT '0',
  `message_type` tinyint(4) NOT NULL DEFAULT '0',
  `message_data_type` tinyint(4) NOT NULL DEFAULT '0',
  `message_data` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `message_box_type` tinyint(4) NOT NULL,
  `scheduled_date` int(11) NOT NULL,
  `date2` int(11) NOT NULL DEFAULT '0',
  `state` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id_2` (`user_id`,`random_id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of scheduled_messages
-- ----------------------------

-- ----------------------------
-- Table structure for secret_chats
-- ----------------------------
DROP TABLE IF EXISTS `secret_chats`;
CREATE TABLE `secret_chats` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `access_hash` bigint(20) NOT NULL,
  `admin_id` int(11) NOT NULL,
  `participant_id` int(11) NOT NULL,
  `admin_auth_key_id` bigint(20) NOT NULL DEFAULT '0',
  `participant_auth_key_id` bigint(20) NOT NULL DEFAULT '0',
  `random_id` int(11) NOT NULL,
  `g_a` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `g_b` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `key_fingerprint` bigint(20) NOT NULL DEFAULT '0',
  `state` tinyint(4) NOT NULL DEFAULT '0',
  `date` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of secret_chats
-- ----------------------------

-- ----------------------------
-- Table structure for secret_chat_messages
-- ----------------------------
DROP TABLE IF EXISTS `secret_chat_messages`;
CREATE TABLE `secret_chat_messages` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `sender_user_id` int(11) NOT NULL,
  `chat_id` int(11) NOT NULL,
  `peer_id` int(11) NOT NULL,
  `random_id` bigint(20) NOT NULL,
  `message_type` tinyint(4) NOT NULL DEFAULT '0',
  `message_data` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `date2` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of secret_chat_messages
-- ----------------------------

-- ----------------------------
-- Table structure for secret_chat_qts_updates
-- ----------------------------
DROP TABLE IF EXISTS `secret_chat_qts_updates`;
CREATE TABLE `secret_chat_qts_updates` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `auth_key_id` bigint(20) NOT NULL,
  `chat_id` int(11) NOT NULL,
  `qts` int(11) NOT NULL,
  `chat_message_id` bigint(20) NOT NULL,
  `date2` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of secret_chat_qts_updates
-- ----------------------------

-- ----------------------------
-- Table structure for sticker_packs
-- ----------------------------
DROP TABLE IF EXISTS `sticker_packs`;
CREATE TABLE `sticker_packs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sticker_set_id` bigint(20) NOT NULL,
  `emoticon` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `document_id` bigint(20) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of sticker_packs
-- ----------------------------

-- ----------------------------
-- Table structure for sticker_sets
-- ----------------------------
DROP TABLE IF EXISTS `sticker_sets`;
CREATE TABLE `sticker_sets` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sticker_set_id` bigint(20) NOT NULL,
  `access_hash` bigint(20) NOT NULL,
  `title` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `short_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `count` int(11) NOT NULL DEFAULT '0',
  `hash` int(11) NOT NULL DEFAULT '0',
  `official` tinyint(4) NOT NULL DEFAULT '0',
  `mask` tinyint(4) NOT NULL DEFAULT '0',
  `masks` tinyint(4) NOT NULL DEFAULT '0',
  `archived` tinyint(4) NOT NULL DEFAULT '0',
  `animated` tinyint(4) NOT NULL DEFAULT '0',
  `installed_date` int(11) NOT NULL DEFAULT '0',
  `thumb` varchar(1000) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `thumb_dc_id` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `sticker_set_id` (`sticker_set_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of sticker_sets
-- ----------------------------

-- ----------------------------
-- Table structure for unregistered_contacts
-- ----------------------------
DROP TABLE IF EXISTS `unregistered_contacts`;
CREATE TABLE `unregistered_contacts` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `phone` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `importer_user_id` int(11) NOT NULL,
  `import_first_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `import_last_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `imported` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`,`importer_user_id`),
  KEY `phone_2` (`phone`,`importer_user_id`,`imported`),
  KEY `phone_3` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of unregistered_contacts
-- ----------------------------

-- ----------------------------
-- Table structure for username
-- ----------------------------
DROP TABLE IF EXISTS `username`;
CREATE TABLE `username` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `peer_type` tinyint(4) NOT NULL DEFAULT '0',
  `peer_id` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=149 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of username
-- ----------------------------
INSERT INTO `username` VALUES ('135', 'BotFather', '2', '6', '0', '2022-08-30 22:24:52', '2022-08-30 22:24:52');
INSERT INTO `username` VALUES ('141', 'gif', '2', '101', '0', '2022-08-30 10:14:26', '2022-08-30 00:39:15');
INSERT INTO `username` VALUES ('143', 'Channel_Bot', '2', '136817688', '0', '2022-08-30 10:14:26', '2022-08-30 09:12:07');
INSERT INTO `username` VALUES ('148', 'pic', '2', '103', '0', '2022-08-30 10:14:26', '2022-08-30 00:39:15');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_type` tinyint(4) NOT NULL DEFAULT '0',
  `access_hash` bigint(20) NOT NULL,
  `secret_key_id` bigint(20) NOT NULL DEFAULT '0',
  `first_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `last_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `username` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `phone` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `country_code` varchar(3) COLLATE utf8mb4_unicode_ci NOT NULL,
  `verified` tinyint(4) NOT NULL DEFAULT '0',
  `about` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `state` int(11) NOT NULL DEFAULT '0',
  `is_bot` tinyint(1) NOT NULL DEFAULT '0',
  `account_days_ttl` int(11) NOT NULL DEFAULT '180',
  `photo` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `profile_photo` varchar(4096) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `photos` varchar(5012) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `min` tinyint(4) NOT NULL DEFAULT '0',
  `restricted` tinyint(4) NOT NULL DEFAULT '0',
  `restriction_reason` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `delete_reason` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`),
  KEY `id` (`id`,`deleted`)
) ENGINE=InnoDB AUTO_INCREMENT=136817689 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES ('6', '2', '6599886787491911851', '0', 'BotFather', '', 'BotFather', '', '', '1', 'BotFather is the one bot to rule them all. Use it to create new bot accounts and manage your existing bots.', '0', '1', '180', '', '', '', '0', '0', '', '0', '', '2022-08-30 21:43:11', '2022-08-30 22:03:14');
INSERT INTO `users` VALUES ('101', '2', '5577006791947779410', '0', 'Giphy GIF Search', '', 'gif', 'gif', '', '1', '', '0', '1', '180', '', '', '', '0', '0', '', '0', '', '2022-08-30 21:43:11', '2022-08-30 23:50:42');
INSERT INTO `users` VALUES ('103', '2', '5577006791947779411', '0', 'Yandex Image Search', '', 'pic', 'pic', '', '0', '', '0', '1', '180', '', '', '', '0', '0', '', '0', '', '2022-08-30 21:43:11', '2022-08-30 09:09:27');
INSERT INTO `users` VALUES ('424000', '1', '6599886787491911852', '0', 'Volunteer Support', '', '', '424000', 'CN', '1', '', '0', '0', '180', '', '', '', '0', '0', '', '0', '', '2022-08-30 21:43:11', '2022-08-30 16:26:32');
INSERT INTO `users` VALUES ('777000', '1', '6599886787491911851', '6895602324158323006', '系统公告', '', 'chat', '42777', '', '1', '', '0', '0', '180', '', '', '', '0', '0', '', '0', '', '2022-08-30 21:43:11', '2022-08-30 16:23:28');
INSERT INTO `users` VALUES ('136817688', '2', '3944724221060850288', '117013984796508050', 'Channel', '', 'Channel_Bot', 'Channel_Bot', 'CN', '0', '', '0', '1', '180', '', '', '', '0', '0', '', '0', '', '2022-08-30 10:14:26', '2022-08-30 22:19:02');

-- ----------------------------
-- Table structure for user_blocks
-- ----------------------------
DROP TABLE IF EXISTS `user_blocks`;
CREATE TABLE `user_blocks` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `block_id` int(11) NOT NULL DEFAULT '0',
  `date` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`,`block_id`),
  KEY `user_id_2` (`user_id`,`block_id`,`deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of user_blocks
-- ----------------------------

-- ----------------------------
-- Table structure for user_contacts
-- ----------------------------
DROP TABLE IF EXISTS `user_contacts`;
CREATE TABLE `user_contacts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `owner_user_id` int(11) NOT NULL,
  `contact_user_id` int(11) NOT NULL,
  `contact_phone` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `contact_first_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `contact_last_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `mutual` tinyint(4) NOT NULL DEFAULT '0',
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  `date2` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `owner_user_id` (`owner_user_id`,`contact_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of user_contacts
-- ----------------------------

-- ----------------------------
-- Table structure for user_notify_settings
-- ----------------------------
DROP TABLE IF EXISTS `user_notify_settings`;
CREATE TABLE `user_notify_settings` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `peer_type` tinyint(4) NOT NULL,
  `peer_id` int(11) NOT NULL,
  `show_previews` tinyint(4) NOT NULL DEFAULT '-1',
  `silent` tinyint(4) NOT NULL DEFAULT '-1',
  `mute_until` int(11) NOT NULL DEFAULT '-1',
  `sound` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'default',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`,`peer_type`,`peer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of user_notify_settings
-- ----------------------------

-- ----------------------------
-- Table structure for user_passwords
-- ----------------------------
DROP TABLE IF EXISTS `user_passwords`;
CREATE TABLE `user_passwords` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `new_algo_salt1` varchar(80) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `v` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `srp_id` bigint(20) NOT NULL DEFAULT '0',
  `srp_b` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `B` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `hint` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `email` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `has_recovery` tinyint(4) NOT NULL DEFAULT '0',
  `code` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `code_expired` int(11) NOT NULL DEFAULT '0',
  `attempts` int(11) NOT NULL DEFAULT '0',
  `state` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of user_passwords
-- ----------------------------

-- ----------------------------
-- Table structure for user_peer_settings
-- ----------------------------
DROP TABLE IF EXISTS `user_peer_settings`;
CREATE TABLE `user_peer_settings` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `peer_type` tinyint(4) NOT NULL,
  `peer_id` int(11) NOT NULL,
  `hide` tinyint(4) NOT NULL DEFAULT '0',
  `report_spam` tinyint(4) NOT NULL DEFAULT '0',
  `add_contact` tinyint(4) NOT NULL DEFAULT '0',
  `block_contact` tinyint(4) NOT NULL DEFAULT '0',
  `share_contact` tinyint(4) NOT NULL DEFAULT '0',
  `need_contacts_exception` tinyint(4) NOT NULL DEFAULT '0',
  `report_geo` tinyint(4) NOT NULL DEFAULT '0',
  `autoarchived` tinyint(4) NOT NULL DEFAULT '0',
  `geo_distance` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`,`peer_type`,`peer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of user_peer_settings
-- ----------------------------

-- ----------------------------
-- Table structure for user_presences
-- ----------------------------
DROP TABLE IF EXISTS `user_presences`;
CREATE TABLE `user_presences` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `last_seen_at` bigint(20) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of user_presences
-- ----------------------------

-- ----------------------------
-- Table structure for user_privacies
-- ----------------------------
DROP TABLE IF EXISTS `user_privacies`;
CREATE TABLE `user_privacies` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `key_type` tinyint(4) NOT NULL DEFAULT '0',
  `rules` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`,`key_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of user_privacies
-- ----------------------------

-- ----------------------------
-- Table structure for user_profile_photos
-- ----------------------------
DROP TABLE IF EXISTS `user_profile_photos`;
CREATE TABLE `user_profile_photos` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `photo_id` bigint(20) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of user_profile_photos
-- ----------------------------

-- ----------------------------
-- Table structure for user_pts_updates
-- ----------------------------
DROP TABLE IF EXISTS `user_pts_updates`;
CREATE TABLE `user_pts_updates` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `pts` int(11) NOT NULL,
  `pts_count` int(11) NOT NULL,
  `update_type` tinyint(4) NOT NULL DEFAULT '0',
  `update_data` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `date2` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`,`pts`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of user_pts_updates
-- ----------------------------

-- ----------------------------
-- Table structure for user_qts_updates
-- ----------------------------
DROP TABLE IF EXISTS `user_qts_updates`;
CREATE TABLE `user_qts_updates` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `qts` int(11) NOT NULL,
  `update_type` tinyint(4) NOT NULL DEFAULT '0',
  `update_data` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `date2` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`,`qts`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of user_qts_updates
-- ----------------------------

-- ----------------------------
-- Table structure for user_settings
-- ----------------------------
DROP TABLE IF EXISTS `user_settings`;
CREATE TABLE `user_settings` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `key2` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `value` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL,
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`,`key2`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of user_settings
-- ----------------------------

-- ----------------------------
-- Table structure for user_sticker_sets
-- ----------------------------
DROP TABLE IF EXISTS `user_sticker_sets`;
CREATE TABLE `user_sticker_sets` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0',
  `sticker_set_id` bigint(20) NOT NULL DEFAULT '0',
  `archived` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq` (`user_id`,`sticker_set_id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of user_sticker_sets
-- ----------------------------

-- ----------------------------
-- Table structure for user_term_of_services
-- ----------------------------
DROP TABLE IF EXISTS `user_term_of_services`;
CREATE TABLE `user_term_of_services` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `term_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `last_version` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of user_term_of_services
-- ----------------------------

-- ----------------------------
-- Table structure for wall_papers
-- ----------------------------
DROP TABLE IF EXISTS `wall_papers`;
CREATE TABLE `wall_papers` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` tinyint(4) NOT NULL DEFAULT '0',
  `title` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `color` int(11) NOT NULL DEFAULT '0',
  `bg_color` int(11) NOT NULL DEFAULT '0',
  `photo_id` bigint(20) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of wall_papers
-- ----------------------------
