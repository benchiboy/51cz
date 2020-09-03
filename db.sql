CREATE TABLE `comment` (
  `id` int(11) DEFAULT NULL,
  `class_code` varchar(10) DEFAULT NULL,
  `from_user_id` varchar(64) DEFAULT NULL,
  `to_user_id` varchar(64) DEFAULT NULL,
  `text` varchar(2048) DEFAULT NULL,
  `insert_time` int(11) DEFAULT NULL,
  `update_time` int(11) DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
