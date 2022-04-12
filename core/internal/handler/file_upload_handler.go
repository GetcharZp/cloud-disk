package handler

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/models"
	"crypto/md5"
	"fmt"
	"net/http"
	"path"

	"cloud-disk/core/internal/logic"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		f, fh, err := r.FormFile("file")
		if err != nil {
			return
		}
		b := make([]byte, fh.Size)
		f.Read(b)

		/// Name string `json:"name,optional"`
		//	Ext  string `json:"ext,optional"`
		//	Size int64  `json:"size,optional"`
		//	Path string `json:"path,optional"`
		hash := fmt.Sprintf("%x", md5.Sum(b))
		// 判断是否在库中存在
		rp := new(models.RepositoryPool)
		has, err := svcCtx.Engine.Where("hash = ?", hash).Get(rp)
		if err != nil {
			return
		}
		if has {
			httpx.OkJson(w, &types.FileUploadReply{Identity: rp.Identity})
			return
		}
		// 上传文件到腾讯云
		cosPath, err := helper.CosUpload(r)
		if err != nil {
			return
		}
		req.Path = cosPath
		req.Hash = hash
		req.Name = fh.Filename
		req.Ext = path.Ext(fh.Filename)
		req.Size = fh.Size

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
