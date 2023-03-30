package logic

import (
	"BaiCloud/internal"
	"BaiCloud/internal/svc"
	"BaiCloud/internal/types"
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
)

type FileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileLogic {
	return &FileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileLogic) File(req *types.DetailReq) (resp *types.DetailReply, err error) {
	// todo: add your logic here and delete this line
	name := internal.JwtVerify(req.AccessToken, l.svcCtx, l.ctx)
	if name == "" {
		return nil, errors.New("accessToken错误")
	}
	FileInfo, _ := l.svcCtx.FileInfoModel.Find(l.ctx, name)
	var r []types.Detail
	for _, i := range *FileInfo {
		var t types.Detail
		t.Name = i.FileName
		t.Type = i.FileType
		t.Size = i.FileSize
		t.Id = i.FileId
		fmt.Printf(t.Name + "\n")
		r = AddPathFile(t, r, i.FileName)
	}
	return &types.DetailReply{File: r}, nil
}

func AddPathFile(detail types.Detail, r []types.Detail, Name string) []types.Detail {
	s := strings.Split(Name, "/")
	if (len(s) <= 3 && detail.Type == "Path") || (len(s) <= 2 && detail.Type != "Path") {
		detail.Name = Name
		r = append(r, detail)
		return r
	}
xh:
	for j := 1; j < len(s); j++ {
		for i, _ := range r {
			if r[i].Type == "Path" && r[i].Name == "/"+s[j]+"/" {
				r[i].Path = AddPathFile(detail, r[i].Path, Name[len(s[j])+1:])
				break xh
			}
		}
	}
	return r
}
