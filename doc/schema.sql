CREATE DATABASE `dz`;

USE dz;

CREATE TABLE `users`(
	id VARCHAR(36) NOT NULL,
	name VARCHAR(36) NOT NULL,
	PRIMARY KEY(`id`),
	KEY `users_name` (`name`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;


INSERT `users` SET id='1',name='elvin';

CREATE TABLE `posts`(
	id INT(36) UNSIGNED NOT NULL,
	title VARCHAR(100) NOT NULL,
	content TEXT NOT NULL,
	PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;