package role

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

// 更新角色
func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateRoleReq) (resp *types.RoleInfo, err error) {
	// // 1. 检查角色是否存在
	// role, err := l.svcCtx.DB.Role.Query().
	// 	Where(role.ID(req.ID)).
	// 	First(l.ctx)
	// if err != nil {
	// 	return nil, errors.New("角色不存在")
	// }

	// // 2. 检查新角色名是否与其他角色冲突
	// if req.Name != role.Name {
	// 	exists, err := l.svcCtx.DB.Role.Query().
	// 		Where(role.And(
	// 			role.Name(req.Name),
	// 			role.IDNEQ(req.ID),
	// 		)).
	// 		Exist(l.ctx)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	if exists {
	// 		return nil, errors.New("角色名已存在")
	// 	}
	// }

	// // 3. 更新角色信息
	// role, err = role.Update().
	// 	SetName(req.Name).
	// 	SetDescription(req.Description).
	// 	ClearPermissions().
	// 	Save(l.ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// // 4. 更新权限
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
