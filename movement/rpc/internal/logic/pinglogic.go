package logic

import (
	"context"

	"iron-go/movement/rpc/internal/svc"
	"iron-go/movement/rpc/movement"

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

func (l *PingLogic) Ping(in *movement.Request) (*movement.Response, error) {
	// todo: add your logic here and delete this line

	return &movement.Response{}, nil
}
