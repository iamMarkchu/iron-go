package logic

import (
	"context"
	errcode "iron-go/common/library/error"
	"iron-go/common/model"
	"strings"

	"iron-go/user/rpc/internal/svc"
	"iron-go/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	username := strings.TrimSpace(in.GetUserName())
	password := strings.TrimSpace(in.GetPassword())
	if len(username) == 0 {
		return nil, errcode.ErrNameError
	}
	if len(password) == 0 {
		return nil, errcode.ErrPassword
	}
	info, err := l.svcCtx.UserModel.FindOneByUserName(username)
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
	if in.GetPassword() != info.Password {
		return nil, errcode.ErrWrongPassword
	}
	return &user.LoginResp{
		UId: uint64(info.Id),
	}, nil
}
