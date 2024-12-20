package auth

import (
	"context"

	"auth/app/auth/internal/svc"
	"auth/app/auth/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 刷新Token
func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshTokenLogic) RefreshToken(req *types.RefreshTokenReq) (resp *types.LoginResp, err error) {
	// // 1. 验证刷新令牌
	// claims, err := utils.ParseToken(req.RefreshToken)
	// if err != nil {
	// 	return nil, errors.New("无效的刷新令牌")
	// }

	// // 2. 检查令牌是否在黑名单中
	// exists, err := l.svcCtx.Redis.Exists(l.ctx, "token_blacklist:"+req.RefreshToken).Result()
	// if err != nil || exists > 0 {
	// 	return nil, errors.New("令牌已失效")
	// }

	// // 3. 生成新的访问令牌和刷新令牌
	// accessToken, err := utils.GenerateToken(claims.UserID, claims.Username, time.Hour*24)
	// if err != nil {
	// 	return nil, err
	// }

	// refreshToken, err := utils.GenerateToken(claims.UserID, claims.Username, time.Hour*24*7)
	// if err != nil {
	// 	return nil, err
	// }

	// return &types.LoginResp{
	// 	AccessToken:  accessToken,
	// 	RefreshToken: refreshToken,
	// 	TokenType:    "Bearer",
	// 	ExpiresIn:    24 * 3600,
	// }, nil

	return nil, nil
}
