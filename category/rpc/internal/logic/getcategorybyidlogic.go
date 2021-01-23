package logic

import (
	"context"

	"iron-go/category/rpc/category"
	"iron-go/category/rpc/internal/svc"
	errcode "iron-go/common/library/error"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetCategoryByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoryByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryByIdLogic {
	return &GetCategoryByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCategoryByIdLogic) GetCategoryById(in *category.GetCategoryByIdReq) (*category.GetCategoryByIdResp, error) {
	if in.GetCatId() == 0 {
		return nil, errcode.ErrNoCatId
	}
	info, err := l.svcCtx.CategoryModel.FindOne(int64(in.GetCatId()))
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, err
	}
	res := &category.CategoryItem{
		Id:     int32(info.Id),
		Name:   info.Name,
		Status: int32(info.Status),
	}
	return &category.GetCategoryByIdResp{Data: res}, nil
}
