package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"iron-go/common/model"
	"iron-go/plan/rpc/internal/config"
)

type ServiceContext struct {
	Config          config.Config
	PlanModel       model.IronPlansModel
	PlanDetailModel model.IronPlanDetailsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	pm := model.NewIronPlansModel(conn)
	pdm := model.NewIronPlanDetailsModel(conn)
	return &ServiceContext{
		Config:          c,
		PlanModel:       pm,
		PlanDetailModel: pdm,
	}
}
