package logic

import (
	"context"

	"iron-go/movement/rpc/internal/svc"
	"iron-go/movement/rpc/movement"

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

func (l *GetListLogic) GetList(in *movement.GetListReq) (*movement.GetListResp, error) {
	list, err := l.svcCtx.MovementModel.GetList(in.GetCId())
	if err != nil {
		l.Logger.Error("GetList", "error", err)
		return nil, err
	}
	if list == nil {
		return nil, err
	}
	res := make([]*movement.ListItem, 0)
	for _, item := range list {
		res = append(res, &movement.ListItem{
			Id:          item.Id,
			CId:         int32(item.CatId),
			MName:       item.Name,
			Description: item.Description,
			UpdatedAt:   item.UpdatedAt.Unix(),
			CreatedAt:   item.CreatedAt.Unix(),
		})
	}
	return &movement.GetListResp{Data: res}, nil
}
