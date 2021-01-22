package logic

import (
	"context"

	"iron-go/internal/svc"
	"iron-go/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddCategoryLogic {
	return AddCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCategoryLogic) AddCategory(req types.AddCategoryReq) (*types.AddCategoryResp, error) {

	return &types.AddCategoryResp{}, nil
}
