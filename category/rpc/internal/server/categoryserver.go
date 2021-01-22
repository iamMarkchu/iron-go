// Code generated by goctl. DO NOT EDIT!
// Source: category.proto

package server

import (
	"context"

	"iron-go/category/rpc/category"
	"iron-go/category/rpc/internal/logic"
	"iron-go/category/rpc/internal/svc"
)

type CategoryServer struct {
	svcCtx *svc.ServiceContext
}

func NewCategoryServer(svcCtx *svc.ServiceContext) *CategoryServer {
	return &CategoryServer{
		svcCtx: svcCtx,
	}
}

func (s *CategoryServer) GetList(ctx context.Context, in *category.Request) (*category.Response, error) {
	l := logic.NewGetListLogic(ctx, s.svcCtx)
	return l.GetList(in)
}
