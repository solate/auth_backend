package auth

import (
	"context"

	"auth/app/auth/internal/svc"
	"auth/app/auth/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取当前用户信息
func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	// // 1. 获取当前用户ID
	// userId, err := l.ctx.Value("userId").(int64)
	// if err != nil {
	// 	return nil, errors.New("未登录")
	// }

	// // 2. 查询用户信息
	// user, err := l.svcCtx.DB.User.Query().
	// 	Where(user.ID(userId)).
	// 	WithRoles().
	// 	First(l.ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// // 3. 获取用户角色和权限
	// var roles []string
	// var permissions []string

	// for _, role := range user.Edges.Roles {
	// 	roles = append(roles, role.Name)

	// 	perms, err := role.QueryPermissions().All(l.ctx)
	// 	if err != nil {
	// 		continue
	// 	}

	// 	for _, perm := range perms {
	// 		permissions = append(permissions, perm.Code)
	// 	}
	// }

	// return &types.UserInfoResp{
	// 	ID:          user.ID,
	// 	Username:    user.Username,
	// 	Nickname:    user.Nickname,
	// 	Avatar:      user.Avatar,
	// 	Email:       user.Email,
	// 	Phone:       user.Phone,
	// 	Status:      user.Status,
	// 	Roles:       roles,
	// 	Permissions: permissions,
	// }, nil

	return nil, nil
}
