// Code generated by goctl. DO NOT EDIT!
// Source: movement.proto

//go:generate mockgen -destination ./movement_mock.go -package movementclient -source $GOFILE

package movementclient

import (
	"context"

	"iron-go/movement/rpc/movement"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	GetListResp = movement.GetListResp
	ListItem    = movement.ListItem
	Request     = movement.Request
	Response    = movement.Response
	AddReq      = movement.AddReq
	AddResp     = movement.AddResp
	GetListReq  = movement.GetListReq

	Movement interface {
		Ping(ctx context.Context, in *Request) (*Response, error)
		Add(ctx context.Context, in *AddReq) (*AddResp, error)
		GetList(ctx context.Context, in *GetListReq) (*GetListResp, error)
	}

	defaultMovement struct {
		cli zrpc.Client
	}
)

func NewMovement(cli zrpc.Client) Movement {
	return &defaultMovement{
		cli: cli,
	}
}

func (m *defaultMovement) Ping(ctx context.Context, in *Request) (*Response, error) {
	client := movement.NewMovementClient(m.cli.Conn())
	return client.Ping(ctx, in)
}

func (m *defaultMovement) Add(ctx context.Context, in *AddReq) (*AddResp, error) {
	client := movement.NewMovementClient(m.cli.Conn())
	return client.Add(ctx, in)
}

func (m *defaultMovement) GetList(ctx context.Context, in *GetListReq) (*GetListResp, error) {
	client := movement.NewMovementClient(m.cli.Conn())
	return client.GetList(ctx, in)
}