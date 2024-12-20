package role

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

// 创建角色
func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateRoleReq) (resp *types.RoleInfo, err error) {
	// // 1. 检查角色名是否已存在
	// exists, err := l.svcCtx.DB.Role.Query().
	// 	Where(role.Name(req.Name)).
	// 	Exist(l.ctx)
	// if err != nil {
	// 	return nil, err
	// }
	// if exists {
	// 	return nil, errors.New("角色名已存在")
	// }

	// // 2. 创建角色
	// role, err := l.svcCtx.DB.Role.Create().
	// 	SetName(req.Name).
	// 	SetDescription(req.Description).
	// 	SetStatus(1).
	// 	Save(l.ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// // 3. 设置权限
	// if len(req.Permissions) > 0 {
	// 	perms, err := l.svcCtx.DB.Permission.Query().
	// 		Where(permission.IDIn(req.Permissions...)).
	// 		All(l.ctx)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	_, err = role.Update().AddPermissions(perms...).Save(l.ctx)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	// return &types.RoleInfo{
	// 	ID:          role.ID,
	// 	Name:        role.Name,
	// 	Description: role.Description,
	// 	Status:      role.Status,
	// 	Permissions: req.Permissions,
	// }, nil

	return nil, nil
}
