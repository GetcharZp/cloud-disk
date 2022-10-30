package handler

import (
	"cloud-disk/core/models"
	"errors"
	"net/http"

	"cloud-disk/core/internal/logic"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadChunkCompleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadChunkCompleteRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		// 判断是否已达用户容量上限
		userIdentity := r.Header.Get("UserIdentity")
		ub := new(models.UserBasic)
		_, err := svcCtx.Engine.Where("identity = ?", userIdentity).Select("now_volume, total_volume").Get(ub)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		if req.Size+ub.NowVolume > ub.TotalVolume {
			httpx.Error(w, errors.New("已超出当前容量"))
			return
		}
		l := logic.NewFileUploadChunkCompleteLogic(r.Context(), svcCtx)
		resp, err := l.FileUploadChunkComplete(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
