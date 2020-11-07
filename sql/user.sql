create table user(
id bigint(20) not null auto_increment,
username varchar(50) not null comment '用户名',
password VARCHAR(50) NOT NULL COMMENT '密码，加密保存',
phone VARCHAR(20) DEFAULT NULL COMMENT '注册手机号',
email varchar(50) default null comment '注册邮箱',
created datetime not null,
updated datetime not null,
primary key(id),
unique key username (username) USING BTREE,
unique key phone (phone) USING BTREE,
UNIQUE KEY email (email) USING BTREE
)ENGINE=INNODB auto_increment=37 DEFAULT CHARSET=utf8 COMMENT='用户表';