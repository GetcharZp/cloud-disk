package logic

import (
	"cloud-disk/core/define"
	"cloud-disk/core/helper"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginReply, err error) {
	// 1、从数据库中查询当前用户
	user := new(models.UserBasic)
	has, err := l.svcCtx.Engine.Where("name = ? AND password = ?", req.Name, helper.Md5(req.Password)).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("用户名或密码错误")
	}
	// 2、生成token
	token, err := helper.GenerateToken(user.Id, user.Identity, user.Name, define.TokenExpire)
	if err != nil {
		return nil, err
	}
	// 3、生成用于刷新token的token
	refreshToken, err := helper.GenerateToken(user.Id, user.Identity, user.Name, define.RefreshTokenExpire)
	if err != nil {
		return nil, err
	}
	resp = new(types.LoginReply)
	resp.Token = token
	resp.RefreshToken = refreshToken
	return
}
