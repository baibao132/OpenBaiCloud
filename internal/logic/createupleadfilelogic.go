package logic

import (
	"context"

	"BaiCloud/internal/svc"
	"BaiCloud/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpleadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUpleadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpleadFileLogic {
	return &CreateUpleadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUpleadFileLogic) CreateUpleadFile(req *types.CreateDataReq) (resp *types.CreateDataReply, err error) {
	// todo: add your logic here and delete this line

	return
}
