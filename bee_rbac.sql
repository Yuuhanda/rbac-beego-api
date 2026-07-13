-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jul 13, 2026 at 07:55 AM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `bee_rbac`
--

-- --------------------------------------------------------

--
-- Table structure for table `api_route`
--

CREATE TABLE `api_route` (
  `id` int(11) NOT NULL,
  `path` varchar(255) NOT NULL,
  `method` varchar(10) NOT NULL,
  `controller` varchar(100) NOT NULL,
  `action` varchar(100) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `api_route`
--

INSERT INTO `api_route` (`id`, `path`, `method`, `controller`, `action`, `description`, `created_at`, `updated_at`) VALUES
(1, '/auth/login', 'POST', 'AuthController', 'Login', 'API endpoint for POST /auth/login', '2024-12-30 08:41:49', '2026-07-13 12:54:12'),
(2, '/api/routes/scan', 'POST', 'APIRouteController', 'ScanRoutes', 'API endpoint for POST /api/routes/scan', '2024-12-30 08:41:49', '2026-07-13 12:54:12'),
(3, '/api/routes/list', 'GET', 'APIRouteController', 'ListRoutes', 'API endpoint for GET /api/routes/list', '2024-12-30 08:41:49', '2026-07-13 12:54:12'),
(4, '/api/roles', 'POST', 'AuthRolesController', 'Create', 'API endpoint for POST /api/roles', '2024-12-30 08:41:49', '2026-07-13 12:54:12'),
(5, '/api/roles/:id', 'GET', 'AuthRolesController', 'Get', 'API endpoint for GET /api/roles/:id', '2024-12-30 08:41:49', '2024-12-30 08:41:49'),
(6, '/api/roles', 'GET', 'AuthRolesController', 'List', 'API endpoint for GET /api/roles', '2024-12-30 08:41:49', '2026-07-13 12:54:12'),
(7, '/api/roles/:id', 'PUT', 'AuthRolesController', 'Update', 'API endpoint for PUT /api/roles/:id', '2024-12-30 08:41:49', '2024-12-30 08:41:49'),
(8, '/api/roles/:id', 'DELETE', 'AuthRolesController', 'Delete', 'API endpoint for DELETE /api/roles/:id', '2024-12-30 08:41:49', '2024-12-30 08:41:49'),
(9, '/api/user-roles', 'POST', 'AuthRolesUserController', 'Create', 'API endpoint for POST /api/user-roles', '2024-12-30 08:41:49', '2026-07-13 12:54:12'),
(10, '/api/user-roles/user/:userId', 'GET', 'AuthRolesUserController', 'GetUserRoles', 'API endpoint for GET /api/user-roles/user/:userId', '2024-12-30 08:41:49', '2026-07-13 12:54:12'),
(11, '/api/user-roles/role/:roleId', 'GET', 'AuthRolesUserController', 'GetRoleUsers', 'API endpoint for GET /api/user-roles/role/:roleId', '2024-12-30 08:41:49', '2026-07-13 12:54:12'),
(12, '/api/user-roles/:userId/:roleId', 'DELETE', 'AuthRolesUserController', 'Delete', 'API endpoint for DELETE /api/user-roles/:userId/:roleId', '2024-12-30 08:41:49', '2026-07-13 12:54:12'),
(13, '/api/auth-items', 'POST', 'AuthItemController', 'Create', 'API endpoint for POST /api/auth-items', '2024-12-30 08:41:49', '2026-07-13 12:54:12'),
(14, '/api/auth-items/:id', 'GET', 'AuthItemController', 'Get', 'API endpoint for GET /api/auth-items/:id', '2024-12-30 08:41:49', '2026-07-13 12:54:12'),
(15, '/api/auth-items', 'GET', 'AuthItemController', 'List', 'API endpoint for GET /api/auth-items', '2024-12-30 08:41:49', '2026-07-13 12:54:12'),
(16, '/api/auth-items/:id', 'PUT', 'AuthItemController', 'Update', 'API endpoint for PUT /api/auth-items/:id', '2024-12-30 08:41:49', '2026-07-13 12:54:12'),
(17, '/api/auth-items/:id', 'DELETE', 'AuthItemController', 'Delete', 'API endpoint for DELETE /api/auth-items/:id', '2024-12-30 08:41:49', '2026-07-13 12:54:12'),
(19, '/user', 'POST', 'UserController', 'CreateUser', 'API endpoint for POST /user', '2024-12-30 08:41:49', '2026-07-13 12:54:12'),
(20, '/auth/logout', 'POST', 'AuthController', 'Logout', 'API endpoint for POST /auth/logout', '2024-12-30 08:41:49', '2026-07-13 12:54:12'),
(21, '/user/:id', 'GET', 'UserController', 'GetUser', 'API endpoint for GET /user/:id', '2024-12-30 08:41:49', '2026-07-13 12:54:12'),
(22, '/user-update/:id', 'PUT', 'UserController', 'UpdateUser', 'API endpoint for PUT /user-update/:id', '2024-12-30 08:41:49', '2026-07-13 12:54:12'),
(23, '/user-delete/:id', 'DELETE', 'UserController', 'DeleteUser', 'API endpoint for DELETE /user-delete/:id', '2024-12-30 08:41:49', '2026-07-13 12:54:12'),
(24, '/users', 'GET', 'UserController', 'ListUsers', 'API endpoint for GET /users', '2024-12-30 08:41:49', '2026-07-13 12:54:12'),
(27, '/user/:id/visits', 'GET', 'UserVisitLogController', 'GetUserVisits', 'API endpoint for GET /user/:id/visits', '2026-07-11 04:40:10', '2026-07-13 12:54:12');

-- --------------------------------------------------------

--
-- Table structure for table `auth_item`
--

CREATE TABLE `auth_item` (
  `id` int(11) NOT NULL,
  `role` varchar(32) NOT NULL,
  `path` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `auth_item`
--

INSERT INTO `auth_item` (`id`, `role`, `path`) VALUES
(1, 'superadmin', '/api/auth-items'),
(2, 'superadmin', '/api/auth-items/:id'),
(4, 'superadmin', '/api/roles'),
(5, 'superadmin', '/api/roles/:id'),
(6, 'superadmin', '/api/routes/list'),
(7, 'superadmin', '/api/routes/scan'),
(8, 'superadmin', '/api/user-roles'),
(9, 'superadmin', '/api/user-roles/:userId/:roleId'),
(10, 'superadmin', '/api/user-roles/role/:roleId'),
(11, 'superadmin', '/api/user-roles/user/:userId'),
(12, 'superadmin', '/auth/login'),
(13, 'superadmin', '/auth/logout'),
(14, 'superadmin', '/user'),
(15, 'superadmin', '/user-delete/:id'),
(16, 'superadmin', '/user-update/:id'),
(17, 'superadmin', '/user/:id'),
(19, 'superadmin', '/users'),
(20, 'admin', '/users'),
(21, 'admin', '/api/roles'),
(22, 'admin', '/api/routes/scan'),
(23, 'admin', '/api/routes/list');

-- --------------------------------------------------------

--
-- Table structure for table `auth_roles`
--

CREATE TABLE `auth_roles` (
  `code` varchar(32) NOT NULL,
  `name` varchar(100) NOT NULL,
  `description` varchar(255) NOT NULL,
  `created_at` datetime(2) NOT NULL,
  `updated_at` datetime(2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `auth_roles`
--

INSERT INTO `auth_roles` (`code`, `name`, `description`, `created_at`, `updated_at`) VALUES
('admin', 'Administrator Updated', 'Updated admin role', '0000-00-00 00:00:00.00', '2026-07-10 06:50:45.70'),
('superadmin', 'Super Admin', 'Super Admin', '2024-12-20 02:39:58.22', '2024-12-20 02:39:58.22'),
('test', 'test Admin', 'test Admin', '2024-12-20 04:04:59.15', '2024-12-20 04:04:59.15'),
('test1', 'test Admin', 'test Admin', '2024-12-20 07:00:22.13', '2024-12-20 07:00:22.13'),
('tester-more', 'tester more', 'tester role more', '2026-07-10 06:58:10.03', '2026-07-10 06:58:10.03');

-- --------------------------------------------------------

--
-- Table structure for table `auth_roles_user`
--

CREATE TABLE `auth_roles_user` (
  `user_id` int(11) NOT NULL,
  `roles_code` varchar(32) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `id` int(11) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password_hash` varchar(255) NOT NULL,
  `status` int(11) NOT NULL DEFAULT 1,
  `superadmin` smallint(6) DEFAULT 0,
  `created_at` datetime(2) NOT NULL,
  `updated_at` datetime(2) NOT NULL,
  `registration_ip` varchar(15) DEFAULT NULL,
  `email` varchar(128) DEFAULT NULL,
  `auth_key` varchar(255) DEFAULT NULL,
  `bind_to_ip` varchar(255) DEFAULT NULL,
  `email_confirmed` int(11) NOT NULL,
  `confirmation_token` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`id`, `username`, `password_hash`, `status`, `superadmin`, `created_at`, `updated_at`, `registration_ip`, `email`, `auth_key`, `bind_to_ip`, `email_confirmed`, `confirmation_token`) VALUES
(1, 'superadmin', '$2y$13$bp2w2.mTeJ/ORRVlEjA.jOHw0o49vwAJ.A15RTPjnSyk05M.20ZyS', 1, 1, '0000-00-00 00:00:00.00', '2026-07-03 07:50:48.39', NULL, 'super@mail.com', '20260703145048_N31APklfHqP/FEKO9/uQHzIDZ7L9ps3m8d8RmOdd1o8=', '', 1, ''),
(2, 'bogosbinted', '$2a$10$z357DoGIsf9taHMhlC.5FeVZInJTu9kcXrq/9X/ZU4Owayv2Swt4W', 1, 1, '2024-12-14 06:51:48.00', '2026-07-10 06:18:01.92', '::1', 'bogos@mail.com', '20260710131801_X6WxxydilZCsy8QAFv/jEIeIZXG1+39cGU5k/wBErTY=', '', 0, ''),
(13, 'Sisyphuss', '$2a$10$eUFZI1QdBB4sccI.5b/phuEz8ZqqVvZZUTcKiV6t5xiS9G36IkiGO', 1, 0, '2024-12-23 03:58:44.11', '2024-12-23 03:58:44.11', '::1', 'sisyphuss@mail.com', '', '', 0, ''),
(14, 'johndoe_updated', '$2a$10$1fNI6PI0WRVkPBpzQG9Jkufg3NHciPjA6EUtQL6bIdT5rEnSic.ZO', 1, 0, '2026-07-09 23:30:22.00', '2026-07-10 06:31:34.73', '::1', 'john@example.com', '', '', 0, '');

-- --------------------------------------------------------

--
-- Table structure for table `user_visit_log`
--

CREATE TABLE `user_visit_log` (
  `id` int(11) NOT NULL,
  `token` varchar(255) NOT NULL,
  `ip` varchar(15) NOT NULL,
  `language` char(2) NOT NULL,
  `user_agent` varchar(255) NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `visit_time` int(11) NOT NULL,
  `browser` varchar(30) DEFAULT NULL,
  `os` varchar(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `user_visit_log`
--

INSERT INTO `user_visit_log` (`id`, `token`, `ip`, `language`, `user_agent`, `user_id`, `visit_time`, `browser`, `os`) VALUES
(1, '241220093715', '::1', '', 'PostmanRuntime/7.43.0', 1, 1734662235, 'Unknown', 'Unknown'),
(2, '241220094651', '::1', '', 'PostmanRuntime/7.43.0', 1, 1734662811, 'Unknown', 'Unknown'),
(3, '241220100006', '::1', '', 'PostmanRuntime/7.43.0', 2, 1734663606, 'Unknown', 'Unknown'),
(4, '241220100052', '::1', '', 'PostmanRuntime/7.43.0', 1, 1734663652, 'Unknown', 'Unknown'),
(5, '241220100358', '::1', '', 'PostmanRuntime/7.43.0', 2, 1734663838, 'Unknown', 'Unknown'),
(6, '241220101007', '::1', '', 'PostmanRuntime/7.43.0', 2, 1734664207, 'Unknown', 'Unknown'),
(7, '241220101019', '::1', '', 'PostmanRuntime/7.43.0', 2, 1734664219, 'Unknown', 'Unknown'),
(8, '241220101034', '::1', '', 'PostmanRuntime/7.43.0', 2, 1734664234, 'Unknown', 'Unknown'),
(9, '241220101352', '::1', '', 'PostmanRuntime/7.43.0', 2, 1734664432, 'Unknown', 'Unknown'),
(10, '241220101359', '::1', '', 'PostmanRuntime/7.43.0', 2, 1734664439, 'Unknown', 'Unknown'),
(11, '241220101811', '::1', '', 'PostmanRuntime/7.43.0', 2, 1734664691, 'Unknown', 'Unknown'),
(12, '241220102303', '::1', '', 'PostmanRuntime/7.43.0', 2, 1734664983, 'Unknown', 'Unknown'),
(13, '241220102530', '::1', '', 'PostmanRuntime/7.43.0', 2, 1734665130, 'Unknown', 'Unknown'),
(14, '241220102718', '::1', '', 'PostmanRuntime/7.43.0', 2, 1734665238, 'Unknown', 'Unknown'),
(15, '241220103359', '::1', '', 'PostmanRuntime/7.43.0', 2, 1734665639, 'Unknown', 'Unknown'),
(16, '241220131508', '::1', '', 'PostmanRuntime/7.43.0', 1, 1734675308, 'Unknown', 'Unknown'),
(17, '241220135952', '::1', '', 'PostmanRuntime/7.43.0', 1, 1734677992, 'Unknown', 'Unknown'),
(18, '241220140344', '::1', '', 'PostmanRuntime/7.43.0', 1, 1734678224, 'Unknown', 'Unknown'),
(19, '241220142927', '::1', '', 'PostmanRuntime/7.43.0', 1, 1734679767, 'Unknown', 'Unknown'),
(20, '241220143039', '::1', '', 'PostmanRuntime/7.43.0', 1, 1734679839, 'Unknown', 'Unknown'),
(21, '241220143805', '::1', 'en', 'PostmanRuntime/7.43.0', 1, 1734680285, 'Unknown', 'Unknown'),
(22, '241220143829', '::1', 'en', 'PostmanRuntime/7.43.0', 1, 1734680309, 'Unknown', 'Unknown'),
(23, '241223094126', '::1', 'en', 'PostmanRuntime/7.43.0', 1, 1734921686, 'Unknown', 'Unknown'),
(24, '241223100322', '::1', 'en', 'PostmanRuntime/7.43.0', 1, 1734923002, 'Unknown', 'Unknown'),
(25, '241223100838', '::1', 'en', 'PostmanRuntime/7.43.0', 1, 1734923318, 'Unknown', 'Unknown'),
(26, '241223104856', '::1', 'en', 'PostmanRuntime/7.43.0', 1, 1734925736, 'Unknown', 'Unknown'),
(27, '241223105008', '::1', 'en', 'PostmanRuntime/7.43.0', 1, 1734925808, 'Unknown', 'Unknown'),
(28, '241223105751', '::1', 'en', 'PostmanRuntime/7.43.0', 1, 1734926271, 'Unknown', 'Unknown'),
(29, '241223105809', '::1', 'en', 'PostmanRuntime/7.43.0', 1, 1734926289, 'Unknown', 'Unknown'),
(30, '241223133052', '::1', 'en', 'PostmanRuntime/7.43.0', 1, 1734935452, 'Unknown', 'Unknown'),
(31, '241230083252', '::1', 'en', 'PostmanRuntime/7.43.0', 1, 1735522372, 'Unknown', 'Unknown'),
(32, '260703143843', '::1', 'en', 'PostmanRuntime/1.9.5', 1, 1783064323, 'Unknown', 'Unknown'),
(33, '260703143938', '::1', 'en', 'PostmanRuntime/1.9.5', 1, 1783064378, 'Unknown', 'Unknown'),
(34, '260703144410', '::1', 'en', 'PostmanRuntime/1.9.5', 1, 1783064650, 'Unknown', 'Unknown'),
(35, '260703144415', '::1', 'en', 'PostmanRuntime/1.9.5', 1, 1783064655, 'Unknown', 'Unknown'),
(36, '260703145028', '::1', 'en', 'PostmanRuntime/1.9.5', 1, 1783065028, 'Unknown', 'Unknown'),
(37, '260703145048', '::1', 'en', 'PostmanRuntime/1.9.5', 1, 1783065048, 'Unknown', 'Unknown'),
(38, '260710131644', '::1', 'en', 'PostmanRuntime/1.9.5', 2, 1783664204, 'Unknown', 'Unknown'),
(39, '260710131739', '::1', 'en', 'PostmanRuntime/1.9.5', 2, 1783664259, 'Unknown', 'Unknown'),
(40, '260710131801', '::1', 'en', 'PostmanRuntime/1.9.5', 2, 1783664281, 'Unknown', 'Unknown');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `api_route`
--
ALTER TABLE `api_route`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `path_method` (`path`,`method`);

--
-- Indexes for table `auth_item`
--
ALTER TABLE `auth_item`
  ADD PRIMARY KEY (`id`),
  ADD KEY `route` (`path`),
  ADD KEY `role` (`role`);

--
-- Indexes for table `auth_roles`
--
ALTER TABLE `auth_roles`
  ADD PRIMARY KEY (`code`);

--
-- Indexes for table `auth_roles_user`
--
ALTER TABLE `auth_roles_user`
  ADD PRIMARY KEY (`user_id`),
  ADD KEY `roles_code` (`roles_code`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user_visit_log`
--
ALTER TABLE `user_visit_log`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `api_route`
--
ALTER TABLE `api_route`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=28;

--
-- AUTO_INCREMENT for table `auth_item`
--
ALTER TABLE `auth_item`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=24;

--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- AUTO_INCREMENT for table `user_visit_log`
--
ALTER TABLE `user_visit_log`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=41;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `auth_item`
--
ALTER TABLE `auth_item`
  ADD CONSTRAINT `role` FOREIGN KEY (`role`) REFERENCES `auth_roles` (`code`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `route` FOREIGN KEY (`path`) REFERENCES `api_route` (`path`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `auth_roles_user`
--
ALTER TABLE `auth_roles_user`
  ADD CONSTRAINT `user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

--
-- Constraints for table `user_visit_log`
--
ALTER TABLE `user_visit_log`
  ADD CONSTRAINT `user_visit_log_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE SET NULL ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
