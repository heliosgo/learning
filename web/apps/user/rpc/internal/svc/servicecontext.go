package svc

import (
	"web/apps/property/rpc/propertyclient"
	"web/apps/user/rpc/internal/config"
	"web/apps/user/rpc/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel

	PropertyRPC propertyclient.Property
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlConn, c.CacheRedis),

		PropertyRPC: propertyclient.NewProperty(zrpc.MustNewClient(c.PropertyRPC)),
	}
}
