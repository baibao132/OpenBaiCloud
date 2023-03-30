package logic

import (
	"context"

	"BaiCloud/internal/svc"
	"BaiCloud/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpleadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpleadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpleadFileLogic {
	return &UpleadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpleadFileLogic) UpleadFile(req *types.UpleadReq) error {
	// todo: add your logic here and delete this line

	return nil
}
