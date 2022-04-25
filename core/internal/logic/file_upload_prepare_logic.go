package logic

import (
	"cloud-disk/core/models"
	"context"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadPrepareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadPrepareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadPrepareLogic {
	return &FileUploadPrepareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadPrepareLogic) FileUploadPrepare(req *types.FileUploadPrepareRequest) (resp *types.FileUploadPrepareReply, err error) {
	rp := new(models.RepositoryPool)
	has, err := l.svcCtx.Engine.Where("hash = ?", req.Md5).Get(rp)
	if err != nil {
		return
	}
	resp = new(types.FileUploadPrepareReply)
	if has {
		// 秒传成功
		resp.Identity = rp.Identity
	} else {
		// TODO : 获取该文件的UploadID,用来进行文件的分片上传
	}

	return
}
