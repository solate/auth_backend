package permission

import (
	"context"

	"auth/app/auth/internal/svc"
	"auth/app/auth/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取权限树
func NewPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionLogic {
	return &PermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PermissionLogic) Permission() (resp *types.PermissionTreeResp, err error) {
	// // 1. 获取所有权限
	// permissions, err := l.svcCtx.Orm.Permission.Query().
	// 	Order(ent.Asc(permission.FieldSort)).
	// 	All(l.ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// // 2. 构建权限树
	// permMap := make(map[int64]*types.PermissionTree)
	// var rootNodes []*types.PermissionTree

	// // 先构建map
	// for _, p := range permissions {
	// 	node := &types.PermissionTree{
	// 		ID:       p.ID,
	// 		ParentID: p.ParentID,
	// 		Name:     p.Name,
	// 		Code:     p.Code,
	// 		Type:     p.Type,
	// 		Path:     p.Path,
	// 		Method:   p.Method,
	// 		Sort:     p.Sort,
	// 		Status:   p.Status,
	// 		Children: make([]*types.PermissionTree, 0),
	// 	}
	// 	permMap[p.ID] = node

	// 	if p.ParentID == 0 {
	// 		rootNodes = append(rootNodes, node)
	// 	}
	// }

	// // 构建树形结构
	// for _, p := range permissions {
	// 	if p.ParentID > 0 {
	// 		if parent, ok := permMap[p.ParentID]; ok {
	// 			parent.Children = append(parent.Children, permMap[p.ID])
	// 		}
	// 	}
	// }

	// return &types.PermissionTreeResp{
	// 	List: rootNodes,
	// }, nil

	return nil, nil
}
