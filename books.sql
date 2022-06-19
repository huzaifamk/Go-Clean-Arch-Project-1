CREATE DATABASE  IF NOT EXISTS `books`;
USE `books`;

DROP TABLE IF EXISTS `books`;

CREATE TABLE `books` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  `content` longtext COLLATE utf8_unicode_ci NOT NULL,
  `author_name` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

LOCK TABLES `books` WRITE;
INSERT INTO `books` VALUES (10,'Physics','<p>Theory and Practicals</p>','Sultan Ahmed','2022-06-20 13:50:19'),(20,'Maths','<p>Past Papers Only</p>','Ali Khan','2022-06-20 13:55:47'),(30,'Biology','<p>Past Papers with Practicals</p>','Taimoor Khan','2022-06-20 14:05:03'),(40,'Chemistry','<p>Model Papers with Theory</p>','Uzair Ahmed K.','2022-06-20 14:15:29');

UNLOCK TABLES;
