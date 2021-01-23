package logic

import (
	"context"

	"iron-go/category/rpc/category"
	"iron-go/category/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetTopListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTopListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopListLogic {
	return &GetTopListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTopListLogic) GetTopList(in *category.GetTopListReq) (*category.GetTopListResp, error) {
	list, err := l.svcCtx.CategoryModel.GetTopList()
	if err != nil {
		l.Logger.Error("GetTopList", "error", err)
		return nil, err
	}
	if list == nil {
		return nil, err
	}
	var ret = make([]*category.CategoryItem, 0)
	for _, item := range list {
		ret = append(ret, &category.CategoryItem{
			Id:   int32(item.Id),
			Name: item.Name,
		})
	}
	return &category.GetTopListResp{
		Data: ret,
	}, nil
}
