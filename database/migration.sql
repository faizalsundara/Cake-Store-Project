CREATE TABLE `cakes` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `title` varchar(255) DEFAULT NULL,
  `description` mediumtext,
  `rating` float DEFAULT NULL,
  `image` longtext,
  PRIMARY KEY (`id`)
);

INSERT INTO `cakes` (title, description, rating, image)
VALUES ('Serabi', 'terbuat dari tepung beras', 8.3, 'serabi.png');