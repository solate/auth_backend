package role

import (
	"context"

	"auth/app/auth/internal/svc"
	"auth/app/auth/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除角色
func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.IDRequest) (resp bool, err error) {
	// // 1. 检查角色是否存在
	// role, err := l.svcCtx.DB.Role.Query().
	// 	Where(role.ID(req.ID)).
	// 	First(l.ctx)
	// if err != nil {
	// 	return false, errors.New("角色不存在")
	// }

	// // 2. 检查角色是否被用户使用
	// count, err := role.QueryUsers().Count(l.ctx)
	// if err != nil {
	// 	return false, err
	// }
	// if count > 0 {
	// 	return false, errors.New("角色正在使用中，无法删除")
	// }

	// // 3. 删除角色
	// err = l.svcCtx.DB.Role.DeleteOne(role).Exec(l.ctx)
	// if err != nil {
	// 	return false, err
	// }

	return true, nil
}
