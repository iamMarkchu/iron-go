package logic

import (
	"context"
	"iron-go/common/library/vars"
	"iron-go/common/model"

	"iron-go/category/rpc/category"
	"iron-go/category/rpc/internal/svc"
	errcode "iron-go/common/library/error"

	"github.com/tal-tech/go-zero/core/logx"
)

type DelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLogic {
	return &DelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelLogic) Del(in *category.DelReq) (*category.DelResp, error) {
	if in.GetUid() == 0 {
		return nil, errcode.ErrNoUid
	}
	if in.GetCatId() == 0 {
		return nil, errcode.ErrNoCatId
	}
	logx.Info("Category Del", "catId", in.GetCatId(), "uid", in.GetUid())
	info, err := l.svcCtx.CategoryModel.FindOne(in.GetCatId())
	if err != nil {
		return nil, errcode.ErrQueryError
	}
	if info == nil {
		return nil, errcode.ErrQueryEmpty
	}
	// 判断状态
	if info.Status != model.StatusOk {
		return nil, errcode.ErrStatusInvalidError
	}
	// 删除
	err = l.svcCtx.CategoryModel.Delete(in.GetCatId())
	if err != nil {
		return nil, errcode.ErrDeleteError
	}
	logx.Info("Category Del Success", "catId", in.GetCatId(), "uid", in.GetUid())
	return &category.DelResp{
		Done: vars.Done,
	}, nil
}
