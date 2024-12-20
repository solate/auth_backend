package auth

import (
	"context"

	"auth/app/auth/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 退出登录
func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout() error {
	// 1. 获取当前用户ID
	// userId, err := l.ctx.Value("userId").(int64)
	// if err != nil {
	// 	return errors.New("未登录")
	// }

	// // 2. 将token加入黑名单
	// token := l.ctx.Value("token").(string)
	// err = l.svcCtx.Redis.Set(l.ctx, "token_blacklist:"+token, userId, time.Hour*24).Err()
	// if err != nil {
	// 	return err
	// }

	return nil
}
