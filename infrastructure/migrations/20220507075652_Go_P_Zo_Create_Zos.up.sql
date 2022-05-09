DROP TABLE IF EXISTS `Zos`;

CREATE TABLE `Zos` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Zoの一意の識別子',
    `achievementDate` datetime NOT NULL COMMENT '達成した日時',
    `exp` int(11) NOT NULL COMMENT '獲得経験値（0以上10000以下）',
    `categoryId` int(11) NOT NULL COMMENT 'カテゴリーID',
    `message` varchar(400) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '達成した内容、自分へのメッセージ（400文字以内）',
    `createdAt` datetime NOT NULL COMMENT 'レコード作成日',
    `updatedAt` datetime DEFAULT NULL COMMENT 'レコード更新日',
    `user_id` int(11) NOT NULL COMMENT '作成したユーザーの識別子',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
