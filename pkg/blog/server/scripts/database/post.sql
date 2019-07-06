CREATE DATABASE goblog

CREATE TABLE blogs (
  `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `title` TEXT,
  `content` TEXT,
  `create_at` DATETIME,
  `modified_at` DATETIME,
)
