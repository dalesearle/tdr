DROP TABLE IF EXISTS tdr.account;
CREATE TABLE tdr.account (
  accountId INT(5) AUTO_INCREMENT,
  email VARCHAR(25) UNIQUE NOT NULL,
  fname VARCHAR(15) NOT NULL,
  lastAccessed DATETIME DEFAULT now(),
  lname VARCHAR(15) NOT NULL,
  pwd varchar(64) NOT NULL,
  hasAccess BOOLEAN NOT NULL DEFAULT FALSE,
  rights BIGINT UNSIGNED NOT NULL DEFAULT 1,
  PRIMARY KEY (accountId),
  INDEX USING BTREE (email)
)ENGINE=InnoDB;
