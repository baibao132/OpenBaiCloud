package logic

import (
	"context"

	"BaiCloud/internal/svc"
	"BaiCloud/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateShareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateShareLogic {
	return &CreateShareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateShareLogic) CreateShare(req *types.ShareReq) (resp *types.ShareReply, err error) {
	// todo: add your logic here and delete this line

	return
}
