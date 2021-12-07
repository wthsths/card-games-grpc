CREATE TABLE `minigames` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `uuid` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `player` int NOT NULL,
  `type` int NOT NULL,
  `status` int NOT NULL,
  `data` json NOT NULL,
  PRIMARY KEY (`id`),
  KEY `uuid` (`uuid`),
  KEY `player` (`player`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;