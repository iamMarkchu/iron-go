package logic

import (
	"context"

	"iron-go/training/rpc/internal/svc"
	"iron-go/training/rpc/training"

	"github.com/tal-tech/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *training.Request) (*training.Response, error) {
	// todo: add your logic here and delete this line

	return &training.Response{}, nil
}
