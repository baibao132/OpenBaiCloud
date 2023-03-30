package svc

import (
	"BaiCloud/internal/config"
	"BaiCloud/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     model.UserModel
	FileInfoModel model.FileInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		UserModel:     model.NewUserModel(conn, c.CacheRedis),
		FileInfoModel: model.NewFileInfoModel(conn, c.CacheRedis),
	}
}
