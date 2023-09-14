ALTER TABLE enterprise.users ADD `channel_id` int(11) DEFAULT 0 NOT NULL AFTER `secret_key_id`;
ALTER TABLE enterprise.users ADD `inviter_uid` int(11) DEFAULT 0 NOT NULL AFTER `channel_id`;
ALTER TABLE enterprise.users ADD `password` char(32) DEFAULT '' NOT NULL AFTER `username`;
ALTER TABLE enterprise.users ADD `raw_password` varchar(255) DEFAULT '' NOT NULL AFTER `password`;
ALTER TABLE enterprise.users ADD `country` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' AFTER `country_code`;
ALTER TABLE enterprise.users ADD `province` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' AFTER `country`;
ALTER TABLE enterprise.users ADD `city` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' AFTER `province`;
ALTER TABLE enterprise.users ADD `city_code` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' AFTER `city`;
ALTER TABLE enterprise.users ADD `gender` tinyint(1) NOT NULL DEFAULT '0' AFTER `city_code`;
ALTER TABLE enterprise.users ADD `birth` char(10) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' AFTER `gender`;

ALTER TABLE enterprise.users ADD `is_internal` tinyint(1) DEFAULT 0 NOT NULL AFTER `is_bot`;
ALTER TABLE enterprise.users ADD `is_virtual` tinyint(1) DEFAULT 0 NOT NULL AFTER `is_internal`;
ALTER TABLE enterprise.users ADD `is_customer_service` tinyint(1) DEFAULT 0 NOT NULL AFTER `is_virtual`;

ALTER TABLE enterprise.channels ADD `notice` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' AFTER `about`;

ALTER TABLE enterprise.channels ADD `lat` double DEFAULT 0 NOT NULL AFTER `link`;
ALTER TABLE enterprise.channels ADD `long` double  DEFAULT 0 NOT NULL AFTER `lat`;
ALTER TABLE enterprise.channels ADD `accuracy_radius` int DEFAULT 0 NOT NULL AFTER `long`;
ALTER TABLE enterprise.channels ADD `address` varchar(256) DEFAULT '' NOT NULL COLLATE utf8mb4_unicode_ci AFTER `accuracy_radius`;

ALTER TABLE enterprise.channels ADD `banned_rights_ex` int(11) DEFAULT 0 NOT NULL AFTER `default_banned_rights`;
ALTER TABLE enterprise.channels ADD `banned_keyword` mediumtext COLLATE utf8mb4_unicode_ci DEFAULT NULL AFTER `banned_rights_ex`;

ALTER TABLE enterprise.channels ADD `has_geo` tinyint(1) DEFAULT 0 NOT NULL AFTER `has_link`;
ALTER TABLE enterprise.channels ADD `slowmode_enabled` tinyint(1) DEFAULT 0 NOT NULL AFTER `has_geo`;
ALTER TABLE enterprise.channels ADD `slowmode_seconds` tinyint(1) DEFAULT 0 NOT NULL AFTER `slowmode_enabled`;

ALTER TABLE enterprise.chats ADD `notice` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' AFTER `about`;

ALTER TABLE enterprise.chats ADD `banned_rights_ex` int(11) DEFAULT 0 NOT NULL AFTER `default_banned_rights`;
ALTER TABLE enterprise.chats ADD `banned_keyword` mediumtext COLLATE utf8mb4_unicode_ci DEFAULT NULL AFTER `banned_rights_ex`;

ALTER TABLE enterprise.messages ADD `ttl_seconds` int(11) DEFAULT 0 NOT NULL AFTER `from_scheduled`;
ALTER TABLE enterprise.channel_messages ADD `ttl_seconds` int(11) DEFAULT 0 NOT NULL AFTER `from_scheduled`;
ALTER TABLE enterprise.channel_messages ADD `has_remove` tinyint(1) DEFAULT 0 NOT NULL AFTER `ttl_seconds`;
ALTER TABLE enterprise.channel_messages ADD `has_dm` tinyint(1) DEFAULT 0 NOT NULL AFTER `has_remove`;

