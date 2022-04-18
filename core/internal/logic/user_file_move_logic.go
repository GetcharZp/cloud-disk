package logic

import (
	"cloud-disk/core/models"
	"context"
	"errors"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest, userIdentity string) (resp *types.UserFileMoveReply, err error) {
	//parentID
	parentData := new(models.UserRepository)
	has, err := l.svcCtx.Engine.Where("identity = ? AND user_identity = ?", req.ParentIdnetity, userIdentity).Get(parentData)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("文件夹不存在")
	}
	// 更新记录的 ParentID
	_, err = l.svcCtx.Engine.Where("identity = ?", req.Idnetity).Update(models.UserRepository{
		ParentId: int64(parentData.Id),
	})
	return
}
