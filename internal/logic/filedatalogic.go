package logic

import (
	"BaiCloud/internal"
	"BaiCloud/internal/svc"
	"BaiCloud/internal/types"
	"bufio"
	"context"
	"errors"
	"io"
	"os"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileDataLogic {
	return &FileDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileDataLogic) FileData(req *types.DataReq) (resp *[]byte, err error) {
	// todo: add your logic here and delete this line
	name := internal.JwtVerify(req.AccessToken, l.svcCtx, l.ctx)
	if name == "" {
		return nil, errors.New("accessToken错误")
	}
	FileInfo, _ := l.svcCtx.FileInfoModel.FindOne(l.ctx, req.Id)
	if FileInfo != nil {
		inputFile, inputError := os.Open(l.svcCtx.Config.DataPath + FileInfo.FilePath)
		if inputError != nil {
			return nil, nil
		}
		inputReader := bufio.NewReader(inputFile)
		buffer := make([]byte, 1024)
		var bytes []byte
		for {
			count, err := inputReader.Read(buffer)
			if err == io.EOF {
				break
			}
			currBytes := buffer[:count]
			bytes = append(bytes, currBytes...)
		}
		return &bytes, nil
	}
	return nil, errors.New("不存在该文件")
}
