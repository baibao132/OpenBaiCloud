package logic

import (
	"BaiCloud/internal"
	"BaiCloud/internal/svc"
	"BaiCloud/internal/types"
	"BaiCloud/model"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginReply, err error) {
	// todo: add your logic here and delete this line
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("参数错误")
	}

	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, req.Username)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errors.New("账户或密码错误")
	default:
		return nil, err
	}

	if userInfo.Pwd != req.Password {
		return nil, errors.New("账户或密码错误")
	}

	jwtToken, err := internal.MakeToken(userInfo.UserName)
	if err != nil {
		return nil, err
	}
	// ---end---

	return &types.LoginReply{
		Name:        userInfo.Name,
		AccessToken: jwtToken,
	}, nil
}
