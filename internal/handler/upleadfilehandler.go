package handler

import (
	"net/http"

	"BaiCloud/internal/logic"
	"BaiCloud/internal/svc"
	"BaiCloud/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpleadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpleadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpleadFileLogic(r.Context(), svcCtx)
		err := l.UpleadFile(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
