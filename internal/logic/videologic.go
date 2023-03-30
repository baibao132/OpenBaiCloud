package logic

import (
	"context"

	"BaiCloud/internal/svc"
	"BaiCloud/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoLogic {
	return &VideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VideoLogic) Video(req *types.DetailReq) (resp *types.DetailReply, err error) {
	// todo: add your logic here and delete this line

	return
}
