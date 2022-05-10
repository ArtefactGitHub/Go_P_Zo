DROP TABLE IF EXISTS `Users`;

CREATE TABLE `Users` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Userの一意の識別子',
    `given_name` varchar(20) NOT NULL COMMENT '名前',
    `family_name` varchar(20) NOT NULL COMMENT '苗字',
    `email` varchar(100) NOT NULL COMMENT 'メールアドレス',
    `password` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'パスワード',
    `createdAt` datetime NOT NULL COMMENT 'レコード作成日',
    `updatedAt` datetime DEFAULT NULL COMMENT 'レコード更新日',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
