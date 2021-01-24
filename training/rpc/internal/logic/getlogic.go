package logic

import (
	"context"

	"iron-go/training/rpc/internal/svc"
	"iron-go/training/rpc/training"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLogic) Get(in *training.GetReq) (*training.GetResp, error) {
	// todo: add your logic here and delete this line

	return &training.GetResp{}, nil
}
