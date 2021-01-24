package logic

import (
	"context"

	errcode "iron-go/common/library/error"
	"iron-go/plan/rpc/internal/svc"
	"iron-go/plan/rpc/plan"

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

func (l *GetListLogic) GetList(in *plan.GetListReq) (*plan.GetListResp, error) {
	if in.GetUid() == 0 {
		return nil, errcode.ErrNoUid
	}
	// todo 检测uid正确性
	list, err := l.svcCtx.PlanModel.GetListByUid(in.GetUid())
	if err != nil {
		return nil, err
	}
	if list == nil {
		return nil, err
	}
	ret := make([]*plan.PlanDetailList, 0)
	for _, ll := range list {
		ret = append(ret, &plan.PlanDetailList{
			Id:        uint64(ll.Id),
			PlanName:  ll.PlanName,
			Status:    uint32(ll.Status),
			Uid:       uint64(ll.UserId),
			CreatedAt: ll.CreatedAt.Unix(),
			UpdatedAt: ll.UpdatedAt.Unix(),
		})
	}
	pidS := make([]int64, 0)
	for _, item := range list {
		// 获取所有的planId
		pidS = append(pidS, item.Id)
	}
	detailMap, err := l.svcCtx.PlanDetailModel.BatchGetDetailsMap(pidS)
	if err != nil {
		return nil, err
	}
	if detailMap == nil {
		return nil, err
	}
	for k, rr := range ret {
		if dd, ok := detailMap[int64(rr.Id)]; ok {
			tmpDetails := make([]*plan.PlanDetail, 0)
			for _, ddd := range dd {
				tmpDetails = append(tmpDetails, &plan.PlanDetail{
					MovementId: uint64(ddd.MovementId),
					Weight:     uint32(ddd.Weight),
					Count:      uint32(ddd.Count),
					Break:      uint32(ddd.Break),
				})
			}
			ret[k].PlanDetails = tmpDetails
		}
	}
	return &plan.GetListResp{
		PlanDetailList: ret,
	}, nil
}
