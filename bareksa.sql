DROP TABLE IF EXISTS `topic`;
CREATE TABLE `topic` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS `news`;
CREATE TABLE `news` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `topic_id` int unsigned NOT NULL,
  `name` varchar(200) NOT NULL,
  `status` varchar(30) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_topic` (`topic_id`),
  CONSTRAINT `FK_topic` FOREIGN KEY (`topic_id`) REFERENCES `topic` (`id`)
);
DROP TABLE IF EXISTS `tags`;
CREATE TABLE `tags` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS `tags_news`;
CREATE TABLE `tags_news` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `news_id` int unsigned NOT NULL,
  `tag_id` int unsigned NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `FK_news_tags` (`news_id`),
  KEY `FK_tags_news` (`tag_id`),
  CONSTRAINT `FK_news_tags` FOREIGN KEY (`news_id`) REFERENCES `news` (`id`),
  CONSTRAINT `FK_tags_news` FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`)
);

