/*
Navicat MariaDB Data Transfer

Source Server         : new
Source Server Version : 100311
Source Host           : localhost:3306
Source Database       : class

Target Server Type    : MariaDB
Target Server Version : 100311
File Encoding         : 65001

Date: 2019-10-16 02:53:34
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for address
-- ----------------------------
DROP TABLE IF EXISTS `address`;
CREATE TABLE `address` (
  `id` int(128) NOT NULL AUTO_INCREMENT,
  `address` varchar(255) DEFAULT NULL,
  `addressDetail` varchar(255) DEFAULT NULL,
  `userID` int(128) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `userName` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of address
-- ----------------------------

-- ----------------------------
-- Table structure for book
-- ----------------------------
DROP TABLE IF EXISTS `book`;
CREATE TABLE `book` (
  `id` int(128) NOT NULL AUTO_INCREMENT,
  `bookname` varchar(255) DEFAULT NULL,
  `bookimgs` varchar(255) DEFAULT NULL,
  `bookprice` varchar(24) DEFAULT NULL,
  `bookmsgs` varchar(255) DEFAULT NULL,
  `booksalenum` int(64) DEFAULT NULL,
  `bookgz` int(128) DEFAULT NULL,
  `bookdetailimgs` varchar(255) DEFAULT NULL,
  `bookclass` varchar(32) DEFAULT NULL,
  `jpclass` int(32) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=73 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of book
-- ----------------------------
INSERT INTO `book` VALUES ('1', '精品课程推广', 'http://www.linchongyang.cn/static/img/jpkc/2.jpg', '45', '这是一部很适合自己学的课程', '12', '56', 'http://www.linchongyang.cn/static/img/kapian/bg3.jpg', '1', '1');
INSERT INTO `book` VALUES ('2', '精品课程推广', 'http://www.linchongyang.cn/static/img/jpkc/3.jpg', '45', '这是一部很适合自己学的课程', '12', '56', 'http://www.linchongyang.cn/static/img/kapian/bg3.jpg', '2', '1');
INSERT INTO `book` VALUES ('3', '精品课程推广', 'http://www.linchongyang.cn/static/img/jpkc/4.jpg', '45', '这是一部很适合自己学的课程', '12', '56', 'http://www.linchongyang.cn/static/img/kapian/bg3.jpg', '3', '2');
INSERT INTO `book` VALUES ('4', '精品课程推广', 'http://www.linchongyang.cn/static/img/jpkc/5.jpg', '45', '这是一部很适合自己学的课程', '12', '56', 'http://www.linchongyang.cn/static/img/kapian/bg3.jpg', '4', '3');
INSERT INTO `book` VALUES ('5', '精品课程推广', 'http://www.linchongyang.cn/static/img/jpkc/6.jpg', '45', '这是一部很适合自己学的课程', '12', '56', 'http://www.linchongyang.cn/static/img/kapian/bg3.jpg', '5', '4');
INSERT INTO `book` VALUES ('6', '精品课程推广', 'http://www.linchongyang.cn/static/img/jpkc/7.jpg', '45', '这是一部很适合自己学的课程', '12', '56', 'http://www.linchongyang.cn/static/img/kapian/bg3.jpg', '1', '2');
INSERT INTO `book` VALUES ('7', '精品课程推广', 'http://www.linchongyang.cn/static/img/jpkc/8.jpg', '45', '这是一部很适合自己学的课程', '12', '56', 'http://www.linchongyang.cn/static/img/kapian/bg3.jpg', '1', '3');
INSERT INTO `book` VALUES ('8', '精品课程推广', 'http://www.linchongyang.cn/static/img/jpkc/9.jpg', '45', '这是一部很适合自己学的课程', '12', '56', 'http://www.linchongyang.cn/static/img/kapian/bg3.jpg', '2', '4');
INSERT INTO `book` VALUES ('9', '精品课程推广', 'http://www.linchongyang.cn/static/img/jpkc/10.jpg', '45', '这是一部很适合自己学的课程', '12', '56', 'http://www.linchongyang.cn/static/img/kapian/bg3.jpg', '3', '2');
INSERT INTO `book` VALUES ('10', '精品课程推广', 'http://www.linchongyang.cn/static/img/jpkc/11.jpg', '45', '这是一部很适合自己学的课程', '12', '56', 'http://www.linchongyang.cn/static/img/kapian/bg3.jpg', '4', '3');
INSERT INTO `book` VALUES ('11', '精品课程推广', 'http://www.linchongyang.cn/static/img/jpkc/12.jpg', '45', '这是一部很适合自己学的课程', '12', '56', 'http://www.linchongyang.cn/static/img/kapian/bg3.jpg', '5', '2');
INSERT INTO `book` VALUES ('12', '精品课程推广', 'http://www.linchongyang.cn/static/img/jpkc/13.jpg', '45', '这是一部很适合自己学的课程', '12', '56', 'http://www.linchongyang.cn/static/img/kapian/bg3.jpg', '1', '4');
INSERT INTO `book` VALUES ('13', '精品课程推广', 'http://www.linchongyang.cn/static/img/jpkc/14.jpg', '45', '这是一部很适合自己学的课程', '12', '56', 'http://www.linchongyang.cn/static/img/kapian/bg3.jpg', '1', '2');
INSERT INTO `book` VALUES ('14', '精品课程推广', 'http://www.linchongyang.cn/static/img/jpkc/15.jpg', '45', '这是一部很适合自己学的课程', '12', '56', 'http://www.linchongyang.cn/static/img/kapian/bg3.jpg', '2', '2');
INSERT INTO `book` VALUES ('15', '精品课程推广', 'http://www.linchongyang.cn/static/img/jpkc/0.jpg', '45', '这是一部很适合自己学的课程', '12', '56', 'http://www.linchongyang.cn/static/img/kapian/bg3.jpg', '3', '3');
INSERT INTO `book` VALUES ('16', '精品课程推广', 'http://www.linchongyang.cn/static/img/jpkc/1.jpg', '45', '这是一部很适合自己学的课程', '12', '56', 'http://www.linchongyang.cn/static/img/kapian/bg3.jpg', '4', '5');
INSERT INTO `book` VALUES ('17', '精品课程推广', 'http://www.linchongyang.cn/static/img/jpkc/2.jpg', '45', '这是一部很适合自己学的课程', '12', '56', 'http://www.linchongyang.cn/static/img/kapian/bg3.jpg', '5', '6');
INSERT INTO `book` VALUES ('18', '精品课程推广', 'http://www.linchongyang.cn/static/img/jpkc/3.jpg', '45', '这是一部很适合自己学的课程', '12', '56', 'http://www.linchongyang.cn/static/img/kapian/bg3.jpg', '1', '3');
INSERT INTO `book` VALUES ('19', '精品课程推广', 'http://www.linchongyang.cn/static/img/jpkc/4.jpg', '45', '这是一部很适合自己学的课程', '12', '56', 'http://www.linchongyang.cn/static/img/kapian/bg3.jpg', '2', '2');
INSERT INTO `book` VALUES ('20', '精品课程推广', 'http://www.linchongyang.cn/static/img/jpkc/5.jpg', '45', '这是一部很适合自己学的课程', '12', '56', 'http://www.linchongyang.cn/static/img/kapian/bg3.jpg', '3', '1');

-- ----------------------------
-- Table structure for bookdetailimgs
-- ----------------------------
DROP TABLE IF EXISTS `bookdetailimgs`;
CREATE TABLE `bookdetailimgs` (
  `id` int(255) NOT NULL AUTO_INCREMENT,
  `bookid` int(255) DEFAULT NULL,
  `bookdetailimgs` varchar(3000) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of bookdetailimgs
-- ----------------------------
INSERT INTO `bookdetailimgs` VALUES ('1', '8', '[\"http://www.linchongyang.cn/static/img/jpkc/2.jpg\",\"http://www.linchongyang.cn/static/img/jpkc/2.jpg\",\"http://www.linchongyang.cn/static/img/jpkc/2.jpg\",\"http://www.linchongyang.cn/static/img/jpkc/2.jpg\"]');

-- ----------------------------
-- Table structure for cart
-- ----------------------------
DROP TABLE IF EXISTS `cart`;
CREATE TABLE `cart` (
  `id` int(128) NOT NULL AUTO_INCREMENT,
  `bookid` int(128) DEFAULT NULL,
  `userid` int(128) DEFAULT NULL,
  `booknum` int(128) DEFAULT NULL,
  `bookstatus` int(4) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='bookStatus  1:加入购物车，未购买     2：已购买';

-- ----------------------------
-- Records of cart
-- ----------------------------
INSERT INTO `cart` VALUES ('1', '2', '1', '1', '0');
INSERT INTO `cart` VALUES ('2', '8', '1', '1', '0');
INSERT INTO `cart` VALUES ('3', '8', '1', '1', '0');
INSERT INTO `cart` VALUES ('4', '6', '1', '1', '0');

-- ----------------------------
-- Table structure for command
-- ----------------------------
DROP TABLE IF EXISTS `command`;
CREATE TABLE `command` (
  `id` int(255) NOT NULL AUTO_INCREMENT,
  `goodsid` int(255) DEFAULT NULL,
  `uid` int(255) DEFAULT NULL,
  `commandmsg` varchar(255) DEFAULT NULL,
  `commandtime` varchar(120) DEFAULT NULL,
  `uhead` varchar(255) DEFAULT NULL,
  `uname` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of command
-- ----------------------------
INSERT INTO `command` VALUES ('1', '8', '1', '这个课程真的很不错', '2019-08-12', 'http://www.linchongyang.cn/static/img/jpkc/6.jpg', '啊哈哈哈');

-- ----------------------------
-- Table structure for havebuy
-- ----------------------------
DROP TABLE IF EXISTS `havebuy`;
CREATE TABLE `havebuy` (
  `id` int(128) NOT NULL AUTO_INCREMENT,
  `userID` int(255) DEFAULT NULL,
  `goodsID` int(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of havebuy
-- ----------------------------

-- ----------------------------
-- Table structure for indexchr
-- ----------------------------
DROP TABLE IF EXISTS `indexchr`;
CREATE TABLE `indexchr` (
  `id` int(30) NOT NULL AUTO_INCREMENT,
  `imgs` varchar(255) DEFAULT NULL,
  `names` varchar(60) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of indexchr
-- ----------------------------
INSERT INTO `indexchr` VALUES ('1', 'http://www.linchongyang.cn/static/img/jpkc/3.jpg', '精品课程');
INSERT INTO `indexchr` VALUES ('2', 'http://www.linchongyang.cn/static/img/jpkc/3.jpg', '精品课程');
INSERT INTO `indexchr` VALUES ('3', 'http://www.linchongyang.cn/static/img/jpkc/3.jpg', '精品课程');
INSERT INTO `indexchr` VALUES ('4', 'http://www.linchongyang.cn/static/img/jpkc/3.jpg', '精品课程');

-- ----------------------------
-- Table structure for lunimgs
-- ----------------------------
DROP TABLE IF EXISTS `lunimgs`;
CREATE TABLE `lunimgs` (
  `id` int(128) NOT NULL AUTO_INCREMENT,
  `imgs` varchar(255) DEFAULT NULL,
  `links` varchar(255) DEFAULT NULL,
  `msgs` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of lunimgs
-- ----------------------------
INSERT INTO `lunimgs` VALUES ('1', 'http://www.linchongyang.cn/static/img/kapian/bg4.jpg', 'http://www.linchongyang.cn/static/img/kapian/bg4.jpg', '精品课程推广');
INSERT INTO `lunimgs` VALUES ('2', 'http://www.linchongyang.cn/static/img/kapian/bg2.jpg', 'http://www.linchongyang.cn/static/img/kapian/bg2.jpg', '精品课程推广');
INSERT INTO `lunimgs` VALUES ('3', 'http://www.linchongyang.cn/static/img/kapian/bg3.jpg', 'http://www.linchongyang.cn/static/img/kapian/bg2.jpg', '精品课程推广');
INSERT INTO `lunimgs` VALUES ('4', 'http://www.linchongyang.cn/static/img/kapian/bg4.jpg', 'http://www.linchongyang.cn/static/img/kapian/bg2.jpg', '精品课程推广');
INSERT INTO `lunimgs` VALUES ('5', 'http://www.linchongyang.cn/static/img/kapian/bg1.jpg', 'http://www.linchongyang.cn/static/img/kapian/bg2.jpg', '精品课程推广');

-- ----------------------------
-- Table structure for savegoods
-- ----------------------------
DROP TABLE IF EXISTS `savegoods`;
CREATE TABLE `savegoods` (
  `id` int(255) NOT NULL AUTO_INCREMENT,
  `userid` int(255) DEFAULT NULL,
  `goodsid` int(255) DEFAULT NULL,
  `time` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='status  1:  收藏   2：取消收藏';

-- ----------------------------
-- Records of savegoods
-- ----------------------------
INSERT INTO `savegoods` VALUES ('2', '1', '5', '2019-10-16');
INSERT INTO `savegoods` VALUES ('4', '1', '6', '2019-10-16');

-- ----------------------------
-- Table structure for slidebar
-- ----------------------------
DROP TABLE IF EXISTS `slidebar`;
CREATE TABLE `slidebar` (
  `id` int(128) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `classid` int(30) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of slidebar
-- ----------------------------
INSERT INTO `slidebar` VALUES ('1', '高等代数', '1');
INSERT INTO `slidebar` VALUES ('2', '物理', '2');
INSERT INTO `slidebar` VALUES ('3', '线性代数', '3');
INSERT INTO `slidebar` VALUES ('4', '泛函分析', '4');
INSERT INTO `slidebar` VALUES ('5', '解析几何', '5');
INSERT INTO `slidebar` VALUES ('6', '高斯定理', '6');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(128) NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NOT NULL,
  `tel` varchar(255) NOT NULL,
  `pass` varchar(255) DEFAULT NULL,
  `status` int(12) DEFAULT NULL,
  `uname` varchar(60) DEFAULT NULL,
  PRIMARY KEY (`id`,`tel`,`email`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', '1653709489@qq.com', '13350058238', '123456', '1', '张三');
