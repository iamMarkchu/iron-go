package logic

import (
	"context"
	"iron-go/common/model"
	"strings"
	"time"

	errcode "iron-go/common/library/error"
	"iron-go/user/rpc/internal/svc"
	"iron-go/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegistreResp, error) {
	username := strings.TrimSpace(in.GetUserName())
	password := strings.TrimSpace(in.GetPassword())
	rePassword := strings.TrimSpace(in.GetRePassword())
	// mobile := strings.TrimSpace(in.GetMobile())
	if len(username) == 0 {
		return nil, errcode.ErrNameError
	}
	if len(password) == 0 || len(rePassword) == 0 {
		return nil, errcode.ErrPassword
	}
	if password != rePassword {
		return nil, errcode.ErrNotEqualPassword
	}
	// 用户名是否存在
	info, err := l.svcCtx.UserModel.FindOneByUserName(username)
	if err != nil {
		if err != model.ErrNotFound {
			return nil, err
		}
	}
	if info != nil {
		return nil, errcode.ErrUserExisit
	}
	data := model.IronUsers{
		UserName:  in.GetUserName(),
		Password:  in.GetPassword(),
		Status:    model.StatusOk,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	res, err := l.svcCtx.UserModel.Insert(data)
	if err != nil {
		return nil, errcode.ErrCreate
	}
	userId, err := res.LastInsertId()
	if err != nil {
		return nil, errcode.ErrCreate
	}
	return &user.RegistreResp{
		UId: uint64(userId),
	}, nil
}
