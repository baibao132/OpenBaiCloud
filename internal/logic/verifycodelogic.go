package logic

import (
	"context"

	"BaiCloud/internal/svc"
	"BaiCloud/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyCodeLogic {
	return &VerifyCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyCodeLogic) VerifyCode(req *types.VerifyReq) (resp *types.VerifyReply, err error) {
	// todo: add your logic here and delete this line

	return
}
