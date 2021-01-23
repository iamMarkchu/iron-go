package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/zrpc"
	"iron-go/category/rpc/category"
	"iron-go/common/model"
	"iron-go/movement/rpc/internal/config"
)

type ServiceContext struct {
	Config            config.Config
	CategoryRpcClient category.CategoryClient
	MovementModel     model.IronMovementsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	mm := model.NewIronMovementsModel(conn)
	return &ServiceContext{
		Config:            c,
		CategoryRpcClient: category.NewCategoryClient(zrpc.MustNewClient(c.ZrpcCategory).Conn()),
		MovementModel:     mm,
	}
}
