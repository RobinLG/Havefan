
-- ----------------------------
-- Table structure for `tb_order`
-- ----------------------------
DROP TABLE IF EXISTS `tb_order`;
CREATE TABLE `tb_order` (
  `txhash` varchar(100) NOT NULL,
  `dishes` varchar(300) NOT NULL,
  `price` varchar(20) NOT NULL,
  `location` varchar(20) NOT NULL,
  `address` varchar(60) NOT NULL,
  `mobile` varchar(15) NOT NULL,
  PRIMARY KEY (`txhash` )
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;

