package logic

import (
	"context"

	"iron-go/internal/svc"
	"iron-go/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type Register22Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegister22Logic(ctx context.Context, svcCtx *svc.ServiceContext) Register22Logic {
	return Register22Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Register22Logic) Register22(req types.RegisterReq) (*types.RegisterResp, error) {
	// todo: add your logic here and delete this line

	return &types.RegisterResp{}, nil
}
