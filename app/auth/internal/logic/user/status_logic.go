package user

import (
	"context"

	"auth/app/auth/internal/svc"
	"auth/app/auth/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改用户状态
func NewStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StatusLogic {
	return &StatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StatusLogic) Status(req *types.StatusRequest) (resp bool, err error) {
	// 1. 获取当前用户ID
	// userId, err := l.ctx.Value("userId").(int64)
	// if err != nil {
	// 	return false, errors.New("未登录")
	// }

	// 2. 更新用户状态
	_, err = l.svcCtx.Orm.User.UpdateOneID(req.ID).SetStatus(req.Status).Save(l.ctx)
	if err != nil {
		l.Logger.Error("StatusLogic.Status.UpdateOneID.Save", err)
		return false, err
	}

	return true, nil
}
