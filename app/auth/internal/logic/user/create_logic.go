package user

import (
	"context"

	"auth/app/auth/internal/svc"
	"auth/app/auth/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateUserReq) (resp *types.UserInfo, err error) {
	// todo: add your logic here and delete this line

	l.Debug("debug")
	l.Error("error")
	l.Info("info")
	l.Slow("slow")

	l.svcCtx.Redis.Set("key", "value111")
	l.Info(l.svcCtx.Redis.Get("key"))

	return
}
