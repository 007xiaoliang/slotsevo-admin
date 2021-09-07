/*
 Navicat Premium Data Transfer

 Source Server         : test
 Source Server Type    : MySQL
 Source Server Version : 80025
 Source Host           : www.xiaoliang007.asia:3306
 Source Schema         : slotsevo_admin

 Target Server Type    : MySQL
 Target Server Version : 80025
 File Encoding         : 65001

 Date: 07/09/2021 11:00:59
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (1, 'p', '1000', '/base/login', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (5, 'p', '1000', '/jwt/jsonInBlacklist', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (3, 'p', '1000', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (4, 'p', '1000', '/menu/getMenuList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (8, 'p', '1000', '/rpc/actInfo', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (7, 'p', '1000', '/rpc/themeInfo', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (6, 'p', '1000', '/rpc/userInfo', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (2, 'p', '1000', '/user/register', 'POST', '', '', '');

-- ----------------------------
-- Table structure for jwt_blacklists
-- ----------------------------
DROP TABLE IF EXISTS `jwt_blacklists`;
CREATE TABLE `jwt_blacklists`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `jwt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'jwt',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_jwt_blacklists_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of jwt_blacklists
-- ----------------------------
INSERT INTO `jwt_blacklists` VALUES (1, '2021-08-18 18:27:28', '2021-08-18 18:27:28', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMzI0ZjdmMDEtMWNmMC00YzQ4LWFlNzEtNjE1MTE0YzUwN2JkIiwiSUQiOjEsIlVzZXJuYW1lIjoieGlhb2xpYW5nIiwiTmlja05hbWUiOiJ4aWFvbGlhbmciLCJBdXRob3JpdHlJZCI6IjEwMDAiLCJCdWZmZXJUaW1lIjo4NjQwMCwiZXhwIjoxNjI5ODgxNjM2LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE2MjkyNzU4MzZ9.Cd5hzxOioqWf-_oh63oxhfC8Ea7J1yFT2psYpMYFXu4');
INSERT INTO `jwt_blacklists` VALUES (2, '2021-08-18 18:30:29', '2021-08-18 18:30:29', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMzI0ZjdmMDEtMWNmMC00YzQ4LWFlNzEtNjE1MTE0YzUwN2JkIiwiSUQiOjEsIlVzZXJuYW1lIjoieGlhb2xpYW5nIiwiTmlja05hbWUiOiJ4aWFvbGlhbmciLCJBdXRob3JpdHlJZCI6IjEwMDAiLCJCdWZmZXJUaW1lIjo4NjQwMCwiZXhwIjoxNjI5ODg3MzAwLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE2MjkyODE1MDB9.twuM6IgkIlvBXAMtWhmJ3F083hgAR21ZuOy-JqU21jw');
INSERT INTO `jwt_blacklists` VALUES (3, '2021-08-18 19:16:32', '2021-08-18 19:16:32', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMzI0ZjdmMDEtMWNmMC00YzQ4LWFlNzEtNjE1MTE0YzUwN2JkIiwiSUQiOjEsIlVzZXJuYW1lIjoieGlhb2xpYW5nIiwiTmlja05hbWUiOiJ4aWFvbGlhbmciLCJBdXRob3JpdHlJZCI6IjEwMDAiLCJCdWZmZXJUaW1lIjo4NjQwMCwiZXhwIjoxNjI5ODg3ODk5LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE2MjkyODIwOTl9.6K9OpwGwGijGgj5_ao4TtxN2UwbkdyUAdj3OOy_WMEE');
INSERT INTO `jwt_blacklists` VALUES (4, '2021-08-19 15:00:48', '2021-08-19 15:00:48', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMzI0ZjdmMDEtMWNmMC00YzQ4LWFlNzEtNjE1MTE0YzUwN2JkIiwiSUQiOjEsIlVzZXJuYW1lIjoieGlhb2xpYW5nIiwiTmlja05hbWUiOiJ4aWFvbGlhbmciLCJBdXRob3JpdHlJZCI6IjEwMDAiLCJCdWZmZXJUaW1lIjo4NjQwMCwiZXhwIjoxNjI5ODkwMjE2LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE2MjkyODQ0MTZ9.KBlYBOm1i79m0Ea17-pB9DnTAciXJ0PXg7fenVux1L8');
INSERT INTO `jwt_blacklists` VALUES (5, '2021-08-20 11:10:06', '2021-08-20 11:10:06', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMzI0ZjdmMDEtMWNmMC00YzQ4LWFlNzEtNjE1MTE0YzUwN2JkIiwiSUQiOjEsIlVzZXJuYW1lIjoieGlhb2xpYW5nIiwiTmlja05hbWUiOiJ4aWFvbGlhbmciLCJBdXRob3JpdHlJZCI6IjEwMDAiLCJCdWZmZXJUaW1lIjo4NjQwMCwiZXhwIjoxNjI5OTYxMjk2LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE2MjkzNTU0OTZ9.8IP_Fhp_IuNKOz6M05wWPbMzkBwYukW_3PrS0BsJRF8');
INSERT INTO `jwt_blacklists` VALUES (6, '2021-08-23 08:36:11', '2021-08-23 08:36:11', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMzI0ZjdmMDEtMWNmMC00YzQ4LWFlNzEtNjE1MTE0YzUwN2JkIiwiSUQiOjEsIlVzZXJuYW1lIjoieGlhb2xpYW5nIiwiTmlja05hbWUiOiJ4aWFvbGlhbmciLCJBdXRob3JpdHlJZCI6IjEwMDAiLCJCdWZmZXJUaW1lIjo4NjQwMCwiZXhwIjoxNjMwMzExMTM5LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE2Mjk3MDUzMzl9.brwXw73_p7lyt9tcem15baErYjvjTB0IpFm533AQye0');
INSERT INTO `jwt_blacklists` VALUES (7, '2021-08-24 06:08:05', '2021-08-24 06:08:05', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMzI0ZjdmMDEtMWNmMC00YzQ4LWFlNzEtNjE1MTE0YzUwN2JkIiwiSUQiOjEsIlVzZXJuYW1lIjoieGlhb2xpYW5nIiwiTmlja05hbWUiOiJ4aWFvbGlhbmciLCJBdXRob3JpdHlJZCI6IjEwMDAiLCJCdWZmZXJUaW1lIjo4NjQwMCwiZXhwIjoxNjMwMzc3OTcxLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE2Mjk3NzIxNzF9.jqVf3HuPy_xTzUcejKM5lbKlDouQmHVANL69zfVFXHo');
INSERT INTO `jwt_blacklists` VALUES (8, '2021-08-24 06:15:46', '2021-08-24 06:15:46', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMzI0ZjdmMDEtMWNmMC00YzQ4LWFlNzEtNjE1MTE0YzUwN2JkIiwiSUQiOjEsIlVzZXJuYW1lIjoieGlhb2xpYW5nIiwiTmlja05hbWUiOiJ4aWFvbGlhbmciLCJBdXRob3JpdHlJZCI6IjEwMDAiLCJCdWZmZXJUaW1lIjo4NjQwMCwiZXhwIjoxNjMwMzkwMDkyLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE2Mjk3ODQyOTJ9.xlsrU6xr42FpAQMbXa4v3DcrFsx_SLRDaQ8_Xp8O4n4');
INSERT INTO `jwt_blacklists` VALUES (9, '2021-08-24 07:06:46', '2021-08-24 07:06:46', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMzI0ZjdmMDEtMWNmMC00YzQ4LWFlNzEtNjE1MTE0YzUwN2JkIiwiSUQiOjEsIlVzZXJuYW1lIjoieGlhb2xpYW5nIiwiTmlja05hbWUiOiJ4aWFvbGlhbmciLCJBdXRob3JpdHlJZCI6IjEwMDAiLCJCdWZmZXJUaW1lIjo4NjQwMCwiZXhwIjoxNjMwMzkwNTU2LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE2Mjk3ODQ3NTZ9.Y8WPg9weUg_wFzE73KNWqB6sSl0URVlMPWR4cPqb4S4');
INSERT INTO `jwt_blacklists` VALUES (10, '2021-08-24 07:14:40', '2021-08-24 07:14:40', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMzI0ZjdmMDEtMWNmMC00YzQ4LWFlNzEtNjE1MTE0YzUwN2JkIiwiSUQiOjEsIlVzZXJuYW1lIjoieGlhb2xpYW5nIiwiTmlja05hbWUiOiJ4aWFvbGlhbmciLCJBdXRob3JpdHlJZCI6IjEwMDAiLCJCdWZmZXJUaW1lIjo4NjQwMCwiZXhwIjoxNjMwMzkzNjEzLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE2Mjk3ODc4MTN9.d8H5JO4Loudw7ScKOn6SPtCTsANGgunBojNT3gAHi0A');
INSERT INTO `jwt_blacklists` VALUES (11, '2021-08-24 07:23:30', '2021-08-24 07:23:30', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMzI0ZjdmMDEtMWNmMC00YzQ4LWFlNzEtNjE1MTE0YzUwN2JkIiwiSUQiOjEsIlVzZXJuYW1lIjoieGlhb2xpYW5nIiwiTmlja05hbWUiOiJ4aWFvbGlhbmciLCJBdXRob3JpdHlJZCI6IjEwMDAiLCJCdWZmZXJUaW1lIjo4NjQwMCwiZXhwIjoxNjMwMzk0MTY4LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE2Mjk3ODgzNjh9.fwJdz29McNpcH-Tw6puw3TCUexUk4XFsqkIexBkh1KU');
INSERT INTO `jwt_blacklists` VALUES (12, '2021-08-24 07:25:12', '2021-08-24 07:25:12', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMzI0ZjdmMDEtMWNmMC00YzQ4LWFlNzEtNjE1MTE0YzUwN2JkIiwiSUQiOjEsIlVzZXJuYW1lIjoieGlhb2xpYW5nIiwiTmlja05hbWUiOiJ4aWFvbGlhbmciLCJBdXRob3JpdHlJZCI6IjEwMDAiLCJCdWZmZXJUaW1lIjo4NjQwMCwiZXhwIjoxNjMwMzk0NjE0LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE2Mjk3ODg4MTR9.-hlrd6zIXuxb12pI47HilvZMUL6Koy_x-56FSW4MfM4');
INSERT INTO `jwt_blacklists` VALUES (13, '2021-08-24 09:22:13', '2021-08-24 09:22:13', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMzI0ZjdmMDEtMWNmMC00YzQ4LWFlNzEtNjE1MTE0YzUwN2JkIiwiSUQiOjEsIlVzZXJuYW1lIjoieGlhb2xpYW5nIiwiTmlja05hbWUiOiJ4aWFvbGlhbmciLCJBdXRob3JpdHlJZCI6IjEwMDAiLCJCdWZmZXJUaW1lIjo4NjQwMCwiZXhwIjoxNjMwMzk0NzE4LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE2Mjk3ODg5MTh9.UompQzpNzEZ7elbGobnn5-RUW5Bl5Kq0hfp9gNrrJxk');
INSERT INTO `jwt_blacklists` VALUES (14, '2021-08-24 10:51:22', '2021-08-24 10:51:22', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMzI0ZjdmMDEtMWNmMC00YzQ4LWFlNzEtNjE1MTE0YzUwN2JkIiwiSUQiOjEsIlVzZXJuYW1lIjoieGlhb2xpYW5nIiwiTmlja05hbWUiOiJ4aWFvbGlhbmciLCJBdXRob3JpdHlJZCI6IjEwMDAiLCJCdWZmZXJUaW1lIjo4NjQwMCwiZXhwIjoxNjMwNDA1NDAwLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE2Mjk3OTk2MDB9.XS2eWPnrKBnP73DldIRfLhOe7_Q4u5wLXqUV--hl-6o');
INSERT INTO `jwt_blacklists` VALUES (15, '2021-09-04 11:07:31', '2021-09-04 11:07:31', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMzI0ZjdmMDEtMWNmMC00YzQ4LWFlNzEtNjE1MTE0YzUwN2JkIiwiSUQiOjEsIlVzZXJuYW1lIjoieGlhb2xpYW5nIiwiTmlja05hbWUiOiJ4aWFvbGlhbmciLCJBdXRob3JpdHlJZCI6IjEwMDAiLCJCdWZmZXJUaW1lIjo4NjQwMCwiZXhwIjoxNjMxMzU3OTc0LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE2MzAyMTQyOTh9.oaZVBW6HGBUtL2xl9xwHBo0XR33BavFBlo2SHk-thpE');

-- ----------------------------
-- Table structure for sys_apis
-- ----------------------------
DROP TABLE IF EXISTS `sys_apis`;
CREATE TABLE `sys_apis`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'POST',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_apis_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 32 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_apis
-- ----------------------------
INSERT INTO `sys_apis` VALUES (1, '2021-08-18 16:53:39', '2021-08-18 16:53:39', NULL, '/base/login', '用户登录', 'base', 'POST');
INSERT INTO `sys_apis` VALUES (2, '2021-08-18 16:53:39', '2021-08-18 16:53:39', NULL, '/user/register', '用户注册', 'user', 'POST');
INSERT INTO `sys_apis` VALUES (3, '2021-08-18 16:53:39', '2021-08-18 16:53:39', NULL, '/menu/getMenu', '获取菜单树', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (4, '2021-08-18 16:53:39', '2021-08-18 16:53:39', NULL, '/menu/getMenuList', '分页获取基础menu列表', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (5, '2021-08-02 17:01:26', '2021-08-02 17:01:26', NULL, '/jwt/jsonInBlacklist', 'jwt加入黑名单(退出)', 'jwt', 'POST');
INSERT INTO `sys_apis` VALUES (31, '2021-08-25 16:12:12', '2021-08-25 16:12:14', NULL, '/rpc/userInfo', 'rpc获取用户信息', 'rpc', 'POST');

-- ----------------------------
-- Table structure for sys_authorities
-- ----------------------------
DROP TABLE IF EXISTS `sys_authorities`;
CREATE TABLE `sys_authorities`  (
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `authority_id` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色ID',
  `authority_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色名',
  `parent_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '父角色ID',
  `default_router` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'dashboard' COMMENT '默认菜单',
  PRIMARY KEY (`authority_id`) USING BTREE,
  UNIQUE INDEX `authority_id`(`authority_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_authorities
-- ----------------------------
INSERT INTO `sys_authorities` VALUES ('2021-08-18 16:53:39', '2021-08-18 16:53:39', NULL, '1000', '系统用户', '0', 'dashboard');

-- ----------------------------
-- Table structure for sys_authority_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_menus`;
CREATE TABLE `sys_authority_menus`  (
  `sys_authority_authority_id` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色ID',
  `sys_base_menu_id` bigint(0) UNSIGNED NOT NULL,
  PRIMARY KEY (`sys_authority_authority_id`, `sys_base_menu_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_authority_menus
-- ----------------------------
INSERT INTO `sys_authority_menus` VALUES ('1000', 1);
INSERT INTO `sys_authority_menus` VALUES ('1000', 2);
INSERT INTO `sys_authority_menus` VALUES ('1000', 3);
INSERT INTO `sys_authority_menus` VALUES ('1000', 4);
INSERT INTO `sys_authority_menus` VALUES ('1000', 5);
INSERT INTO `sys_authority_menus` VALUES ('1000', 6);
INSERT INTO `sys_authority_menus` VALUES ('1000', 7);
INSERT INTO `sys_authority_menus` VALUES ('1000', 8);
INSERT INTO `sys_authority_menus` VALUES ('1000', 9);

-- ----------------------------
-- Table structure for sys_base_menu_parameters
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menu_parameters`;
CREATE TABLE `sys_base_menu_parameters`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `sys_base_menu_id` bigint(0) UNSIGNED NULL DEFAULT NULL,
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址栏携带参数为params还是query',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址栏携带参数的key',
  `value` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址栏携带参数的值',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_base_menu_parameters_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_base_menu_parameters
-- ----------------------------

-- ----------------------------
-- Table structure for sys_base_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menus`;
CREATE TABLE `sys_base_menus`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `menu_level` bigint(0) UNSIGNED NULL DEFAULT NULL,
  `parent_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '父菜单ID',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '路由path',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '路由name',
  `hidden` tinyint(1) NULL DEFAULT NULL COMMENT '是否在列表隐藏',
  `component` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '对应前端文件路径',
  `sort` bigint(0) NULL DEFAULT NULL COMMENT '排序标记',
  `keep_alive` tinyint(1) NULL DEFAULT NULL COMMENT '附加属性',
  `default_menu` tinyint(1) NULL DEFAULT NULL COMMENT '附加属性',
  `title` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '附加属性',
  `icon` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '附加属性',
  `close_tab` tinyint(1) NULL DEFAULT NULL COMMENT '附加属性',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_base_menus_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_base_menus
-- ----------------------------
INSERT INTO `sys_base_menus` VALUES (1, '2021-08-24 17:20:19', '2021-08-24 17:20:17', NULL, 0, '0', 'dashboard', 'dashboard', 0, 'view/dashboard/index.vue', 1, 0, 0, '首页', 's-home', 0);
INSERT INTO `sys_base_menus` VALUES (2, '2021-08-18 16:53:39', '2021-08-18 16:53:39', NULL, 0, '0', 'backend', 'backend', 0, 'view/backend/index.vue', 2, 0, 0, '后端数据', 's-data', 0);
INSERT INTO `sys_base_menus` VALUES (3, '2021-08-18 16:53:39', '2021-08-18 16:53:39', NULL, 0, '2', 'user', 'user', 0, 'view/backend/user/user.vue', 3, 0, 0, '用户信息', 's-custom', 0);
INSERT INTO `sys_base_menus` VALUES (4, '2021-08-18 16:53:39', '2021-08-18 16:53:39', NULL, 0, '2', 'theme', 'theme', 0, 'view/backend/theme/theme.vue', 4, 0, 0, '主题配置', 's-grid', 0);
INSERT INTO `sys_base_menus` VALUES (5, '2021-08-24 10:32:30', '2021-08-24 10:32:34', NULL, 0, '2', 'activity', 'activity', 0, 'view/backend/activity/activity.vue', 5, 0, 0, '活动配置', 'menu', 0);
INSERT INTO `sys_base_menus` VALUES (6, '2021-08-24 10:32:30', '2021-08-24 10:32:34', NULL, 0, '0', 'control', 'control', 0, 'view/control/index.vue', 6, 0, 0, '产品数值', 's-marketing', 0);
INSERT INTO `sys_base_menus` VALUES (7, '2021-08-24 10:32:30', '2021-08-24 10:32:34', NULL, 0, '6', 'product', 'product', 0, 'view/control/product/product.vue', 7, 0, 0, '产品文案', 'folder', 0);
INSERT INTO `sys_base_menus` VALUES (8, '2021-08-24 10:32:30', '2021-08-24 10:32:34', NULL, 0, '6', 'data', 'data', 0, 'view/control/data/data.vue', 8, 0, 0, '数值控制', 'edit', 0);

-- ----------------------------
-- Table structure for sys_data_authority_id
-- ----------------------------
DROP TABLE IF EXISTS `sys_data_authority_id`;
CREATE TABLE `sys_data_authority_id`  (
  `sys_authority_authority_id` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色ID',
  `data_authority_id_authority_id` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_authority_authority_id`, `data_authority_id_authority_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_data_authority_id
-- ----------------------------
INSERT INTO `sys_data_authority_id` VALUES ('1000', '1000');

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `uuid` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户UUID',
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户登录名',
  `password` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户登录密码',
  `nick_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '系统用户' COMMENT '用户昵称',
  `header_img` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'http://qmplusimg.henrongyi.top/head.png' COMMENT '用户头像',
  `authority_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '888' COMMENT '用户角色ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_users_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
INSERT INTO `sys_users` VALUES (1, '2021-08-18 16:53:39', '2021-08-18 16:53:39', NULL, '324f7f01-1cf0-4c48-ae71-615114c507bd', 'xiaoliang', 'e10adc3949ba59abbe56e057f20f883e', 'xiaoliang', 'http://qmplusimg.henrongyi.top/Slots_header.jpg', '1000');

-- ----------------------------
-- View structure for authority_menu
-- ----------------------------
DROP VIEW IF EXISTS `authority_menu`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `authority_menu` AS select `sys_base_menus`.`id` AS `id`,`sys_base_menus`.`created_at` AS `created_at`,`sys_base_menus`.`updated_at` AS `updated_at`,`sys_base_menus`.`deleted_at` AS `deleted_at`,`sys_base_menus`.`menu_level` AS `menu_level`,`sys_base_menus`.`parent_id` AS `parent_id`,`sys_base_menus`.`path` AS `path`,`sys_base_menus`.`name` AS `name`,`sys_base_menus`.`hidden` AS `hidden`,`sys_base_menus`.`component` AS `component`,`sys_base_menus`.`title` AS `title`,`sys_base_menus`.`icon` AS `icon`,`sys_base_menus`.`sort` AS `sort`,`sys_authority_menus`.`sys_authority_authority_id` AS `authority_id`,`sys_authority_menus`.`sys_base_menu_id` AS `menu_id`,`sys_base_menus`.`keep_alive` AS `keep_alive`,`sys_base_menus`.`default_menu` AS `default_menu` from (`sys_authority_menus` join `sys_base_menus` on((`sys_authority_menus`.`sys_base_menu_id` = `sys_base_menus`.`id`)));

SET FOREIGN_KEY_CHECKS = 1;
