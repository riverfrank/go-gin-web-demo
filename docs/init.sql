
CREATE TABLE `student` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `stud_name` varchar(32) NOT NULL DEFAULT '' COMMENT '学生姓名',
  `stud_age` int(11) NOT NULL DEFAULT '0' COMMENT '学生年龄',
  `stud_sex` varchar(8) NOT NULL DEFAULT '' COMMENT '学生性别',
  `create_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_stud_name` (`stud_name`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='学生表';


CREATE TABLE `teacher` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `teacher_name` varchar(32) NOT NULL DEFAULT '' COMMENT '老师姓名',
  `create_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='老师表';


CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(30) DEFAULT NULL,
  `password` varchar(40) DEFAULT NULL,
  `name` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

INSERT INTO ginhello.student (stud_name,stud_age,stud_sex,create_time,update_time) VALUES
('张三',18,'男','2020-04-17 11:17:12','2020-04-17 11:17:12')
,('李四',19,'男','2020-04-17 11:17:12','2020-04-17 11:17:12')
,('如花',16,'女','2020-04-17 11:17:12','2020-04-17 11:17:12')
;

INSERT INTO ginhello.teacher (teacher_name,create_time,update_time) VALUES
('苍老师','2020-04-17 11:17:12','2020-04-17 11:17:12')
;
