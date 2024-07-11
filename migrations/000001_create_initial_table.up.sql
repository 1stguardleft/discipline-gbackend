-- ----------------------------
-- Table structure for blog_auth
-- ----------------------------
DROP TABLE IF EXISTS `dscp_auth`;
CREATE TABLE `dscp_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '이름',
  `password` varchar(50) DEFAULT '' COMMENT '비밀번호',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

INSERT INTO `dscp_auth` (`id`, `username`, `password`) VALUES ('1', 'test', 'test123');