package logic

import (
	"context"
	"fmt"

	"iron-go/category/rpc/category"
	"iron-go/category/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListLogic {
	return &GetListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetListLogic) GetList(in *category.Request) (*category.Response, error) {
	list, err := l.svcCtx.CategoryModel.GetAll()
	fmt.Println(list)
	if err != nil {
		logx.Error("GetList Error", "error", err)
	}
	var res = make([]*category.CategoryItem, 0)
	for _, item := range list {
		t := new(category.CategoryItem)
		t.Id = int32(item.Id)
		t.Name = item.Name
		res = append(res, t)
	}
	return &category.Response{Data: res}, nil
}
