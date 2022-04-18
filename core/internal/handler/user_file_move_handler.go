package handler

import (
	"net/http"

	"cloud-disk/core/internal/logic"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserFileMoveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserFileMoveRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserFileMoveLogic(r.Context(), svcCtx)
		resp, err := l.UserFileMove(&req, r.Header.Get("UserIdentity"))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
