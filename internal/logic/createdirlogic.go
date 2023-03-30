package logic

import (
	"BaiCloud/internal"
	"BaiCloud/internal/svc"
	"BaiCloud/internal/types"
	"BaiCloud/model"
	"context"
	"errors"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDirLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDirLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDirLogic {
	return &CreateDirLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateDirLogic) CreateDir(req *types.CreateDirReq) error {
	// todo: add your logic here and delete this line
	name := internal.JwtVerify(req.AccessToken, l.svcCtx, l.ctx)
	if name == "" {
		return errors.New("accessToken错误")
	}
	if strings.Index(req.DirName, "/") >= 0 || strings.Index(req.DirName, "\\") >= 0 {
		return errors.New("文件夹命名错误")
	}
	FileInfo, _ := l.svcCtx.FileInfoModel.FindOne(l.ctx, req.Id)
	if FileInfo != nil {
		random := "qwertyuiop123456asdfghjkl789zxcvbnm0"
		var Rd string
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 0; i < 10; i++ {
			num := r.Intn(len(random) - 1)
			Rd += random[num : num+1]

		}
		err := os.Mkdir(l.svcCtx.Config.DataPath+"/"+FileInfo.FilePath+Rd, os.ModePerm)
		if err != nil {
			return errors.New("创建文件夹失败")
		}
		l.svcCtx.FileInfoModel.Insert(l.ctx, &model.FileInfo{FileId: Rd, FileName: FileInfo.FileName + req.DirName + "/", FileSize: 0, FileType: "Path", FileUserName: name, FilePath: FileInfo.FilePath + Rd + "/"})
		return nil
	}
	return errors.New("id错误")
}
