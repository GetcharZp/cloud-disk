package logic

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepositorySaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositorySaveLogic {
	return &UserRepositorySaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositorySaveLogic) UserRepositorySave(req *types.UserRepositorySaveRequest, userIdentity string) (resp *types.UserRepositorySaveReply, err error) {
	// 判断文件是否超容量
	rp := new(models.RepositoryPool)
	_, err = l.svcCtx.Engine.Select("size").Where("identity = ?", req.RepositoryIdentity).Get(rp)
	if err != nil {
		fmt.Println(err)
		return
	}
	ub := new(models.UserBasic)
	_, err = l.svcCtx.Engine.Select("now_volume, total_volume").Where("identity = ?", userIdentity).Get(ub)
	if err != nil {
		fmt.Println(err)
		return
	}
	if ub.NowVolume+rp.Size > ub.TotalVolume {
		err = errors.New("已超出当前容量")
		return
	}

	// 更新当前容量
	_, err = l.svcCtx.Engine.Exec("UPDATE user_basic SET now_volume = now_volume + ? WHERE identity = ?", rp.Size, userIdentity)
	if err != nil {
		return
	}
	// 新增关联记录
	ur := &models.UserRepository{
		Identity:           helper.UUID(),
		UserIdentity:       userIdentity,
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Ext:                req.Ext,
		Name:               req.Name,
	}
	_, err = l.svcCtx.Engine.Insert(ur)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
