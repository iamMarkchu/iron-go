package logic

import (
	"context"
	"errors"
	"iron-go/common/model"
	"strings"
	"time"

	"iron-go/category/rpc/category"
	"iron-go/category/rpc/internal/svc"

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

func (l *AddLogic) Add(in *category.AddReq) (*category.AddResp, error) {
	var (
		trimName = strings.TrimSpace(in.GetCatName())
		trimDesc = strings.TrimSpace(in.GetDescription())
		err      error
	)
	if len(trimName) == 0 {
		return nil, errors.New("种类名不能为空")
	}
	if len(trimDesc) == 0 {
		return nil, errors.New("描述不能为空")
	}
	if in.GetUid() == 0 {
		return nil, errors.New("请提供uid")
	}
	// 检测父类是否存在
	if in.GetParentId() > 0 {
		_, err = l.svcCtx.CategoryModel.FindOne(in.GetParentId())
		if err != nil {
			return nil, errors.New("查询父类出错")
		}
	}
	// 检测名字是否重复
	info, err := l.svcCtx.CategoryModel.FindOneByName(in.GetCatName())
	if err != nil {
		if err != model.ErrNotFound {
			return nil, errors.New("查询去重失败")
		}
	}
	// 同父类id下不允许重复
	if info != nil {
		if info.Id > 0 && (info.ParentId == in.GetParentId()) {
			return nil, errors.New("已存在同名数据")
		}
	}
	data := model.IronCategories{
		ParentId:    in.GetParentId(),
		Description: in.GetDescription(),
		Status:      model.StatusOk,
		UserId:      in.GetUid(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Name:        in.GetCatName(),
	}
	ret, err := l.svcCtx.CategoryModel.Insert(data)
	if err != nil {
		l.Logger.Error("CategoryModel.Insert", "error", err)
		return nil, errors.New("创建失败")
	}
	data.Id, err = ret.LastInsertId()
	if err != nil {
		l.Logger.Error("ret.LastInsertId", "error", err)
		return nil, errors.New("创建失败2")
	}
	return &category.AddResp{
		CatId: data.Id,
	}, nil
}
