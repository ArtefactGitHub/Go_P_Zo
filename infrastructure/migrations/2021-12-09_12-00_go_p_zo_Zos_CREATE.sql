# ************************************************************
# Sequel Ace SQL dump
# バージョン 20016
#
# https://sequel-ace.com/
# https://github.com/Sequel-Ace/Sequel-Ace
#
# ホスト: localhost (MySQL 8.0.12)
# データベース: Go_P_Zo
# 生成時間: 2021-12-09 03:00:48 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE='NO_AUTO_VALUE_ON_ZERO', SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# テーブルのダンプ Zos
# ------------------------------------------------------------
USE Go_P_Zo;

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




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
