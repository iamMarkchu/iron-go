package logic

import (
	"context"
	"iron-go/category/rpc/category"

	"iron-go/internal/svc"
	"iron-go/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ListCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListCategoryLogic {
	return ListCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCategoryLogic) ListCategory(req types.ListCategoryReq) (*types.ListCategoryResp, error) {
	list, err := l.svcCtx.CategoryRpcClient.GetList(l.ctx, &category.Request{Source: "hello"})
	if err != nil {
		logx.Error("l.svcCtx.CategoryRpcClient.GetList", "error", err)
	}
	var res = make([]types.CategoryItem, 0)
	for _, item := range list.GetData() {
		tt := types.CategoryItem{
			Id:   int(item.GetId()),
			Name: item.GetName(),
		}
		res = append(res, tt)
	}
	return &types.ListCategoryResp{Data: res}, nil
}
