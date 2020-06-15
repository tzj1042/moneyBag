CREATE TABLE `t_money_bag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user` char(3) CHARACTER SET utf8 DEFAULT NULL COMMENT '用户',
  `amount` decimal(10,2) NOT NULL COMMENT '金额',
  `income_pay` tinyint(2) NOT NULL COMMENT '收支类型1收入2支出',
  `source` varchar(255) CHARACTER SET utf8 NOT NULL COMMENT '来源',
  `use` varchar(255) CHARACTER SET utf8 NOT NULL COMMENT '用途',
  `mark` varchar(255) CHARACTER SET utf8 NOT NULL COMMENT '备注',
  `describe` varchar(255) CHARACTER SET utf8 NOT NULL COMMENT '描述',
  `is_must` tinyint(2) NOT NULL DEFAULT '0' COMMENT '是否必须',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='我的钱包';
