DROP TABLE IF EXISTS `task`;
CREATE TABLE `task` (
  id INTEGER(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  priority VARCHAR(1) NOT NULL,
  user_id integer(11) NOT NULL,
  description VARCHAR(256) NOT NULL,
  date DATE NOT NULL,
  time BIGINT,
  done tinyint(1) NOT NULL DEFAULT 0,
  canceled tinyint(1) NOT NULL DEFAULT 0,
  INDEX (user_id, date)
);
