CREATE DATABASE goblog;

USE goblog;

CREATE TABLE posts (
  id INT AUTO_INCREMENT,
  title TEXT,
  content TEXT,
  create_at DATETIME,
  modified_at DATETIME,
  primary key (id)
);

