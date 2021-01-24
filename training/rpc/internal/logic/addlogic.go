package logic

import (
	"context"
	errcode "iron-go/common/library/error"
	"iron-go/common/model"
	"time"

	"iron-go/training/rpc/internal/svc"
	"iron-go/training/rpc/training"

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

func (l *AddLogic) Add(in *training.AddReq) (*training.AddResp, error) {
	if in.GetUId() == 0 {
		return nil, errcode.ErrNoUid
	}
	if in.GetPlanId() == 0 {
		return nil, errcode.ErrNoPlanId
	}
	now := time.Now().Unix()
	if in.GetStartTime() <= uint64(now) {
		return nil, errcode.ErrStartTime
	}
	if in.GetEndTime() <= uint64(now) {
		return nil, errcode.ErrEndTime
	}
	if in.GetStartTime() >= in.GetEndTime() {
		return nil, errcode.ErrStartEnd
	}
	data := model.IronTrainings{
		TrainingDate: time.Time{},
		PlanId:       int64(in.GetPlanId()),
		StartTime:    time.Unix(int64(in.GetStartTime()), 0),
		EndTime:      time.Unix(int64(in.GetEndTime()), 0),
		Description:  in.GetDescription(),
		Status:       model.StatusOk,
		UserId:       int64(in.GetUId()),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	l.svcCtx.TrainingModel.Insert(data)
	return &training.AddResp{}, nil
}
