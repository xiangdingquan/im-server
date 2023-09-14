SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `imweb`
--

-- --------------------------------------------------------

--
-- 表的结构 `admin_logs`
--

CREATE TABLE `admin_logs` (
  `id` int(11) NOT NULL,
  `admin_id` int(11) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `route` varchar(255) DEFAULT NULL,
  `type` varchar(32) DEFAULT NULL,
  `request` text,
  `response` text,
  `request_ip` varchar(32) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- --------------------------------------------------------

--
-- 表的结构 `agents`
--

CREATE TABLE `agents` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `username` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `api_token` varchar(255) CHARACTER SET utf8mb4 DEFAULT NULL,
  `channel_id` int(10) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

-- --------------------------------------------------------

--
-- 表的结构 `async_routers`
--

CREATE TABLE `async_routers` (
  `id` int(11) NOT NULL,
  `name` varchar(32) NOT NULL,
  `parentId` int(11) DEFAULT NULL,
  `redirect` varchar(255) DEFAULT NULL,
  `component` varchar(255) DEFAULT NULL,
  `meta` text,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `path` varchar(255) DEFAULT NULL,
  `permission` text
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

--
-- 转存表中的数据 `async_routers`
--

INSERT INTO `async_routers` (`id`, `name`, `parentId`, `redirect`, `component`, `meta`, `created_at`, `updated_at`, `path`, `permission`) VALUES
(1, 'Member', 0, '/members/member-list', 'RouteView', '{\"title\":\"\\u7528\\u6237\\u7ba1\\u7406\",\"icon\":\"table\",\"permission\":[\"table\"]}', '2022-08-30 04:01:55', '2022-08-30 14:20:23', '/members', NULL),
(2, 'MemberList', 1, NULL, 'MemberList', '{\"title\":\"\\u7528\\u6237\\u5217\\u8868\"}', '2022-08-30 04:01:55', '2022-08-30 06:13:52', '/members/member-list', '[{\"action\":\"test\",\"describe\":\"\\u6d4b\\u8bd5\",\"defaultCheck\":false},{\"action\":\"export\",\"describe\":\"\\u5bfc\\u51fa\",\"defaultCheck\":false},{\"action\":\"import\",\"describe\":\"\\u5bfc\\u5165\",\"defaultCheck\":false}]'),
(3, 'GroupList', 1, NULL, 'GroupList', '{\"title\":\"\\u7fa4\\u804a\\u7ba1\\u7406\",\"keepAlive\":\"true\",\"permission\":[\"table\"]}', '2022-08-30 04:01:55', '2022-08-30 04:01:55', '/members/member-group', NULL),
(4, 'finance_center', 0, '/members/member-list', 'RouteView', '{\"title\":\"\\u8d22\\u52a1\\u4e2d\\u5fc3\",\"icon\":\"money-collect\",\"permission\":[\"table\"]}', '2022-08-30 04:01:55', '2022-08-30 14:20:26', '/finance_center', NULL),
(5, 'RedPacketList', 4, NULL, 'RedPacketList', '{\"title\":\"\\u7ea2\\u5305\\u8bb0\\u5f55\",\"keepAlive\":\"true\",\"permission\":[\"table\"]}', '2022-08-30 04:01:55', '2022-08-30 04:01:55', '/finance_center/redpacket_list', NULL),
(6, 'RechargeList', 4, NULL, 'RechargeList', '{\"title\":\"\\u5145\\u503c\\u8bb0\\u5f55\",\"keepAlive\":\"true\",\"permission\":[\"table\"]}', '2022-08-30 04:01:55', '2022-08-30 04:01:55', '/finance_center/recharge_list', NULL),
(7, 'WithdrawalList', 4, NULL, 'WithdrawalList', '{\"title\":\"\\u63d0\\u73b0\\u8bb0\\u5f55\",\"keepAlive\":\"true\",\"permission\":[\"table\"]}', '2022-08-30 04:01:55', '2022-08-30 04:01:55', '/finance_center/withdrawal_list', NULL),
(8, 'Verify', 0, NULL, 'RouteView', '{\"title\":\"\\u5ba1\\u6838\\u7ba1\\u7406\",\"icon\":\"check-square\"}', '2022-08-30 04:01:55', '2022-08-30 00:34:24', '/verify', 'null'),
(9, 'VerifyWithdrawalList', 8, NULL, 'VerifyWithdrawalList', '{\"title\":\"\\u63d0\\u73b0\\u5ba1\\u6838\",\"keepAlive\":\"true\",\"permission\":[\"table\"]}', '2022-08-30 04:01:56', '2022-08-30 13:39:33', '/verify/withdrawal_list', NULL),
(10, 'ProxyManage', 0, NULL, 'RouteView', '{\"title\":\"\\u6e20\\u9053\\u7ba1\\u7406\",\"icon\":\"profile\"}', '2022-08-30 04:01:56', '2022-08-30 00:34:32', '/proxy_manage', 'null'),
(11, 'SystemManager', 10, NULL, 'SystemManager', '{\"title\":\"\\u7ba1\\u7406\\u5458\\u5217\\u8868\",\"keepAlive\":\"true\",\"permission\":[\"table\"]}', '2022-08-30 04:01:56', '2022-08-30 04:01:56', '/channel_manage/manager', NULL),
(12, 'discoverGroups', 10, NULL, 'discoverGroups', '{\"title\":\"\\u53d1\\u73b0\\u9875\\u9762\",\"keepAlive\":\"true\",\"permission\":[\"table\"]}', '2022-08-30 04:01:56', '2022-08-30 04:01:56', '/channel_manage/discoverMenus', NULL),
(13, 'Message', 0, NULL, 'RouteView', '{\"title\":\"\\u6d88\\u606f\\u7ba1\\u7406\",\"icon\":\"message\"}', '2022-08-30 04:01:56', '2022-08-30 00:00:46', '/message', 'null'),
(14, 'GroupUsers', 13, NULL, 'GroupUsers', '{\"title\":\"\\u7528\\u6237\\u5206\\u7ec4\",\"keepAlive\":\"true\",\"permission\":[\"table\"]}', '2022-08-30 04:01:56', '2022-08-30 04:01:56', '/message/group_users', NULL),
(15, 'MessageGroup', 13, NULL, 'MessageGroup', '{\"title\":\"\\u7fa4\\u53d1\\u6d88\\u606f\",\"keepAlive\":\"true\",\"permission\":[\"table\"]}', '2022-08-30 04:01:56', '2022-08-30 04:01:56', '/message/group', NULL),
(16, 'MessageGroupRecord', 13, NULL, 'MessageGroupRecord', '{\"title\":\"\\u5386\\u53f2\\u7fa4\\u53d1\",\"keepAlive\":\"true\",\"permission\":[\"table\"]}', '2022-08-30 04:01:56', '2022-08-30 04:01:56', '/message/record', NULL),
(17, 'System', 0, NULL, 'RouteView', '{\"title\":\"\\u7cfb\\u7edf\\u8bbe\\u7f6e\",\"permission\":\"table\",\"icon\":\"cluster\"}', '2022-08-30 04:01:56', '2022-08-30 06:14:38', '/system', NULL),
(18, 'Personal', 17, NULL, 'Personal', '{\"title\":\"\\u4e2a\\u4eba\\u8d44\\u6599\",\"permission\":\"table\"}', '2022-08-30 04:01:56', '2022-08-30 22:00:41', '/system/personal', NULL),
(19, 'SysConfig', 17, NULL, 'SysConfig', '{\"title\":\"\\u5ba2\\u670d\\u7aef\\u914d\\u7f6e\",\"permission\":[\"table\"]}', '2022-08-30 04:01:56', '2022-08-30 21:41:20', '/system/sysconfig', NULL),
(20, 'ServerSysConfig', 17, NULL, 'ServerSysConfig', '{\"title\":\"\\u670d\\u52a1\\u7aef\\u914d\\u7f6e\",\"keepAlive\":\"true\",\"permission\":[\"table\"]}', '2022-08-30 04:01:56', '2022-08-30 04:01:56', '/system/serversysconfig', NULL),
(21, 'RechargeScene', 17, NULL, 'RechargeScene', '{\"title\":\"\\u652f\\u4ed8\\u901a\\u9053\",\"keepAlive\":\"true\",\"permission\":[\"table\"]}', '2022-08-30 04:01:56', '2022-08-30 04:01:56', '/system/recharge_scene', NULL),
(22, 'WithdrawSence', 17, NULL, 'WithdrawSence', '{\"title\":\"\\u63d0\\u73b0\\u901a\\u9053\",\"permission\":\"table\"}', '2022-08-30 04:01:56', '2022-08-30 22:00:58', '/system/withdraw_sence', NULL),
(23, 'Support', 0, NULL, 'RouteView', '{\"title\":\"\\u8d85\\u7ea7\\u7ba1\\u7406\",\"permission\":\"table1\",\"icon\":\"global\"}', '2022-08-30 04:01:56', '2022-08-30 22:22:15', '/support', NULL),
(24, 'MenuNav', 23, NULL, 'MenuNav', '{\"title\":\"\\u83dc\\u5355\\u914d\\u7f6e\",\"permission\":\"table1\",\"icon\":null}', '2022-08-30 04:01:56', '2022-08-30 22:22:09', '/support/menu_nav', NULL),
(25, 'RoleList', 23, NULL, 'support/RoleList', '{\"title\":\"\\u89d2\\u8272\\u7ba1\\u7406\"}', '2022-08-30 00:42:01', '2022-08-30 14:40:51', '/support/rolelist', 'null'),
(26, 'AdminList', 23, NULL, 'support/AdminList', '{\"title\":\"\\u540e\\u53f0\\u7528\\u6237\"}', '2022-08-30 03:23:31', '2022-08-30 03:45:23', '/system/admin_list', NULL),
(27, 'DataCenter', 1, NULL, 'members/DataCenter', '{\"title\":\"\\u6570\\u636e\\u7edf\\u8ba1\"}', '2022-08-30 20:46:42', '2022-08-30 20:50:30', '/members/data_center', NULL),
(28, 'AdminLog', 23, NULL, 'support/AdminLog', '{\"title\":\"\\u64cd\\u4f5c\\u65e5\\u5fd7\"}', '2022-08-30 18:46:39', '2022-08-30 00:51:24', '/support/adminlog', NULL),
(29, 'BannedIps', 1, NULL, 'members/BannedIps', '{\"title\":\"IP\\u9ed1\\u540d\\u5355\"}', '2022-08-30 03:05:59', '2022-08-30 03:18:45', '/members/banned_ips', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `failed_jobs`
--

CREATE TABLE `failed_jobs` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `connection` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `queue` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `payload` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `exception` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `failed_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

-- --------------------------------------------------------

--
-- 表的结构 `menu_actions`
--

CREATE TABLE `menu_actions` (
  `id` int(11) NOT NULL,
  `actions` varchar(32) DEFAULT NULL,
  `describe` varchar(32) DEFAULT NULL,
  `defaultCheck` int(1) NOT NULL DEFAULT '0',
  `menu_id` int(10) NOT NULL,
  `controller` varchar(32) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

--
-- 转存表中的数据 `menu_actions`
--

INSERT INTO `menu_actions` (`id`, `actions`, `describe`, `defaultCheck`, `menu_id`, `controller`, `created_at`, `updated_at`) VALUES
(1, 'add', '添加', 0, 1, NULL, '2022-08-30 18:16:30', '2022-08-30 18:16:30'),
(2, 'edit', '修改资料', 0, 2, NULL, '2022-08-30 18:40:18', '2022-08-30 23:14:14'),
(3, 'add_group', '添加分组', 0, 12, NULL, '2022-08-30 20:05:27', '2022-08-30 20:05:27'),
(4, 'add_menu', '添加菜单', 0, 12, NULL, '2022-08-30 20:05:58', '2022-08-30 20:05:58'),
(5, 'del', '删除', 0, 12, NULL, '2022-08-30 20:07:05', '2022-08-30 05:00:14'),
(6, 'add', '添加权限', 0, 24, NULL, '2022-08-30 05:03:36', '2022-08-30 05:03:36'),
(7, 'edit', '编辑', 0, 24, NULL, '2022-08-30 05:03:47', '2022-08-30 05:03:47'),
(8, 'edit', '编辑', 0, 14, NULL, '2022-08-30 19:06:04', '2022-08-30 19:06:04');

-- --------------------------------------------------------

--
-- 表的结构 `messages`
--

CREATE TABLE `messages` (
  `id` int(11) NOT NULL,
  `title` varchar(30) NOT NULL,
  `type` int(2) DEFAULT '1' COMMENT '1消息，2通知',
  `remarks` varchar(255) DEFAULT NULL,
  `messages` varchar(255) NOT NULL,
  `users_ids` text NOT NULL,
  `status` int(4) DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- 表的结构 `message_groups`
--

CREATE TABLE `message_groups` (
  `id` int(11) NOT NULL,
  `name` varchar(32) DEFAULT NULL,
  `remarks` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- --------------------------------------------------------

--
-- 表的结构 `message_group_details`
--

CREATE TABLE `message_group_details` (
  `id` int(11) NOT NULL,
  `group_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- --------------------------------------------------------

--
-- 表的结构 `migrations`
--

CREATE TABLE `migrations` (
  `id` int(10) UNSIGNED NOT NULL,
  `migration` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `batch` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

--
-- 转存表中的数据 `migrations`
--

INSERT INTO `migrations` (`id`, `migration`, `batch`) VALUES
(1, '2014_10_12_000000_create_users_table', 1),
(2, '2014_10_12_100000_create_password_resets_table', 1),
(3, '2019_08_19_000000_create_failed_jobs_table', 1),
(4, '2021_04_12_040605_create_red_packets_table', 1);

-- --------------------------------------------------------

--
-- 表的结构 `password_resets`
--

CREATE TABLE `password_resets` (
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `token` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

-- --------------------------------------------------------

--
-- 表的结构 `recharge_scenes`
--

CREATE TABLE `recharge_scenes` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `title` varchar(255) NOT NULL,
  `desc` varchar(255) NOT NULL,
  `scene` varchar(255) NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `img` varchar(255) DEFAULT NULL,
  `sort` int(11) DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- --------------------------------------------------------

--
-- 表的结构 `roles`
--

CREATE TABLE `roles` (
  `id` int(255) NOT NULL,
  `menu_ids` text,
  `action_ids` text,
  `name` varchar(255) DEFAULT NULL,
  `describe` varchar(255) DEFAULT NULL,
  `key` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

--
-- 转存表中的数据 `roles`
--

INSERT INTO `roles` (`id`, `menu_ids`, `action_ids`, `name`, `describe`, `key`, `created_at`, `updated_at`) VALUES
(1, '[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,22,21,23,24,25]', '[2,1,3,4,5,7,6,8]', '超级管理员', NULL, 'admin', '2022-08-30 06:43:30', '2022-08-30 19:06:14');

-- --------------------------------------------------------

--
-- 表的结构 `sys_configs`
--

CREATE TABLE `sys_configs` (
  `id` int(11) NOT NULL,
  `key` varchar(60) NOT NULL,
  `value` text,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- 表的结构 `users`
--

CREATE TABLE `users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `remember_token` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `api_token` varchar(255) CHARACTER SET utf8mb4 DEFAULT NULL,
  `googlekey` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL,
  `googlekey_imgurl` varchar(255) CHARACTER SET utf8mb4 DEFAULT NULL,
  `roleId` int(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

--
-- 转存表中的数据 `users`
--

INSERT INTO `users` (`id`, `name`, `password`, `remember_token`, `created_at`, `updated_at`, `api_token`, `googlekey`, `googlekey_imgurl`, `roleId`) VALUES
(1, 'admin', '$2y$10$fRgCZRXBpOr4Dki2/zxaeO51GanZz1aYLm.mkaHsBiINhNTYOVEUq', NULL, NULL, '2022-08-30 16:34:38', 'ZItRD598JseqwY99UtcwTSdizLXvQuRixDotRVpj4EorKitN3pRTuKbbdL3Z', 'ZWE3TJT72UMJTLAR', 'https://api.qrserver.com/v1/create-qr-code/?size=200x200&data=otpauth%3A%2F%2Ftotp%2Fadmin%40jk.gochat8.com%3Fsecret%3DZWE3TJT72UMJTLAR&ecc=M', 1);

-- --------------------------------------------------------

--
-- 表的结构 `withdraws`
--

CREATE TABLE `withdraws` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `order_token` varchar(255) DEFAULT NULL,
  `order_id` varchar(80) NOT NULL,
  `user_id` int(11) NOT NULL,
  `money` int(11) NOT NULL DEFAULT '0',
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `reply` varchar(255) DEFAULT NULL,
  `server_money` float NOT NULL DEFAULT '0',
  `alipay_account` varchar(50) DEFAULT NULL,
  `bank_account` varchar(50) DEFAULT NULL,
  `bank_branch` varchar(100) DEFAULT NULL,
  `bank_name` varchar(50) DEFAULT NULL,
  `alipay_name` varchar(50) DEFAULT NULL,
  `remark` varchar(200) DEFAULT NULL,
  `scene` varchar(50) DEFAULT 'self',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `real_name` varchar(50) DEFAULT NULL,
  `pay_type` varchar(50) NOT NULL DEFAULT 'Alipay'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- 表的结构 `withdraw_scenes`
--

CREATE TABLE `withdraw_scenes` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `title` varchar(255) NOT NULL,
  `desc` varchar(255) NOT NULL,
  `scene` varchar(255) NOT NULL,
  `pay_type` varchar(20) DEFAULT 'Alipay',
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

--
-- 转储表的索引
--

--
-- 表的索引 `admin_logs`
--
ALTER TABLE `admin_logs`
  ADD PRIMARY KEY (`id`) USING BTREE;

--
-- 表的索引 `agents`
--
ALTER TABLE `agents`
  ADD PRIMARY KEY (`id`) USING BTREE;

--
-- 表的索引 `async_routers`
--
ALTER TABLE `async_routers`
  ADD PRIMARY KEY (`id`) USING BTREE;

--
-- 表的索引 `failed_jobs`
--
ALTER TABLE `failed_jobs`
  ADD PRIMARY KEY (`id`) USING BTREE;

--
-- 表的索引 `menu_actions`
--
ALTER TABLE `menu_actions`
  ADD PRIMARY KEY (`id`) USING BTREE;

--
-- 表的索引 `messages`
--
ALTER TABLE `messages`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `message_groups`
--
ALTER TABLE `message_groups`
  ADD PRIMARY KEY (`id`) USING BTREE;

--
-- 表的索引 `message_group_details`
--
ALTER TABLE `message_group_details`
  ADD PRIMARY KEY (`id`) USING BTREE;

--
-- 表的索引 `migrations`
--
ALTER TABLE `migrations`
  ADD PRIMARY KEY (`id`) USING BTREE;

--
-- 表的索引 `password_resets`
--
ALTER TABLE `password_resets`
  ADD KEY `password_resets_email_index` (`email`) USING BTREE;

--
-- 表的索引 `recharge_scenes`
--
ALTER TABLE `recharge_scenes`
  ADD PRIMARY KEY (`id`) USING BTREE;

--
-- 表的索引 `roles`
--
ALTER TABLE `roles`
  ADD PRIMARY KEY (`id`) USING BTREE;

--
-- 表的索引 `sys_configs`
--
ALTER TABLE `sys_configs`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`) USING BTREE;

--
-- 表的索引 `withdraws`
--
ALTER TABLE `withdraws`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `withdraw_scenes`
--
ALTER TABLE `withdraw_scenes`
  ADD PRIMARY KEY (`id`) USING BTREE;

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `admin_logs`
--
ALTER TABLE `admin_logs`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `agents`
--
ALTER TABLE `agents`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `async_routers`
--
ALTER TABLE `async_routers`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=30;

--
-- 使用表AUTO_INCREMENT `failed_jobs`
--
ALTER TABLE `failed_jobs`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `menu_actions`
--
ALTER TABLE `menu_actions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- 使用表AUTO_INCREMENT `messages`
--
ALTER TABLE `messages`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `message_groups`
--
ALTER TABLE `message_groups`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `message_group_details`
--
ALTER TABLE `message_group_details`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `migrations`
--
ALTER TABLE `migrations`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- 使用表AUTO_INCREMENT `recharge_scenes`
--
ALTER TABLE `recharge_scenes`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `roles`
--
ALTER TABLE `roles`
  MODIFY `id` int(255) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- 使用表AUTO_INCREMENT `sys_configs`
--
ALTER TABLE `sys_configs`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- 使用表AUTO_INCREMENT `withdraws`
--
ALTER TABLE `withdraws`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `withdraw_scenes`
--
ALTER TABLE `withdraw_scenes`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
