DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  id INTEGER(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  email CHAR(64) UNIQUE NOT NULL,
  password CHAR(64) NOT NULL,
  deleted TINYINT(1) NOT NULL DEFAULT 0,

  INDEX (email)
);

DROP TABLE IF EXISTS `task`;
CREATE TABLE `task` (
  id INTEGER(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  priority VARCHAR(1) NOT NULL,
  user_id integer(11) NOT NULL,
  description VARCHAR(256) NOT NULL,
  date DATE NOT NULL,
  estimated_time BIGINT,
  done tinyint(1) NOT NULL DEFAULT 0,
  canceled tinyint(1) NOT NULL DEFAULT 0,

  INDEX (user_id, date)
);
