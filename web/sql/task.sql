CREATE DATABASE IF NOT EXISTS task;

USE task;

CREATE TABLE IF NOT EXISTS `tasks` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '任务 ID',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '任务名称',
  `desc` varchar(200) NOT NULL DEFAULT '' COMMENT '任务描述',
  `type` int(11) NOT NULL DEFAULT '0' COMMENT '任务类型：新手、每日等',
  `event` int(11) NOT NULL DEFAULT '0' COMMENT '任务事件：签到、胜利等',
  `target` int(11) NOT NULL DEFAULT '1' COMMENT '任务所需完成次数',
  `reward` json NOT NULL COMMENT '奖励', 
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `status` tinyint(3) NOT NULL DEFAULT '0' COMMENT '状态：0-下线 1-上线 -1-删除',
  `jump_type` int(11) NOT NULL DEFAULT '0' COMMENT '跳转类：0-uri 1-个人信息页面',
  `jump_uri` varchar(300) NOT NULL DEFAULT '' COMMENT '跳转地址',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY(`id`),
  KEY `type_status` (`type`, `status`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '任务表';

DELIMITER //
CREATE PROCEDURE IF NOT EXISTS `create_daily_task_table`()
BEGIN
  DECLARE nextDay varchar(20);
  DECLARE table_prefix varchar(20);
  DECLARE table_name varchar(20);
  DECLARE csql varchar(5210);

  SELECT replace(DATE_SUB(curdate(), INTERVAL -1 DAY), '-', '') INTO @nextDay;

  SET @table_prefix = 'user_daily_task_';

  SET @table_name = CONCAT(@table_prefix, @nextDay);
  
  SET @csql = CONCAT("CREATE TABLE IF NOT EXISTS ", @table_name, "(
      `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 记录 ID',
      `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户 ID',
      `task_id` int(11) NOT NULL DEFAULT '0' COMMENT '任务 ID',
      `task_type` int(11) NOT NULL DEFAULT '0' COMMENT '任务类型：新手、每日等',
      `task_event` int(11) NOT NULL DEFAULT '0' COMMENT '任务事件',
      `task_target` int(11) NOT NULL DEFAULT '1' COMMENT '任务目标',
      `task_progress` int(11) NOT NULL DEFAULT '0' COMMENT '任务进度',
      `status` tinyint(3) NOT NULL DEFAULT '0' COMMENT '状态：0-未完成 1-已完成未领取 2-已领取',
      `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
      `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
      PRIMARY KEY(`id`),
      KEY `user_id` (`user_id`),
      KEY `task_id` (`task_id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '任务进度表';");

  PREPARE create_stmt from @csql;
    EXECUTE create_stmt;
  DEALLOCATE PREPARE create_stmt;

END //
DELIMITER ;

CREATE EVENT IF NOT EXISTS `create_daily_task_table_by_day`
ON SCHEDULE EVERY 1 DAY STARTS '2023-03-30 23:00:00' 
ON COMPLETION PRESERVE
DO call create_daily_task_table();

