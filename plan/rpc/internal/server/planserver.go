// Code generated by goctl. DO NOT EDIT!
// Source: plan.proto

package server

import (
	"context"

	"iron-go/plan/rpc/internal/logic"
	"iron-go/plan/rpc/internal/svc"
	"iron-go/plan/rpc/plan"
)

type PlanServer struct {
	svcCtx *svc.ServiceContext
}

func NewPlanServer(svcCtx *svc.ServiceContext) *PlanServer {
	return &PlanServer{
		svcCtx: svcCtx,
	}
}

func (s *PlanServer) Ping(ctx context.Context, in *plan.Request) (*plan.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}

func (s *PlanServer) Create(ctx context.Context, in *plan.CreateReq) (*plan.CreateResp, error) {
	l := logic.NewCreateLogic(ctx, s.svcCtx)
	return l.Create(in)
}
