package auth

import (
	"context"

	"auth/app/auth/internal/svc"
	"auth/app/auth/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登录
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// // 1. 验证用户名密码
	// user, err := l.svcCtx.Orm.User.Query().
	// 	Where(builder.Eq(sqlbuilder.Field{
	// 		Table:  "user",
	// 		Column: "username",
	// 	}, req.Username)).
	// 	First(l.ctx)
	// if err != nil {
	// 	return nil, errors.New("用户名或密码错误")
	// }

	// // 2. 验证密码
	// if !utils.ComparePassword(user.Password, req.Password) {
	// 	return nil, errors.New("用户名或密码错误")
	// }

	// // 3. 生成访问令牌和刷新令牌
	// accessToken, err := utils.GenerateToken(user.ID, user.Username, time.Hour*24)
	// if err != nil {
	// 	return nil, err
	// }

	// refreshToken, err := utils.GenerateToken(user.ID, user.Username, time.Hour*24*7)
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
