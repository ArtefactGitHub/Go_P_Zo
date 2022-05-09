DROP TABLE IF EXISTS `UserTokens`;

CREATE TABLE `UserTokens` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ユーザートークン識別子',
    `user_id` int(11) NOT NULL COMMENT 'ユーザーの識別子',
    `token` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'トークン',
    `expiredAt` datetime NOT NULL COMMENT '有効期限日',
    `createdAt` datetime NOT NULL COMMENT '作成日',
    `updatedAt` datetime DEFAULT NULL COMMENT '更新日',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
