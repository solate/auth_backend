package role

import (
	"context"

	"auth/app/auth/internal/svc"
	"auth/app/auth/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取角色列表
func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.RoleListReq) (resp *types.RoleListResp, err error) {
	// // 1. 构建查询条件
	// query := l.svcCtx.DB.Role.Query()
	// if req.Name != "" {
	// 	query = query.Where(role.NameContains(req.Name))
	// }
	// if req.Status > 0 {
	// 	query = query.Where(role.Status(req.Status))
	// }

	// // 2. 获取总数
	// total, err := query.Count(l.ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// // 3. 分页查询
	// roles, err := query.
	// 	Limit(int(req.PageSize)).
	// 	Offset(int((req.Page - 1) * req.PageSize)).
	// 	Order(ent.Desc(role.FieldCreatedAt)).
	// 	WithPermissions().
	// 	All(l.ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// // 4. 构建返回数据
	// list := make([]*types.RoleInfo, 0)
	// for _, r := range roles {
	// 	permissions := make([]int64, 0)
	// 	for _, p := range r.Edges.Permissions {
	// 		permissions = append(permissions, p.ID)
	// 	}
	// 	list = append(list, &types.RoleInfo{
	// 		ID:          r.ID,
	// 		Name:        r.Name,
	// 		Description: r.Description,
	// 		Status:      r.Status,
	// 		Permissions: permissions,
	// 	})
	// }

	// return &types.RoleListResp{
	// 	Total: total,
	// 	List:  list,
	// }, nil

	return nil, nil
}
