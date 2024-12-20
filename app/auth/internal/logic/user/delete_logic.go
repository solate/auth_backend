package user

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

// 删除用户
func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.IDRequest) (resp bool, err error) {
	// TODO 删除用户暂时没有需求

	// // 1. 检查用户是否存在
	// user, err := l.svcCtx.DB.User.Query().
	// 	Where(user.ID(req.ID)).
	// 	First(l.ctx)
	// if err != nil {
	// 	return false, errors.New("用户不存在")
	// }

	// // 2. 删除用户
	// err = l.svcCtx.DB.User.DeleteOne(user).Exec(l.ctx)
	// if err != nil {
	// 	return false, err
	// }

	return true, nil
}
