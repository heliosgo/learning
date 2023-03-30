package svc

import (
	"web/apps/task/rpc/internal/config"
	"web/apps/task/rpc/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	TasksModel model.TasksModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		TasksModel: model.NewTasksModel(sqlConn, c.CacheRedis),
	}
}
