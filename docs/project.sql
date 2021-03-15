/*
 Navicat Premium Data Transfer

 Source Server         : 本地MySQL
 Source Server Type    : MySQL
 Source Server Version : 80016
 Source Host           : localhost:3306
 Source Schema         : project

 Target Server Type    : MySQL
 Target Server Version : 80016
 File Encoding         : 65001

 Date: 15/03/2021 23:20:49
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `username` varchar(64) NOT NULL COMMENT '用户状态',
  `password` varchar(200) NOT NULL COMMENT '用户密码',
  `email` varchar(255) NOT NULL COMMENT '用户邮箱',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1 启用 2 停用',
  `last_login_ip` varchar(20) NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `last_login_time` datetime NOT NULL COMMENT '最后登录时间',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `unq_username` (`username`) COMMENT '用户名称唯一'
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='菜单信息';

-- ----------------------------
-- Records of admin
-- ----------------------------
BEGIN;
INSERT INTO `admin` VALUES (1, 'admin', '$2a$10$8nFmszU3KJ4Eb4oBy2Krse5XfNriwmHPTtVjnponCkcZfbXQ3yvgu', 'admin@gmail.com', 1, '127.0.0.1:52873', '2021-03-15 22:51:19', '2021-03-10 17:38:48', '2021-03-15 22:51:19');
COMMIT;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` int(11) NOT NULL COMMENT '创建者',
  `title` varchar(64) NOT NULL COMMENT '标题',
  `content` varchar(64) NOT NULL COMMENT '内容',
  `img` varchar(255) NOT NULL COMMENT '图片',
  `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '类型',
  `see_num` tinyint(1) NOT NULL DEFAULT '1' COMMENT '类型',
  `comment_num` tinyint(1) NOT NULL DEFAULT '1' COMMENT '类型',
  `recommend` tinyint(1) NOT NULL DEFAULT '1' COMMENT '类型',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1 启用 2 停用',
  `sort` mediumint(4) NOT NULL DEFAULT '100' COMMENT '排序',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='菜单信息';

-- ----------------------------
-- Records of article
-- ----------------------------
BEGIN;
INSERT INTO `article` VALUES (3, 1, '添加统一商户审计', '1222222222222', '/static/uploads/202103/5577006791947779410.png', 0, 11, 0, 0, 1, 100, '2021-03-10 09:00:51', '2021-03-15 22:52:52');
INSERT INTO `article` VALUES (4, 1, '测试文章', '我的测试文档', '/static/uploads/202103/5577006791947779410.jpg', 1, 3, 1, 1, 1, 100, '2021-03-15 22:51:50', '2021-03-15 22:52:42');
INSERT INTO `article` VALUES (5, 1, '轮播图一', '测试图片上传', '/static/uploads/202103/8674665223082153551.jpg', 1, 1, 1, 1, 1, 100, '2021-03-15 22:54:25', '2021-03-15 22:54:25');
COMMIT;

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `pid` int(11) NOT NULL COMMENT '父ID',
  `path` varchar(200) NOT NULL COMMENT '内容',
  `cate_name` varchar(255) NOT NULL COMMENT '图片',
  `recommend` tinyint(1) NOT NULL DEFAULT '1' COMMENT '类型',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1 启用 2 停用',
  `sort` mediumint(4) NOT NULL COMMENT '排序',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='菜单信息';

-- ----------------------------
-- Records of category
-- ----------------------------
BEGIN;
INSERT INTO `category` VALUES (1, 0, '', 'PHP', 1, 1, 100, '2021-03-11 21:42:15', '2021-03-15 21:51:23');
INSERT INTO `category` VALUES (2, 0, '', 'Golang', 1, 1, 200, '2021-03-11 13:46:09', '2021-03-15 22:12:24');
INSERT INTO `category` VALUES (3, 0, '', 'Java', 1, 1, 300, '2021-03-11 15:34:22', '2021-03-15 21:54:12');
INSERT INTO `category` VALUES (4, 0, '', 'Javascript', 1, 1, 400, '2021-03-15 21:53:16', '2021-03-15 21:54:22');
INSERT INTO `category` VALUES (5, 0, '', 'Vue', 1, 1, 500, '2021-03-15 21:53:36', '2021-03-15 21:54:29');
INSERT INTO `category` VALUES (6, 0, '', 'React', 1, 1, 600, '2021-03-15 21:53:49', '2021-03-15 21:54:37');
INSERT INTO `category` VALUES (7, 1, '', 'Laravel', 1, 1, 100, '2021-03-15 22:11:28', '2021-03-15 22:11:28');
INSERT INTO `category` VALUES (8, 1, '', 'Yii2', 1, 1, 100, '2021-03-15 22:11:52', '2021-03-15 22:11:52');
INSERT INTO `category` VALUES (9, 1, '', 'ThinkPHP', 1, 1, 100, '2021-03-15 22:12:07', '2021-03-15 22:12:07');
INSERT INTO `category` VALUES (10, 2, '', 'Beego', 1, 1, 100, '2021-03-15 22:12:58', '2021-03-15 22:12:58');
INSERT INTO `category` VALUES (11, 2, '', 'Gin', 1, 1, 100, '2021-03-15 22:13:11', '2021-03-15 22:13:11');
COMMIT;

-- ----------------------------
-- Table structure for image
-- ----------------------------
DROP TABLE IF EXISTS `image`;
CREATE TABLE `image` (
  `image_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` int(11) NOT NULL COMMENT '所属用户',
  `title` varchar(64) NOT NULL COMMENT '用户状态',
  `description` varchar(200) NOT NULL COMMENT '用户密码',
  `url` varchar(255) NOT NULL COMMENT '用户邮箱',
  `type` tinyint(1) NOT NULL COMMENT '图片类型 1 轮播图 2  普通图片',
  `sort` mediumint(4) NOT NULL COMMENT '排序',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1 启用 2 停用',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`image_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='菜单信息';

-- ----------------------------
-- Records of image
-- ----------------------------
BEGIN;
INSERT INTO `image` VALUES (2, 1, '添加统一商户审计', '121212', '/static/uploads/202103/5577006791947779410.png', 2, 0, 1, '2021-03-10 08:03:37', '2021-03-10 08:03:37');
INSERT INTO `image` VALUES (3, 1, '添加统一商户审计', '121212', '/static/uploads/202103/5577006791947779410.jpg', 2, 0, 1, '2021-03-10 08:07:20', '2021-03-10 08:07:20');
INSERT INTO `image` VALUES (4, 1, '1212', 'gjgkvm', '/static/uploads/202103/5577006791947779410.jpg', 2, 0, 1, '2021-03-15 22:32:47', '2021-03-15 22:32:47');
INSERT INTO `image` VALUES (5, 1, '1212', '12121212', '/static/uploads/202103/8674665223082153551.jpg', 2, 0, 1, '2021-03-15 22:33:05', '2021-03-15 22:33:05');
INSERT INTO `image` VALUES (6, 1, '轮播图', '轮播图一', '/static/uploads/202103/6129484611666145821.jpg', 1, 0, 1, '2021-03-15 22:55:03', '2021-03-15 22:55:03');
COMMIT;

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `pid` int(11) NOT NULL COMMENT '父类ID',
  `menu_name` varchar(64) NOT NULL COMMENT '菜单名称',
  `icons` varchar(32) NOT NULL COMMENT '图标',
  `url` varchar(255) NOT NULL COMMENT '菜单地址',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1 启用 2 停用',
  `sort` mediumint(4) NOT NULL COMMENT '排序',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='菜单信息';

-- ----------------------------
-- Records of menu
-- ----------------------------
BEGIN;
INSERT INTO `menu` VALUES (1, 0, '菜单首页', 'fa fa-home', '/admin/menu/index', 1, 100, '2021-03-10 16:29:12', '2021-03-15 21:47:43');
INSERT INTO `menu` VALUES (2, 0, '分类管理', 'fa fa-list', '/admin/category/index', 1, 100, '2021-03-10 16:29:14', '2021-03-15 21:20:08');
INSERT INTO `menu` VALUES (3, 0, '用户列表', 'fa fa-user', '/admin/admin/index', 1, 100, '2021-03-11 13:44:20', '2021-03-11 13:44:20');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
