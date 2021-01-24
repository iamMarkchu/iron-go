package logic

import (
	"context"
	errcode "iron-go/common/library/error"
	"iron-go/common/model"

	"iron-go/user/rpc/internal/svc"
	"iron-go/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	if in.GetUId() == 0 {
		return nil, errcode.ErrNoUid
	}
	info, err := l.svcCtx.UserModel.FindOne(int64(in.GetUId()))
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errcode.ErrQueryEmpty
		}
		return nil, errcode.ErrQueryError
	}
	if info.Id == 0 {
		return nil, errcode.ErrUserNotExisit
	}
	if info.Status != model.StatusOk {
		return nil, errcode.ErrUserStatusInvalidError
	}
	ret := user.GetUserInfoResp{
		UId:       uint64(info.Id),
		UserName:  info.UserName,
		NickName:  info.NickName,
		Mobile:    info.Mobile,
		Status:    uint32(info.Status),
		CreatedAt: info.CreatedAt.Unix(),
		UpdatedAt: info.UpdatedAt.Unix(),
	}
	return &ret, nil
}
