// Code generated by goctl. DO NOT EDIT!
// Source: category.proto

//go:generate mockgen -destination ./category_mock.go -package categoryclient -source $GOFILE

package categoryclient

import (
	"context"

	"iron-go/category/rpc/category"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	CategoryItem = category.CategoryItem
	Request      = category.Request
	Response     = category.Response

	Category interface {
		GetList(ctx context.Context, in *Request) (*Response, error)
	}

	defaultCategory struct {
		cli zrpc.Client
	}
)

func NewCategory(cli zrpc.Client) Category {
	return &defaultCategory{
		cli: cli,
	}
}

func (m *defaultCategory) GetList(ctx context.Context, in *Request) (*Response, error) {
	client := category.NewCategoryClient(m.cli.Conn())
	return client.GetList(ctx, in)
}