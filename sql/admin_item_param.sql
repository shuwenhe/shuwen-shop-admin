/*
 Navicat MySQL Data Transfer

 Source Server         : shuwenshopadmin
 Source Server Type    : MySQL
 Source Server Version : 50647
 Source Host           : 62.234.11.179:3306
 Source Schema         : shuwenshopadmin

 Target Server Type    : MySQL
 Target Server Version : 50647
 File Encoding         : 65001

 Date: 07/11/2020 20:04:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for item_param
-- ----------------------------
DROP TABLE IF EXISTS `admin_item_param`;
CREATE TABLE `admin_item_param` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `item_cat_id` bigint(20) DEFAULT NULL COMMENT '商品类目',
  `param_data` text COMMENT '参数数据，格式为json格式',
  `created` datetime DEFAULT NULL COMMENT '创建时间',
  `updated` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
	KEY `item_cat_id` (`item_cat_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品规格参数';

SET FOREIGN_KEY_CHECKS = 1;
