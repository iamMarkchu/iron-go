package logic

import (
	"context"
	"iron-go/category/rpc/category"
	"iron-go/common/model"
	"strings"
	"time"

	errcode "iron-go/common/library/error"
	"iron-go/movement/rpc/internal/svc"
	"iron-go/movement/rpc/movement"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *movement.AddReq) (*movement.AddResp, error) {
	// 判断类别
	if in.GetCId() == 0 {
		return nil, errcode.ErrNoCatId
	}
	// 调用服务判断类别状态
	cateInfo, err := l.svcCtx.CategoryRpcClient.GetCategoryById(l.ctx, &category.GetCategoryByIdReq{CatId: in.GetCId()})
	if err != nil {
		return nil, errcode.ErrGrpcErr
	}
	if cateInfo == nil {
		return nil, errcode.ErrNotFoundCate
	}
	if cateInfo.GetData().GetId() == 0 {
		return nil, errcode.ErrNotFoundCate
	}
	if cateInfo.GetData().GetStatus() == model.StatusDeleted {
		return nil, errcode.ErrCateStatusInvalidError
	}
	// 判断名字
	name := strings.TrimSpace(in.GetMName())
	if len(name) == 0 {
		return nil, errcode.ErrNameError
	}
	// 判断描述
	description := strings.TrimSpace(in.GetDescription())
	if len(description) == 0 {
		return nil, errcode.ErrNameError
	}
	// 判断重复
	mInfo, err := l.svcCtx.MovementModel.FindOneByName(in.GetMName())
	if err != nil {
		return nil, errcode.ErrQueryError
	}
	if mInfo != nil {
		if mInfo.Id > 0 {
			return nil, errcode.ErrDuplicate
		}
	}
	// 创建
	data := model.IronMovements{
		Id:          0,
		CatId:       int64(in.GetCId()),
		Name:        in.GetMName(),
		Description: in.GetDescription(),
		Status:      model.StatusOk,
		UserId:      1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	res, err := l.svcCtx.MovementModel.Insert(data)
	if err != nil {
		return nil, err
	}
	data.Id, err = res.LastInsertId()
	return &movement.AddResp{MId: int32(data.Id)}, nil
}
