package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"iron-go/category/rpc/internal/config"
	"iron-go/common/model"
)

type ServiceContext struct {
	Config        config.Config
	CategoryModel model.IronCategoriesModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	cm := model.NewIronCategoriesModel(conn)
	return &ServiceContext{
		Config:        c,
		CategoryModel: cm,
	}
}
