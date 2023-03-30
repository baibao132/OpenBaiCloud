package logic

import (
	"context"

	"BaiCloud/internal/svc"
	"BaiCloud/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhotoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPhotoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhotoLogic {
	return &PhotoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PhotoLogic) Photo(req *types.DetailReq) (resp *types.DetailReply, err error) {
	// todo: add your logic here and delete this line

	return
}
