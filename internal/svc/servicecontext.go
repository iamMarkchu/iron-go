package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"iron-go/category/rpc/category"
	"iron-go/internal/config"
)

type ServiceContext struct {
	Config            config.Config
	CategoryRpcClient category.CategoryClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		CategoryRpcClient: category.NewCategoryClient(zrpc.MustNewClient(c.ZrpcCategory).Conn()),
	}
}
