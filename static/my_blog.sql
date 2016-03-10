/*
Navicat MySQL Data Transfer

Source Server         : 我的数据库
Source Server Version : 50520
Source Host           : localhost:3306
Source Database       : my_blog

Target Server Type    : MYSQL
Target Server Version : 50520
File Encoding         : 65001

Date: 2016-03-10 18:57:47
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for my_blog_admin
-- ----------------------------
DROP TABLE IF EXISTS `my_blog_admin`;
CREATE TABLE `my_blog_admin` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(20) NOT NULL COMMENT '账号',
  `password` char(40) NOT NULL COMMENT '密码(使用sha1()加密)',
  `email` varchar(50) DEFAULT NULL COMMENT '管理员邮箱',
  `power` varchar(255) DEFAULT NULL COMMENT '权限',
  `auto_key` char(20) DEFAULT NULL COMMENT '自动登录的KEY',
  `access_token` char(40) DEFAULT NULL COMMENT '自动登录TOKEN',
  `status` tinyint(2) NOT NULL COMMENT '管理员状态',
  `create_time` int(11) DEFAULT NULL COMMENT '注册时间',
  `last_time` int(11) DEFAULT NULL COMMENT '最后登录的时间',
  `last_ip` varchar(20) DEFAULT NULL COMMENT '最后登录IP',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='后台管理员信息表';

-- ----------------------------
-- Records of my_blog_admin
-- ----------------------------
INSERT INTO `my_blog_admin` VALUES ('1', 'gongyan', 'eda3d8cb5282a4522ad1f1209891ba8a9b321d6f', '610455122@qq.com', '', '', '', '1', '1457604078', '1457606724', '[::1]:50921');
INSERT INTO `my_blog_admin` VALUES ('2', 'liujinxing', 'a75d5dd01d9a3ff45da4a304fdbbaf80596a0bc3', '821901008@qq.com', '', '', '', '1', '1457606311', '1457606311', '[::1]:50480');

-- ----------------------------
-- Table structure for my_blog_category
-- ----------------------------
DROP TABLE IF EXISTS `my_blog_category`;
CREATE TABLE `my_blog_category` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '分类ID',
  `pid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '父类ID(0最高级)',
  `path` varchar(255) NOT NULL DEFAULT '0' COMMENT '分类路径',
  `cate_name` varchar(255) NOT NULL COMMENT '分类名称',
  `sort` int(11) NOT NULL DEFAULT '1' COMMENT '分类排序',
  `recommend` tinyint(2) NOT NULL DEFAULT '0' COMMENT '是否为推荐',
  `status` tinyint(2) NOT NULL DEFAULT '1' COMMENT '分类状态(1启用2 停用)',
  `create_time` int(11) NOT NULL COMMENT '创建时间',
  `create_id` int(11) NOT NULL COMMENT '创建用户',
  `update_time` int(11) NOT NULL COMMENT '修改时间',
  `update_id` int(11) NOT NULL COMMENT '修改者',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8 COMMENT='文章分类信息表';

-- ----------------------------
-- Records of my_blog_category
-- ----------------------------
INSERT INTO `my_blog_category` VALUES ('1', '0', '0', '新闻公告', '1', '0', '1', '1453270463', '1', '1453270490', '1');
INSERT INTO `my_blog_category` VALUES ('2', '0', '0', '游戏资料', '2', '0', '1', '1450921203', '0', '0', '0');
INSERT INTO `my_blog_category` VALUES ('3', '0', '0', '游戏攻略', '3', '0', '1', '1450921233', '0', '0', '0');
INSERT INTO `my_blog_category` VALUES ('4', '1', '0', '公告', '1', '1', '1', '1450921565', '0', '0', '0');
INSERT INTO `my_blog_category` VALUES ('5', '1', '0', '新闻', '2', '1', '1', '1450921585', '0', '0', '0');
INSERT INTO `my_blog_category` VALUES ('6', '1', '0', '活动', '3', '1', '1', '1450921606', '0', '0', '0');
INSERT INTO `my_blog_category` VALUES ('7', '2', '0', '新手指南', '1', '1', '1', '1450921661', '0', '0', '0');
INSERT INTO `my_blog_category` VALUES ('8', '2', '0', '系统介绍', '2', '1', '1', '1450921705', '0', '0', '0');
INSERT INTO `my_blog_category` VALUES ('9', '2', '0', '高手进阶', '3', '1', '1', '1450921730', '0', '0', '0');
INSERT INTO `my_blog_category` VALUES ('10', '2', '0', '特色玩法', '4', '1', '1', '1450921749', '0', '0', '0');
INSERT INTO `my_blog_category` VALUES ('11', '3', '0', '攻略指南', '1', '1', '0', '1450921885', '0', '1453195691', '1');
INSERT INTO `my_blog_category` VALUES ('12', '0', '0', '我的测试数据', '1', '1', '1', '1453195665', '1', '1453195682', '1');
INSERT INTO `my_blog_category` VALUES ('13', '0', '0', '测试数据345', '1', '0', '1', '1456736196', '1', '1456737077', '1');

-- ----------------------------
-- Table structure for my_blog_menu
-- ----------------------------
DROP TABLE IF EXISTS `my_blog_menu`;
CREATE TABLE `my_blog_menu` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `menu_name` varchar(20) DEFAULT '' COMMENT '导航名称',
  `icons` varchar(20) DEFAULT NULL COMMENT '使用的图标',
  `controller_name` varchar(20) DEFAULT NULL COMMENT '访问的控制器',
  `action_name` varchar(20) DEFAULT 'index' COMMENT '访问的控制器方法名',
  `parent_id` int(11) DEFAULT '0' COMMENT '父级导航ID',
  `status` tinyint(1) DEFAULT '1' COMMENT '导航栏的状态（0 停用 1 启用）',
  `sort` int(6) DEFAULT '100' COMMENT '导航的排序',
  `create_time` int(11) DEFAULT NULL COMMENT '创建时间',
  `create_id` int(11) DEFAULT NULL COMMENT '创建管理员',
  `update_time` int(11) DEFAULT NULL,
  `update_id` int(11) DEFAULT NULL COMMENT '修改管理员',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8 COMMENT='后台导航栏信息表';

-- ----------------------------
-- Records of my_blog_menu
-- ----------------------------
INSERT INTO `my_blog_menu` VALUES ('1', '管理员信息', 'glyphicon-user', 'admin', 'index', '0', '1', '1', '1453199489', '1', '1453270710', '1');
INSERT INTO `my_blog_menu` VALUES ('2', '文章分类信息', 'glyphicon-list', 'category', 'index', '0', '1', '2', '1453200049', '1', '1453200049', '1');
INSERT INTO `my_blog_menu` VALUES ('3', '导航栏信息', 'glyphicon-th', 'menu', 'index', '0', '1', '3', '1453200107', '1', '1453200107', '1');
INSERT INTO `my_blog_menu` VALUES ('4', 'UI界面&amp;元素', 'fa-desktop', '', '', '0', '1', '4', '1453200234', '1', '1453200234', '1');
INSERT INTO `my_blog_menu` VALUES ('5', '布局', 'fa-desktop', '', '', '4', '1', '5', '1453200309', '1', '1453200309', '1');
INSERT INTO `my_blog_menu` VALUES ('6', '头部导航', 'fa-desktop', 'other', 'top', '5', '1', '6', '1453200429', '1', '1453276171', '1');
INSERT INTO `my_blog_menu` VALUES ('7', '其他页面', 'fa-file-o', '', '', '0', '1', '8', '1453201293', '1', '1453201293', '1');
INSERT INTO `my_blog_menu` VALUES ('8', 'Error 404', 'fa-file-o', 'other', 'error404', '7', '1', '1', '1453201342', '1', '1453201342', '1');
INSERT INTO `my_blog_menu` VALUES ('9', 'Error 500', 'fa-file-o', 'other', 'error500', '7', '1', '2', '1453201375', '1', '1453201375', '1');
INSERT INTO `my_blog_menu` VALUES ('10', '空白页', 'fa-file-o', 'other', 'blankpage', '7', '1', '3', '1453201421', '1', '1453201421', '1');

-- ----------------------------
-- Table structure for my_blog_menus
-- ----------------------------
DROP TABLE IF EXISTS `my_blog_menus`;
CREATE TABLE `my_blog_menus` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '栏目ID',
  `pid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '父类ID(只支持两级分类)',
  `menu_name` varchar(50) NOT NULL COMMENT '栏目名称',
  `icons` varchar(50) NOT NULL DEFAULT 'icon-desktop' COMMENT '使用的icons',
  `url` varchar(50) DEFAULT NULL COMMENT '访问地址',
  `status` tinyint(2) NOT NULL DEFAULT '1' COMMENT '状态（1启用 0 停用）',
  `sort` int(4) NOT NULL DEFAULT '100' COMMENT '排序字段',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '修改时间',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建用户',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改用户',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='使用SimpliQ的样式的导航栏样式';

-- ----------------------------
-- Records of my_blog_menus
-- ----------------------------
INSERT INTO `my_blog_menus` VALUES ('1', '0', '后台管理', 'icon-folder-close-alt', '', '1', '1', '0', '1457590913', '0', '1');
INSERT INTO `my_blog_menus` VALUES ('2', '1', '管理员管理', 'icon-user', '/admin/index', '1', '1', '0', '1457585793', '0', '1');
INSERT INTO `my_blog_menus` VALUES ('3', '1', '后台栏目管理', 'icon-reorder', '/menus/index', '1', '2', '0', '1457585799', '0', '1');
INSERT INTO `my_blog_menus` VALUES ('4', '1', 'Icons预览', 'icon-star', '/other/index', '1', '3', '0', '1457585807', '0', '1');
