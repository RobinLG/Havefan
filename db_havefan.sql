
-- ----------------------------
-- Table structure for `tb_order`
-- ----------------------------
DROP TABLE IF EXISTS `tb_order`;
CREATE TABLE `tb_order` (
  `txhash` int NOT NULL AUTO_INCREMENT,
  `dishes` varchar(300) NOT NULL,
  `price` varchar(20) NOT NULL,
  `location` varchar(20) NOT NULL,
  `address` varchar(60) NOT NULL,
  `mobile` varchar(15) NOT NULL,
  `time` varchar(20) NOT NULL,
  `flag` varchar(2) NOT NULL,
  PRIMARY KEY (`txhash` )
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for `tb_account`
-- ----------------------------
DROP TABLE IF EXISTS `tb_account`;
CREATE TABLE `tb_account` (
  `id` int NOT NULL AUTO_INCREMENT,
  `wallet` varchar(100) NOT NULL,
  `useridhash` varchar(100) NOT NULL,
  `mobile` varchar(15) NOT NULL,
  PRIMARY KEY (`id` )
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;

