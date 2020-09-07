CREATE TABLE `comments` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `class_no` varchar(10) DEFAULT '',
  `from_user_id` varchar(64) DEFAULT '',
  `to_user_id` varchar(64) DEFAULT '',
  `text` varchar(2048) DEFAULT '',
  `insert_time` int(11) DEFAULT '0',
  `update_time` int(11) DEFAULT '0',
  `version` int(11) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8



CREATE TABLE `items` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `item_no` varchar(20) DEFAULT '',
  `item_title` varchar(128) DEFAULT '',
  `item_desc` varchar(252) DEFAULT '',
  `item_pic` varchar(512) DEFAULT '',
  `item_context` varchar(2048) DEFAULT '',
  `insert_time` int(11) DEFAULT '0',
  `update_time` int(11) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8


CREATE TABLE `tests` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `item_no` varchar(10) DEFAULT '',
  `test_no` varchar(10) DEFAULT '',
  `test_desc` varchar(256) DEFAULT '',
  `test_option` varchar(256) DEFAULT '',
  `test_answer` varchar(128) DEFAULT '',
  `insert_time` int(11) DEFAULT NULL,
  `update_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8


CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(32) DEFAULT '',
  `user_name` varchar(64) DEFAULT '',
  `phone` varchar(21) DEFAULT '',
  `head_image` varchar(252) DEFAULT '',
  `insert_time` int(11) DEFAULT '0',
  `update_time` int(11) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
