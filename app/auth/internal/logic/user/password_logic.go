package user

import (
	"context"

	"auth/app/auth/internal/svc"
	"auth/app/auth/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改用户密码
func NewPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PasswordLogic {
	return &PasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PasswordLogic) Password(req *types.IDRequest) (resp bool, err error) {
	
	return
}
