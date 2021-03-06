// Code generated by goctl. DO NOT EDIT!
// Source: plan.proto

//go:generate mockgen -destination ./plan_mock.go -package planclient -source $GOFILE

package planclient

import (
	"context"

	"iron-go/plan/rpc/plan"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	PlanDetail     = plan.PlanDetail
	PlanDetailList = plan.PlanDetailList
	CreateResp     = plan.CreateResp
	GetListReq     = plan.GetListReq
	GetListResp    = plan.GetListResp
	Request        = plan.Request
	Response       = plan.Response
	CreateReq      = plan.CreateReq

	Plan interface {
		Ping(ctx context.Context, in *Request) (*Response, error)
		Create(ctx context.Context, in *CreateReq) (*CreateResp, error)
		GetList(ctx context.Context, in *GetListReq) (*GetListResp, error)
	}

	defaultPlan struct {
		cli zrpc.Client
	}
)

func NewPlan(cli zrpc.Client) Plan {
	return &defaultPlan{
		cli: cli,
	}
}

func (m *defaultPlan) Ping(ctx context.Context, in *Request) (*Response, error) {
	client := plan.NewPlanClient(m.cli.Conn())
	return client.Ping(ctx, in)
}

func (m *defaultPlan) Create(ctx context.Context, in *CreateReq) (*CreateResp, error) {
	client := plan.NewPlanClient(m.cli.Conn())
	return client.Create(ctx, in)
}

func (m *defaultPlan) GetList(ctx context.Context, in *GetListReq) (*GetListResp, error) {
	client := plan.NewPlanClient(m.cli.Conn())
	return client.GetList(ctx, in)
}
