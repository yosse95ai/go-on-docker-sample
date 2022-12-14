DROP DATABASE IF EXISTS monshin;
CREATE DATABASE monshin;
USE monshin;
DROP TABLE IF EXISTS patient;

CREATE TABLE patient (
  id int unsigned NOT NULL AUTO_INCREMENT,
  pid VARCHAR(20) NOT NULL,
  name VARCHAR(20) NOT NULL,
  birthday DATE NOT NULL,
  sex VARCHAR(10) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP NOT NULL DEFAULT now() ON UPDATE now(),
  PRIMARY KEY  (id)
);

INSERT INTO patient (pid,name,birthday,sex) VALUES (123123,'山田太郎','1990-01-01','male');

CREATE TABLE p (
  id int unsigned NOT NULL AUTO_INCREMENT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY  (id)
);