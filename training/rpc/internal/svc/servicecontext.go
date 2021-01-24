package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/zrpc"
	"iron-go/common/model"
	"iron-go/plan/rpc/plan"
	"iron-go/training/rpc/internal/config"
	"iron-go/user/rpc/user"
)

type ServiceContext struct {
	Config           config.Config
	UserRpcClient    user.UserClient
	TrainingModel    model.IronTrainingsModel
	TrainingLogModel model.IronTrainingLogsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	tm := model.NewIronTrainingsModel(conn, plan.NewPlanClient(zrpc.MustNewClient(c.ZrpcPlan).Conn()))
	tlm := model.NewIronTrainingLogsModel(conn)
	return &ServiceContext{
		Config:           c,
		TrainingLogModel: tlm,
		TrainingModel:    tm,
		UserRpcClient:    user.NewUserClient(zrpc.MustNewClient(c.ZrpcUser).Conn()),
	}
}
