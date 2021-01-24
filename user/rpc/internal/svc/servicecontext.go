package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"iron-go/common/model"
	"iron-go/user/rpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.IronUsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	um := model.NewIronUsersModel(conn)
	return &ServiceContext{
		Config:    c,
		UserModel: um,
	}
}
