create table t2 (id int NOT NULL AUTO_INCREMENT,
name varchar(255) NOT NULL,status int NOT NULL,PRIMARY KEY(id),KEY `KEY_STATUS` (status),
INDEX `name_index` (name));