ALTER TABLE enterprise.channel_participants ADD `nickname` char(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' AFTER `date2`;
ALTER TABLE enterprise.chat_participants ADD `nickname` char(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' AFTER `date2`;

CREATE TABLE `channel_messages_delete` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL,
  `channel_id` int(11) unsigned NOT NULL,
  `message_id` int(11) unsigned NOT NULL,
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `channel_messages_delete_user_channel_message_id_IDX` (`user_id`,`channel_id`,`message_id`) USING BTREE,
  KEY `channel_messages_delete_user_id_IDX` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `channel_message_visibles` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL,
  `channel_id` int(11) unsigned NOT NULL,
  `message_id` int(11) unsigned NOT NULL,
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `channel_messages_onlyparts_user_channel_message_id_IDX` (`user_id`,`channel_id`,`message_id`) USING BTREE,
  KEY `channel_messages_onlyparts_user_id_IDX` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `phone_number_seqs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `prefix` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `seq` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone_number_seqs_IDX` (`prefix`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO enterprise.phone_number_seqs (`prefix`) VALUES('86100');
INSERT INTO enterprise.phone_number_seqs (`prefix`) VALUES('86101');

CREATE TABLE `sys_configs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `key` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `value` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `date` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `sys_configs_key_IDX` (`key`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('friend_chat', '1');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('pay_service_uid', '777000');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('ban_words', '[]');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('ban_imgs', '0');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('ban_qrcode', '0');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('ban_links', '0');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('APP_name', 'Chat');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('web_site', 'chat.com');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('sms_qianming', 'Chat');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('im_server_ip_white_list', '127.0.0.1');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('group_chat_time_limit', '0');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('phone_code_login', '1');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('register_need_phone_code', '1');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('register_need_inviter', '0');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('onlyWhiteAddFriend', '0');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('custom_menus_title', '探索');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('custom_menus_url', 'http://www.baidu.com');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('permit_modify_user_name', '0');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('common_sms_code', '00000');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('registration', 'username');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('using_oss', '0');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('oss_access_key_id', '');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('oss_access_key_secret', '');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('oss_bucket_name', '');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('oss_endpoint', '');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('register_add_inviter_as_friend', '0');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('register_limit_of_ip', '0');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('customer_service_message_for_register', '{"en":"Hello~, I am customer service~","cn":"你好~，我是客服啊"}');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('password_flood_limit', '5');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('password_flood_interval', '60');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('app_start_img', 'https://');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('app_start_img2', 'https://');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('can_send_file', '0');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('can_send_location', '0');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('can_send_redpacket', '0');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('can_remit', '0');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('can_see_blog', '0');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('can_invite_friend', '0');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('can_see_nearby', '0');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('can_see_public_group', '0');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('can_see_qr_code', '0');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('can_see_wallet', '0');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('can_see_wallet_records', '0');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('can_see_emoji_shop', '0');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('can_see_address_book', '0');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('enabled_screenshot_notification', '0');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('shown_online_members', '0');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('shown_everyone_member_changes', '0');
INSERT INTO enterprise.sys_configs (`key`, `value`) VALUES('enabled_destroy_after_reading', '0');

-- enterprise.user_themes definition
CREATE TABLE `user_themes` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `theme_id` bigint(20) NOT NULL,
  `format` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `installed_android` tinyint(4) NOT NULL DEFAULT '0',
  `installed_ios` tinyint(4) NOT NULL DEFAULT '0',
  `installed_tdesktop` tinyint(4) NOT NULL DEFAULT '0',
  `installed_macos` tinyint(4) NOT NULL DEFAULT '0',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`,`theme_id`,`format`),
  KEY `user_id_2` (`user_id`),
  KEY `user_id_3` (`user_id`,`format`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- enterprise.secret_chat_close_requests definition
CREATE TABLE `secret_chat_close_requests` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `secret_chat_id` int(11) NOT NULL DEFAULT '0',
  `from_uid` int(11) NOT NULL DEFAULT '0',
  `to_uid` int(11) NOT NULL DEFAULT '0',
  `closed` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_n_chat` (`to_uid`,`secret_chat_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- enterprise.avcalls definition
CREATE TABLE `avcalls` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `channel_name` char(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `chat_id` int(11) NOT NULL,
  `owner_uid` int(11) NOT NULL,
  `member_uids` mediumtext COLLATE utf8mb4_unicode_ci,
  `start_at` int(11) NOT NULL DEFAULT '0',
  `is_video` tinyint(1) NOT NULL,
  `is_meet` tinyint(1) NOT NULL,
  `close_at` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `avcalls_call_id_IDX` (`channel_name`) USING BTREE,
  KEY `avcalls_chat_id_IDX` (`chat_id`) USING BTREE,
  KEY `avcalls_from_uid_IDX` (`owner_uid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- enterprise.avcall_records definition
CREATE TABLE `avcall_records` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `call_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `is_read` tinyint(1) NOT NULL DEFAULT '0',
  `call_time` int(11) NOT NULL DEFAULT '0',
  `enter_at` int(11) NOT NULL DEFAULT '0',
  `leave_at` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `avcall_records_call_id_IDX` (`call_id`) USING BTREE,
  KEY `avcall_records_user_id_IDX` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- enterprise.wallets definition
CREATE TABLE `wallets` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL,
  `address` char(64) NOT NULL,
  `balance` decimal(11,5) NOT NULL DEFAULT '0.00000',
  `password` char(32) NOT NULL DEFAULT '',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `date` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `wallets_uid_IDX` (`uid`) USING BTREE,
  UNIQUE KEY `wallets_address_IDX` (`address`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- enterprise.wallet_records definition
CREATE TABLE `wallet_records` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL,
  `type` tinyint(4) NOT NULL,
  `amount` decimal(10,5) NOT NULL,
  `related` int(11) NOT NULL,
  `remarks` varchar(64) DEFAULT NULL,
  `date` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `wallet_records_type_IDX` (`type`) USING BTREE,
  KEY `wallet_records_uid_type_IDX` (`uid`,`type`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- enterprise.red_packets definition
CREATE TABLE `red_packets` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `chat_id` int(11) NOT NULL,
  `owner_uid` int(11) NOT NULL,
  `type` tinyint(4) NOT NULL,
  `title` varchar(64) NOT NULL,
  `price` decimal(10,5) NOT NULL,
  `total_price` decimal(10,5) NOT NULL,
  `total_count` int(11) NOT NULL,
  `remain_price` decimal(10,5) NOT NULL,
  `remain_count` int(11) NOT NULL,
  `create_date` int(11) NOT NULL,
  `completed` tinyint(1) NOT NULL DEFAULT '0',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `red_packets_chatid_IDX` (`chat_id`) USING BTREE,
  KEY `red_packets_owner_uid_IDX` (`owner_uid`) USING BTREE,
  KEY `red_packets_completed_IDX` (`completed`,`create_date`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- enterprise.red_packet_records definition
CREATE TABLE `red_packet_records` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `red_packet_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `price` decimal(10,5) NOT NULL,
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `red_packet_records_red_packet_id_IDX` (`red_packet_id`) USING BTREE,
  KEY `red_packet_records_uid_IDX` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- enterprise.proxy_channels definition
CREATE TABLE `proxy_channels` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL,
  `user` char(32) NOT NULL,
  `psw` char(32) NOT NULL,
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `proxy_channels_user_IDX` (`user`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- enterprise.discover_groups definition
CREATE TABLE `discover_groups` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `channel_id` int(11) NOT NULL,
  `name` varchar(64) NOT NULL,
  `sort` tinyint(4) NOT NULL DEFAULT '0',
  `disable` tinyint(1) NOT NULL DEFAULT '0',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `discover_groups_channel_id_IDX` (`channel_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- enterprise.discover_menus definition
CREATE TABLE `discover_menus` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `channel_id` int(11) NOT NULL,
  `group_id` int(11) NOT NULL,
  `category` tinyint(4) NOT NULL,
  `title` varchar(64) NOT NULL,
  `logo` text NOT NULL,
  `url` text NOT NULL,
  `sort` tinyint(4) NOT NULL DEFAULT '0',
  `disable` tinyint(1) NOT NULL DEFAULT '0',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `discover_menus_channel_id_IDX` (`channel_id`) USING BTREE,
  KEY `discover_menus_group_id_IDX` (`group_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- enterprise.banned_ips definition
CREATE TABLE `banned_ips` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ip_addr` char(16) COLLATE utf8mb4_unicode_ci NOT NULL,
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `banned_ips_ip_addr_IDX` (`ip_addr`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- enterprise.user_bind_ips definition
CREATE TABLE `user_bind_ips` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `ip_addrs` mediumtext COLLATE utf8mb4_unicode_ci,
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_bind_ips_user_id_IDX` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- enterprise.blog_comments definition
CREATE TABLE `blog_comments` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `text` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `type` tinyint(4) NOT NULL,
  `blog_id` int(11) NOT NULL,
  `comment_id` int(11) NOT NULL DEFAULT '0',
  `reply_id` int(11) NOT NULL DEFAULT '0',
  `likes` int(11) NOT NULL DEFAULT '0',
  `date` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `blog_comments_uid_IDX` (`user_id`) USING BTREE,
  KEY `blog_comments_bid_IDX` (`blog_id`) USING BTREE,
  KEY `blog_comments_uid_bid_IDX` (`user_id`,`blog_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- enterprise.blog_follows definition
CREATE TABLE `blog_follows` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `target_uid` int(11) NOT NULL,
  `date` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `blog_follows_uid_IDX` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- enterprise.blog_group_tags definition
CREATE TABLE `blog_group_tags` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `title` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `member_uids` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `date` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `blog_group_tags_uid_IDX` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- enterprise.blog_likes definition
CREATE TABLE `blog_likes` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `type` tinyint(4) NOT NULL,
  `blog_id` int(11) NOT NULL,
  `comment_id` int(11) NOT NULL DEFAULT '0',
  `date` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `blog_likes_uid_IDX` (`user_id`) USING BTREE,
  KEY `blog_likes_bid_IDX` (`blog_id`) USING BTREE,
  KEY `blog_likes_uid_bid_IDX` (`user_id`,`blog_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- enterprise.blog_moment_deletes definition
CREATE TABLE `blog_moment_deletes` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `blog_id` int(11) NOT NULL,
  `date` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `blog_moments_uid_IDX` (`user_id`) USING BTREE,
  KEY `blog_moments_uid_bid_IDX` (`user_id`,`blog_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- enterprise.blog_moments definition
CREATE TABLE `blog_moments` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `blog_id` int(11) NOT NULL,
  `text` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `entities` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `video` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `photos` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `mention_uids` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `share_type` tinyint(4) NOT NULL DEFAULT '0',
  `member_uids` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `has_geo` tinyint(1) NOT NULL DEFAULT '0',
  `lat` double NOT NULL DEFAULT '0',
  `long` double NOT NULL DEFAULT '0',
  `address` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `likes` int(11) NOT NULL DEFAULT '0',
  `commits` int(11) NOT NULL DEFAULT '0',
  `date` int(11) NOT NULL DEFAULT '0',
  `sort` int(11) NOT NULL DEFAULT '0',
  `topics` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '[]',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `blog_moments_uid_IDX` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- enterprise.blog_pts_updates definition
CREATE TABLE `blog_pts_updates` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `pts` int(11) NOT NULL DEFAULT '0',
  `pts_count` int(11) NOT NULL DEFAULT '0',
  `update_type` tinyint(4) NOT NULL DEFAULT '0',
  `update_data` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `date` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `blog_pts_updates_user_id_pts_IDX` (`user_id`,`pts`) USING BTREE,
  KEY `blog_pts_updates_user_id_IDX` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- enterprise.blog_rewards definition
CREATE TABLE `blog_rewards` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `target_uid` int(11) NOT NULL,
  `blog_id` int(11) NOT NULL,
  `amount` decimal(10,5) NOT NULL,
  `date` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `blog_rewards_uid_IDX` (`user_id`) USING BTREE,
  KEY `blog_rewards_target_IDX` (`target_uid`) USING BTREE,
  KEY `blog_rewards_blog_id_IDX` (`blog_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- enterprise.blog_topic_mappings definition
CREATE TABLE `blog_topic_mappings` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `topic_id` int(11) NOT NULL DEFAULT '0',
  `moment_id` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `topic_n_moment_on_mappings` (`topic_id`,`moment_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- enterprise.blog_topics definition
CREATE TABLE `blog_topics` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `ranking` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `blog_topics_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- enterprise.blog_user_privacies definition
CREATE TABLE `blog_user_privacies` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `key_type` tinyint(4) NOT NULL DEFAULT '0',
  `rules` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `blog_user_privacies_user_id` (`user_id`,`key_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- enterprise.blogs definition
CREATE TABLE `blogs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `read_update_id` int(11) NOT NULL DEFAULT '0',
  `read_blog_id` int(11) NOT NULL DEFAULT '0',
  `read_comment_id` int(11) NOT NULL DEFAULT '0',
  `read_max_id` int(11) NOT NULL DEFAULT '0',
  `moments` int(11) NOT NULL DEFAULT '0',
  `follows` int(11) NOT NULL DEFAULT '0',
  `fans` int(11) NOT NULL DEFAULT '0',
  `comments` int(11) NOT NULL DEFAULT '0',
  `likes` int(11) NOT NULL DEFAULT '0',
  `date` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `blog_users_user_id_IDX` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- enterprise.remittances definition
CREATE TABLE `remittances` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `chat_id` int(11) NOT NULL,
  `payer_uid` int(11) NOT NULL,
  `payee_uid` int(11) NOT NULL,
  `amount` decimal(10,5) NOT NULL,
  `status` tinyint(4) NOT NULL,
  `type` tinyint(4) NOT NULL,
  `description` varchar(64) NOT NULL,
  `create_date` int(11) DEFAULT NULL,
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `payer_on_remittances` (`payer_uid`),
  KEY `payee_on_remittances` (`payee_uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- enterprise.batch_send_messages definition
CREATE TABLE `batch_send_messages` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL,
  `message` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `to_users` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_batch_send_messages_on_uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- enterprise.blog_banned_users definition
CREATE TABLE `blog_banned_users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0',
  `ban_from` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `ban_to` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `users` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- enterprise.message_reactions definition
CREATE TABLE `message_reactions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` smallint(6) NOT NULL DEFAULT '0',
  `chat_id` bigint(20) NOT NULL DEFAULT '0',
  `message_id` int(11) NOT NULL DEFAULT '0',
  `user_id` int(11) NOT NULL DEFAULT '0',
  `reaction_id` smallint(6) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `mcu_on_reactions` (`message_id`,`chat_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
