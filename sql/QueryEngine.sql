/*
	Author: Yuu Leii
	Data: 02/06/2022 
*/

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`(
	`id` int(11) unsigned AUTO_INCREMENT,
	`identity` varchar(255) NOT NULL UNIQUE,
	`name` varchar(40) NOT NULL,
	`password` varchar(40) NOT NULL,
	PRIMARY KEY(id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ---------------------
-- Table structrue for favorite_list
-- ---------------------
DROP TABLE IF EXISTS favorite_list;
CREATE TABLE favorite_list(
	-- id int unsigned AUTO_INCREMENT ,
	identity varchar(255) NOT NULL,
	name varchar(255) NOT NULL,

	UNIQUE KEY identity_name(identity, name),
	-- CONSTRAINT `id_name_restrict`
	FOREIGN KEY(`identity`) references `users` (`identity`) ON DELETE CASCADE ON UPDATE CASCADE
	-- PRIMARY KEY(`identity`, `name`)

	-- PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------
-- Table structrue for favorites
-- --------------------
DROP TABLE IF EXISTS favorite_info;
CREATE TABLE favorite_info (
	identity varchar(255) NOT NULL,
	name varchar(255) NOT NULL,
	doc_id int(11) NOT NULL,
	UNIQUE KEY identity_name_doc_id(identity, name, doc_id)
	-- url varchar(1024) NOT NULL, 
	-- PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

