package handler

import (
	"net/http"

	"BaiCloud/internal/logic"
	"BaiCloud/internal/svc"
	"BaiCloud/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileDataHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DataReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFileDataLogic(r.Context(), svcCtx)
		resp, err := l.FileData(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			w.Write(*resp)
			httpx.Ok(w)
		}
	}
}
