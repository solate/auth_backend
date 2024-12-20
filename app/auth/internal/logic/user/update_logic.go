package user

import (
	"context"

	"auth/app/auth/internal/svc"
	"auth/app/auth/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新用户
func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateUserReq) (resp *types.UserInfoResponse, err error) {

	// TODO 更新用户暂时没有需求
	// // 1. 检查用户是否存在
	// user, err := l.svcCtx.DB.User.Query().
	// 	Where(user.ID(req.ID)).
	// 	First(l.ctx)
	// if err != nil {
	// 	return nil, errors.New("用户不存在")
	// }

	// // 2. 检查用户名是否与其他用户冲突
	// if req.Username != user.Username {
	// 	exists, err := l.svcCtx.DB.User.Query().
	// 		Where(user.And(
	// 			user.Username(req.Username),
	// 			user.IDNEQ(req.ID),
	// 		)).
	// 		Exist(l.ctx)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	if exists {
	// 		return nil, errors.New("用户名已存在")
	// 	}
	// }

	// // 3. 更新用户信息
	// updateUser := user.Update().
	// 	SetUsername(req.Username).
	// 	SetNickname(req.Nickname).
	// 	SetEmail(req.Email).
	// 	SetPhone(req.Phone).
	// 	ClearRoles()

	// // 如果有密码更新
	// if req.Password != "" {
	// 	hashedPassword, err := utils.HashPassword(req.Password)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	updateUser.SetPassword(hashedPassword)
	// }

	// user, err = updateUser.Save(l.ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// // 4. 更新角色
	// if len(req.Roles) > 0 {
	// 	roles, err := l.svcCtx.DB.Role.Query().
	// 		Where(role.IDIn(req.Roles...)).
	// 		All(l.ctx)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	_, err = user.Update().AddRoles(roles...).Save(l.ctx)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	// return &types.UserInfo{
	// 	ID:       user.ID,
	// 	Username: user.Username,
	// 	Nickname: user.Nickname,
	// 	Email:    user.Email,
	// 	Phone:    user.Phone,
	// 	Status:   user.Status,
	// 	Roles:    req.Roles,
	// }, nil

	return nil, nil
}
