package role

import (
	"context"

	"auth/app/auth/internal/svc"
	"auth/app/auth/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取角色详情
func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.IDRequest) (resp *types.RoleInfo, err error) {
	// // 1. 查询角色详情
	// role, err := l.svcCtx.DB.Role.Query().
	// 	Where(role.ID(req.ID)).
	// 	WithPermissions().
	// 	First(l.ctx)
	// if err != nil {
	// 	return nil, errors.New("角色不存在")
	// }

	// // 2. 构建权限列表
	// permissions := make([]int64, 0)
	// for _, p := range role.Edges.Permissions {
	// 	permissions = append(permissions, p.ID)
	// }

	// return &types.RoleInfo{
	// 	ID:          role.ID,
	// 	Name:        role.Name,
	// 	Description: role.Description,
	// 	Status:      role.Status,
	// 	Permissions: permissions,
	// }, nil

	return nil, nil
}
