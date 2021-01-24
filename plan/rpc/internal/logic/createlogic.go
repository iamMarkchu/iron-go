package logic

import (
	"context"
	"iron-go/common/model"
	"strings"
	"time"

	errcode "iron-go/common/library/error"
	"iron-go/plan/rpc/internal/svc"
	"iron-go/plan/rpc/plan"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *plan.CreateReq) (*plan.CreateResp, error) {
	// 判断名字
	name := strings.TrimSpace(in.GetPlanName())
	if len(name) == 0 {
		return nil, errcode.ErrNameError
	}
	// 检测uid
	if in.GetUid() == 0 {
		return nil, errcode.ErrNoUid
	}
	//todo 获取用户信息
	if len(in.GetPlanDetails()) == 0 {
		return nil, errcode.ErrNoPlanDetails
	}
	details := make([]model.IronPlanDetails, 0)
	for _, detail := range in.GetPlanDetails() {
		// 判断动作是否合法
		err := func() error {
			if detail.GetWeight() <= 0 {
				return errcode.ErrWeight
			}
			if detail.GetCount() <= 0 {
				return errcode.ErrCount
			}
			if detail.GetBreak() <= 0 {
				return errcode.ErrBreak
			}
			return nil
		}()
		if err != nil {
			return nil, err
		}
		details = append(details, model.IronPlanDetails{
			MovementId: int64(detail.GetMovementId()),
			Weight:     int64(detail.GetWeight()),
			Count:      int64(detail.GetCount()),
			Break:      int64(detail.GetBreak()),
			Status:     model.StatusOk,
			UserId:     int64(in.GetUid()),
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		})
	}
	data := model.IronPlans{
		PlanName:  in.GetPlanName(),
		Status:    model.StatusOk,
		UserId:    int64(in.GetUid()),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	planId, err := l.svcCtx.PlanModel.Create(data, details)
	if err != nil {
		return nil, err
	}
	return &plan.CreateResp{PlanId: uint64(planId)}, nil
}
