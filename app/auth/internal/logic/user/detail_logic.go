package user

import (
	"context"

	"auth/app/auth/internal/svc"
	"auth/app/auth/internal/types"
	"auth/pkg/ent/user"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户详情
func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.IDRequest) (resp *types.UserInfoResponse, err error) {
	// 1. 查询用户详情
	user, err := l.svcCtx.Orm.User.Query().
		Where(user.ID(req.ID)).
		WithRoles().
		First(l.ctx)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	// 2. 构建角色列表
	roles := make([]int, 0)
	for _, r := range user.Edges.Roles {
		roles = append(roles, r.ID)
	}

	// 4. 返回用户信息
	userInfo := &types.UserInfoResponse{
		ID:    user.ID,
		Uid:   user.ID,
		Name:  user.Name,
		No:    user.No,
		Phone: user.Phone,
		Email: user.Email,
		// Gender: user.Gender,
		// Icon:   user.Icon,
		// Roles:  user.Edges.Roles,
		Status: user.DisableStatus,
	}

	return userInfo, nil
}
