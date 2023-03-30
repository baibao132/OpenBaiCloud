package logic

import (
	"context"

	"BaiCloud/internal/svc"
	"BaiCloud/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteShareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteShareLogic {
	return &DeleteShareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteShareLogic) DeleteShare(req *types.DataReq) error {
	// todo: add your logic here and delete this line

	return nil
}
