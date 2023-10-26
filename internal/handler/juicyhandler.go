package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"juicy/internal/logic"
	"juicy/internal/svc"
	"juicy/internal/types"
)

func JuicyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewJuicyLogic(r.Context(), svcCtx)
		resp, err := l.Juicy(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
