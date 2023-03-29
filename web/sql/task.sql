CREATE DATABASE IF NOT EXISTS task;

USE task;

CREATE TABLE IF NOT EXISTS `tasks` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '任务 ID',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '任务名称',
  `desc` varchar(200) NOT NULL DEFAULT '' COMMENT '任务描述',
  `type` int(11) NOT NULL DEFAULT '0' COMMENT '任务类型：新手、每日等',
  `event` int(11) NOT NULL DEFAULT '0' COMMENT '任务事件：签到、胜利等',
  `target` int(11) NOT NULL DEFAULT '1' COMMENT '任务所需完成次数',
  `reward` json NOT NULL DEFAULT '{}' COMMENT '奖励', 
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `status` tinyint(3) NOT NULL DEFAULT '0' COMMENT '状态：0-下线 1-上线 -1-删除'
  `jump_type` int(11) NOT NULL DEFAULT '' COMMENT '跳转类：0-uri 1-个人信息页面',
  `jump_uri` varchar(300) NOT NULL DEFAULT '' COMMENT '跳转地址',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY(`id`),
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '任务表';

create table if not exists `user_tasks` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 记录 ID',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户 ID',
  `task_id` int(11) NOT NULL DEFAULT '0' COMMENT '任务 ID',
  `task_type` int(11) NOT NULL DEFAULT '0' COMMENT '任务类型：新手、每日等',
  `task_event` int(11) NOT NULL DEFAULT '0' COMMENT '任务事件',
  `task_target` int(11) NOT NULL DEFAULT '1' COMMENT '任务目标',
  `task_progress` int(11) NOT NULL DEFAULT '0' COMMENT '任务进度',
  `status` tinyint(3) NOT NULL DEFAULT '0' COMMENT '状态：0-未完成 1-已完成未领取 2-已领取'
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY(`id`),
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '任务进度表';
