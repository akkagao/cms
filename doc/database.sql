SET NAMES UTF8;
  create database if not exists cmsadmin default charset utf8 collate utf8_general_ci;
  grant select,update,delete,insert,alter,create,drop on cmsadmin.* to "cmsadmin"@"%" identified by "cmsadmin";
  grant select,update,delete,insert,alter,create,drop on cmsadmin.* to "cmsadmin"@"localhost" identified by "cmsadmin";
USE cmsadmin;


/*
Navicat MySQL Data Transfer

Source Server         : root
Source Server Version : 50629
Source Host           : localhost:3306
Source Database       : cmsadmin

Target Server Type    : MYSQL
Target Server Version : 50629
File Encoding         : 65001

Date: 2016-09-04 17:49:48
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for t_admuser
-- ----------------------------
DROP TABLE IF EXISTS `t_admuser`;
CREATE TABLE `t_admuser` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `accout` varchar(255) NOT NULL DEFAULT '',
  `mail` varchar(255) NOT NULL DEFAULT '',
  `name` varchar(255) NOT NULL DEFAULT '',
  `phone` varchar(255) NOT NULL DEFAULT '',
  `department` varchar(255) NOT NULL DEFAULT '',
  `password` varchar(255) NOT NULL DEFAULT '',
  `createtime` datetime NOT NULL,
  `updatetime` datetime NOT NULL,
  `isdel` tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `admUserAcout` (`accout`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of t_admuser
-- ----------------------------
INSERT INTO `t_admuser` VALUES ('1', 'admin', 'gaojianwen@hotmail.com', 'CrazyWolf', '1391387065', '研发', 'SddggEssSsmtsJgFUdFgInnFGsEfgBs2Sd==', '2016-08-27 15:05:11', '2016-09-03 10:56:45', '1');
INSERT INTO `t_admuser` VALUES ('2', 'test', 'test@test.com', 'test', 'test', 'test', 'SddggEssSsmt1dDGdsdgsdGfdDgRdfgsDd==', '2016-09-03 03:02:28', '2016-09-04 14:57:20', '1');

-- ----------------------------
-- Table structure for t_admusergroup
-- ----------------------------
DROP TABLE IF EXISTS `t_admusergroup`;
CREATE TABLE `t_admusergroup` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `groupname` varchar(255) NOT NULL DEFAULT '',
  `des` varchar(255) NOT NULL DEFAULT '',
  `createtime` datetime NOT NULL,
  `updatetime` datetime NOT NULL,
  `isdel` tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of t_admusergroup
-- ----------------------------
INSERT INTO `t_admusergroup` VALUES ('1', '超级管理员', '超级管理员', '2016-09-03 03:01:42', '2016-09-03 03:01:42', '1');
INSERT INTO `t_admusergroup` VALUES ('2', 'test', 'test', '2016-09-04 07:55:52', '2016-09-04 07:55:52', '1');
INSERT INTO `t_admusergroup` VALUES ('3', 'test2', 'test2', '2016-09-04 06:41:06', '2016-09-04 06:41:06', '0');

-- ----------------------------
-- Table structure for t_group_role_rel
-- ----------------------------
DROP TABLE IF EXISTS `t_group_role_rel`;
CREATE TABLE `t_group_role_rel` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `groupid` bigint(20) NOT NULL DEFAULT '0',
  `roleid` bigint(20) NOT NULL DEFAULT '0',
  `isdel` tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=310 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of t_group_role_rel
-- ----------------------------
INSERT INTO `t_group_role_rel` VALUES ('291', '2', '1', '1');
INSERT INTO `t_group_role_rel` VALUES ('292', '2', '18', '1');
INSERT INTO `t_group_role_rel` VALUES ('293', '2', '19', '1');
INSERT INTO `t_group_role_rel` VALUES ('294', '2', '20', '1');
INSERT INTO `t_group_role_rel` VALUES ('295', '2', '21', '1');
INSERT INTO `t_group_role_rel` VALUES ('296', '2', '2', '1');
INSERT INTO `t_group_role_rel` VALUES ('297', '2', '3', '1');
INSERT INTO `t_group_role_rel` VALUES ('298', '2', '6', '1');
INSERT INTO `t_group_role_rel` VALUES ('299', '2', '13', '1');
INSERT INTO `t_group_role_rel` VALUES ('300', '2', '15', '1');
INSERT INTO `t_group_role_rel` VALUES ('301', '2', '16', '1');
INSERT INTO `t_group_role_rel` VALUES ('302', '2', '4', '1');
INSERT INTO `t_group_role_rel` VALUES ('303', '2', '25', '1');
INSERT INTO `t_group_role_rel` VALUES ('304', '2', '26', '1');
INSERT INTO `t_group_role_rel` VALUES ('305', '2', '27', '1');
INSERT INTO `t_group_role_rel` VALUES ('306', '2', '28', '1');
INSERT INTO `t_group_role_rel` VALUES ('307', '2', '29', '1');
INSERT INTO `t_group_role_rel` VALUES ('308', '2', '30', '1');
INSERT INTO `t_group_role_rel` VALUES ('309', '2', '42', '1');

-- ----------------------------
-- Table structure for t_role
-- ----------------------------
DROP TABLE IF EXISTS `t_role`;
CREATE TABLE `t_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) DEFAULT '0',
  `name` varchar(255) NOT NULL DEFAULT '',
  `roleurl` varchar(255) NOT NULL DEFAULT '',
  `ismenu` tinyint(4) NOT NULL DEFAULT '0',
  `des` varchar(255) NOT NULL DEFAULT '',
  `module` varchar(50) NOT NULL DEFAULT '',
  `action` varchar(50) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of t_role
-- ----------------------------
INSERT INTO `t_role` VALUES ('0', null, 'Root', '3123', '1', '根节点', '', '');
INSERT INTO `t_role` VALUES ('1', '0', '公共权限', '', '1', '公共权限 所有账号都应该有', '', '');
INSERT INTO `t_role` VALUES ('2', '0', '账号管理', '', '0', '账号管理目录', '', '');
INSERT INTO `t_role` VALUES ('3', '2', '管理员管理', 'admuser/list', '0', '', 'AdmUserController', 'List');
INSERT INTO `t_role` VALUES ('4', '2', '管理员组管理', 'admusergroup/list', '0', '', 'AdmUserGroupController', 'List');
INSERT INTO `t_role` VALUES ('5', '2', '权限管理', 'role/list', '0', '', 'RoleController', 'List');
INSERT INTO `t_role` VALUES ('6', '3', '获取管理员列表', 'admuser/list', '1', '', 'AdmUserController', 'Gridlist');
INSERT INTO `t_role` VALUES ('13', '3', '查看所有管理员', 'admuser/gridlist', '1', '', 'AdmUserController', 'Gridlist');
INSERT INTO `t_role` VALUES ('15', '3', '进入添加管理员', 'admuser/toaddadmuser', '1', '进入添加管理员页面', 'AdmUserController', 'Toaddadmuser');
INSERT INTO `t_role` VALUES ('16', '3', '添加管理员', 'admuser/addadmuser', '1', '执行添加管理员操作', 'AdmUserController', 'Addadmuser');
INSERT INTO `t_role` VALUES ('18', '1', '进入欢迎页', '/welcome', '1', '进入欢迎页', 'MainController', 'Welcome');
INSERT INTO `t_role` VALUES ('19', '1', '展示导航页面', '/leftMenu', '1', '展示导航页面', 'MainController', 'LeftMenu');
INSERT INTO `t_role` VALUES ('20', '1', '展示头部信息', '/header', '1', '展示头部信息', 'MainController', 'Header');
INSERT INTO `t_role` VALUES ('21', '1', '获取菜单数据', '/loadMenu', '1', '获取菜单数据', 'MainController', 'LoadMenu');
INSERT INTO `t_role` VALUES ('25', '4', '进入添加页面', '', '1', '进入添加页面', 'AdmUserGroupController', 'Toadd');
INSERT INTO `t_role` VALUES ('26', '4', '添加管理员组', '', '1', '添加管理员组', 'AdmUserGroupController', 'Addadmusergroup');
INSERT INTO `t_role` VALUES ('27', '4', '进入修改页面', '', '1', '进入修改页面', 'AdmUserGroupController', 'Tomodify');
INSERT INTO `t_role` VALUES ('28', '4', '修改管理员组', '', '1', '修改管理员组', 'AdmUserGroupController', 'Modifyadmusergroup');
INSERT INTO `t_role` VALUES ('29', '4', '删除管理员组', '', '1', '删除管理员组', 'AdmUserGroupController', 'Delete');
INSERT INTO `t_role` VALUES ('30', '4', '获取权限树', '', '1', '添加管理员组的时候展示权限树', 'AdmUserGroupController', 'Loadtreechecked');
INSERT INTO `t_role` VALUES ('32', '5', '查询', '', '1', '查询列表', 'RoleController', 'Gridlist');
INSERT INTO `t_role` VALUES ('33', '5', '加载左侧树', '', '1', '加载左侧树', 'RoleController', 'Listtree');
INSERT INTO `t_role` VALUES ('34', '5', '进入添加页面', '', '1', '进入添加页面', 'RoleController', 'Toadd');
INSERT INTO `t_role` VALUES ('35', '5', '添加权限', '', '1', '添加权限', 'RoleController', 'Addrole');
INSERT INTO `t_role` VALUES ('36', '5', '进入修改页面', '', '1', '进入修改页面', 'RoleController', 'Tomodify');
INSERT INTO `t_role` VALUES ('37', '5', '修改权限', '', '1', '修改权限', 'RoleController', 'Modify');
INSERT INTO `t_role` VALUES ('38', '5', '删除权限', '', '1', '删除权限', 'RoleController', 'Deleterole');
INSERT INTO `t_role` VALUES ('42', '4', '查询', '', '1', '查询列表', 'AdmUserGroupController', 'Gridlist');

-- ----------------------------
-- Table structure for t_role_bak
-- ----------------------------
DROP TABLE IF EXISTS `t_role_bak`;
CREATE TABLE `t_role_bak` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) DEFAULT '0',
  `name` varchar(255) NOT NULL DEFAULT '',
  `roleurl` varchar(255) NOT NULL DEFAULT '',
  `ismenu` tinyint(4) NOT NULL DEFAULT '0',
  `des` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of t_role_bak
-- ----------------------------
INSERT INTO `t_role_bak` VALUES ('0', null, 'Root', '3123', '123', '123123');
INSERT INTO `t_role_bak` VALUES ('1', '0', '用户管污染物而', '534', '127', '345345');
INSERT INTO `t_role_bak` VALUES ('2', '0', '系统管理', '345', '127', '45345');
INSERT INTO `t_role_bak` VALUES ('4', '1', '用户管理', 'aaa', '1', '阿斯蒂芬');
INSERT INTO `t_role_bak` VALUES ('5', '2', '管理员管理', '阿斯蒂芬', '4', '阿迪飞');

-- ----------------------------
-- Table structure for t_user_group_rel
-- ----------------------------
DROP TABLE IF EXISTS `t_user_group_rel`;
CREATE TABLE `t_user_group_rel` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `userid` bigint(20) NOT NULL DEFAULT '0',
  `groupid` bigint(20) NOT NULL DEFAULT '0',
  `isdel` tinyint(4) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of t_user_group_rel
-- ----------------------------
INSERT INTO `t_user_group_rel` VALUES ('4', '1', '1', '1');
INSERT INTO `t_user_group_rel` VALUES ('7', '2', '2', '1');
