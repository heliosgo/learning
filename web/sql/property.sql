CREATE DATABASE IF NOT EXISTS property;

USE property;

CREATE TABLE IF NOT EXISTS `user_score` (
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户 ID',
  `level` tinyint(3) NOT NULL DEFAULT '0' COMMENT '等级',
  `score` int(11) NOT NULL DEFAULT '0' COMMENT '总积分',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY(`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '用户积分等级表';

CREATE TABLE IF NOT EXISTS `level_score` (
  `level` tinyint(3) NOT NULL DEFAULT '0' COMMENT '等级',
  `score` int(11) NOT NULL DEFAULT '0' COMMENT '等级积分',
  PRIMARY KEY(`level`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '等级积分定义表';

INSERT INTO level_score(level, score)
VALUES
(0, 0),
(1, 10),
(2, 50),
(3, 100),
(4, 200),
(5, 400),
(6, 800),
(7, 1600),
(8, 3200),
(9, 6400),
(10, 10000);

CREATE TABLE IF NOT EXISTS `score_record` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '记录 ID',
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户 ID',
  `change_score` int(11) NOT NULL DEFAULT '0' COMMENT '变化分值',
  `after_score` int(11) NOT NULL DEFAULT '0' COMMENT '变化后分值',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY(`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '积分变化记录表';
