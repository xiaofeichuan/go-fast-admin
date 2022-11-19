/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : localhost:3306
 Source Schema         : gin_fast_admin

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 19/11/2022 11:56:56
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config`  (
  `id` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_config
-- ----------------------------

-- ----------------------------
-- Table structure for sys_dict
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict`;
CREATE TABLE `sys_dict`  (
  `id` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dict
-- ----------------------------

-- ----------------------------
-- Table structure for sys_dict_detail
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_detail`;
CREATE TABLE `sys_dict_detail`  (
  `id` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dict_detail
-- ----------------------------

-- ----------------------------
-- Table structure for sys_gen_table
-- ----------------------------
DROP TABLE IF EXISTS `sys_gen_table`;
CREATE TABLE `sys_gen_table`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `table_name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '表名称',
  `table_comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '表注释',
  `model_name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '实体名称',
  `business_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '业务名称',
  `function_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '功能名称',
  `param_name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '参数名称',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '代码生成器-业务表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_gen_table
-- ----------------------------
INSERT INTO `sys_gen_table` VALUES (1, 'sys_user', '用户信息表', 'SysUser', 'user', '用户信息', 'sysUser', '2022-10-29 14:30:21', '2022-10-29 14:30:21', NULL);

-- ----------------------------
-- Table structure for sys_gen_table_column
-- ----------------------------
DROP TABLE IF EXISTS `sys_gen_table_column`;
CREATE TABLE `sys_gen_table_column`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `table_id` bigint NULL DEFAULT NULL COMMENT '表编号',
  `column_name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字段名',
  `column_comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字段注释',
  `column_type` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字段类型',
  `param_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '参数名称',
  `go_field` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'go字段名',
  `go_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'go类型',
  `is_pk` tinyint(1) NULL DEFAULT NULL COMMENT '是否主键（0否 1是）',
  `is_edit` tinyint NULL DEFAULT NULL COMMENT '是否编辑字段（0否 1是）',
  `is_list` tinyint NULL DEFAULT NULL COMMENT '是否列表字段（0否 1是）',
  `is_query` tinyint NULL DEFAULT NULL COMMENT '是否查询字段（0否 1是）',
  `query_method` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '查询方式',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '代码生成器-业务表字段' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_gen_table_column
-- ----------------------------
INSERT INTO `sys_gen_table_column` VALUES (1, 1, 'id', '用户ID', 'bigint', 'id', 'Id', 'uint64', 1, 0, 1, 0, '', '2022-10-29 14:30:21', '2022-10-29 14:30:21', NULL);
INSERT INTO `sys_gen_table_column` VALUES (2, 1, 'account', '用户账号', 'varchar(30)', 'account', 'Account', 'string', 0, 0, 1, 0, '', '2022-10-29 14:30:21', '2022-10-29 14:30:21', NULL);
INSERT INTO `sys_gen_table_column` VALUES (3, 1, 'nick_name', '用户昵称', 'varchar(30)', 'nickName', 'NickName', 'string', 0, 0, 1, 0, '', '2022-10-29 14:30:21', '2022-10-29 14:30:21', NULL);
INSERT INTO `sys_gen_table_column` VALUES (4, 1, 'user_type', '用户类型（0普通账号，1超级管理员）', 'tinyint', 'userType', 'UserType', 'int', 0, 0, 1, 0, '', '2022-10-29 14:30:21', '2022-10-29 14:30:21', NULL);
INSERT INTO `sys_gen_table_column` VALUES (5, 1, 'email', '用户邮箱', 'varchar(50)', 'email', 'Email', 'string', 0, 0, 1, 0, '', '2022-10-29 14:30:21', '2022-10-29 14:30:21', NULL);
INSERT INTO `sys_gen_table_column` VALUES (6, 1, 'phone', '手机号码', 'varchar(20)', 'phone', 'Phone', 'string', 0, 0, 1, 0, '', '2022-10-29 14:30:21', '2022-10-29 14:30:21', NULL);
INSERT INTO `sys_gen_table_column` VALUES (7, 1, 'password', '密码', 'varchar(100)', 'password', 'Password', 'string', 0, 0, 1, 0, '', '2022-10-29 14:30:21', '2022-10-29 14:30:21', NULL);
INSERT INTO `sys_gen_table_column` VALUES (8, 1, 'sex', '用户性别（0未知，1男，2女）', 'tinyint', 'sex', 'Sex', 'bool', 0, 0, 1, 0, '', '2022-10-29 14:30:21', '2022-10-29 14:30:21', NULL);
INSERT INTO `sys_gen_table_column` VALUES (9, 1, 'avatar', '头像地址', 'varchar(255)', 'avatar', 'Avatar', 'string', 0, 0, 1, 0, '', '2022-10-29 14:30:21', '2022-10-29 14:30:21', NULL);
INSERT INTO `sys_gen_table_column` VALUES (10, 1, 'status', '帐号状态（0正常 1停用）', 'tinyint', 'status', 'Status', 'int', 0, 0, 1, 0, '', '2022-10-29 14:30:21', '2022-10-29 14:30:21', NULL);
INSERT INTO `sys_gen_table_column` VALUES (11, 1, 'created_at', '创建时间', 'datetime', 'createdAt', 'CreatedAt', 'time.Time', 0, 0, 1, 0, '', '2022-10-29 14:30:21', '2022-10-29 14:30:21', NULL);
INSERT INTO `sys_gen_table_column` VALUES (12, 1, 'updated_at', '更新时间', 'datetime', 'updatedAt', 'UpdatedAt', 'time.Time', 0, 0, 1, 0, '', '2022-10-29 14:30:21', '2022-10-29 14:30:21', NULL);
INSERT INTO `sys_gen_table_column` VALUES (13, 1, 'deleted_at', '删除时间', 'datetime', 'deletedAt', 'DeletedAt', 'time.Time', 0, 0, 1, 0, '', '2022-10-29 14:30:21', '2022-10-29 14:30:21', NULL);
INSERT INTO `sys_gen_table_column` VALUES (14, 1, 'remark', '备注', 'varchar(500)', 'remark', 'Remark', 'string', 0, 0, 1, 0, '', '2022-10-29 14:30:21', '2022-10-29 14:30:21', NULL);

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `parent_id` bigint NULL DEFAULT 0 COMMENT '父菜单ID',
  `menu_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单名称',
  `path` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '路由地址',
  `is_iframe` tinyint NULL DEFAULT NULL COMMENT '是否外链',
  `is_hide` tinyint NULL DEFAULT NULL COMMENT '是否隐藏',
  `is_cache` tinyint NULL DEFAULT NULL COMMENT '是否缓存',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '组件路径',
  `icon` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '#' COMMENT '菜单图标',
  `status` tinyint NOT NULL COMMENT '菜单状态（0正常 1停用）',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色名称',
  `role_code` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色编码',
  `status` tinyint NOT NULL COMMENT '角色状态（0正常 1停用）',
  `is_deleted` tinyint NOT NULL COMMENT '删除标记（0未删除，1已删除）',
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建者',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新者',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role
-- ----------------------------

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu`  (
  `id` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `account` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户账号',
  `user_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名称',
  `user_type` tinyint NOT NULL COMMENT '用户类型（0普通账号，1超级管理员）',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户邮箱',
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机号码',
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `salt` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '密码盐',
  `sex` tinyint NULL DEFAULT NULL COMMENT '用户性别（0未知，1男，2女）',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '头像地址',
  `status` tinyint NOT NULL COMMENT '帐号状态（0正常 1停用）',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 'superAdmin', '超级管理员', 1, '514634736@qq.com', '19999999999', 'ec71d37f4340c223896afd45aaf3cf06', '41b278387b85404', 0, NULL, 0, '2022-11-08 14:27:47', '2022-11-08 14:27:51', NULL, NULL);

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role`  (
  `user_id` bigint NOT NULL COMMENT '用户ID',
  `role_id` bigint NOT NULL COMMENT '角色ID'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户和角色关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
INSERT INTO `sys_user_role` VALUES (1, 1);
INSERT INTO `sys_user_role` VALUES (1, 1);
INSERT INTO `sys_user_role` VALUES (1, 1);

SET FOREIGN_KEY_CHECKS = 1;
