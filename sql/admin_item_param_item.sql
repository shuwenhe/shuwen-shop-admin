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

 Date: 07/11/2020 20:29:18
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for item_param_item
-- ----------------------------
DROP TABLE IF EXISTS `item_param_item`;
CREATE TABLE `item_param_item` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `item_id` bigint(20) DEFAULT NULL COMMENT '商品ID',
  `param_data` text COMMENT '参数数据，格式为json格式',
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
	KEY `item_id` (`item_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品规格和商品的关系表';

SET FOREIGN_KEY_CHECKS = 1;
