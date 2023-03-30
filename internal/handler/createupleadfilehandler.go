package handler

import (
	"net/http"

	"BaiCloud/internal/logic"
	"BaiCloud/internal/svc"
	"BaiCloud/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateUpleadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateDataReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCreateUpleadFileLogic(r.Context(), svcCtx)
		resp, err := l.CreateUpleadFile(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
