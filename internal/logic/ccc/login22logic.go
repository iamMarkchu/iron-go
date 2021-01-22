package logic

import (
	"context"

	"iron-go/internal/svc"
	"iron-go/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type Login22Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogin22Logic(ctx context.Context, svcCtx *svc.ServiceContext) Login22Logic {
	return Login22Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Login22Logic) Login22(req types.LoginReq) (*types.LoginResp, error) {
	// todo: add your logic here and delete this line

	return &types.LoginResp{}, nil
}
