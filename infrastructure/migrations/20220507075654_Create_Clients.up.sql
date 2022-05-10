DROP TABLE IF EXISTS `Clients`;

CREATE TABLE `Clients` (
   `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'クライアント識別子',
   `secret` varchar(400) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'シークレット',
   `createdAt` datetime NOT NULL COMMENT 'レコード作成日',
   `updatedAt` datetime DEFAULT NULL COMMENT 'レコード更新日',
   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
