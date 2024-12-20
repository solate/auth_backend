package user

import (
	"context"

	"auth/app/auth/internal/svc"
	"auth/app/auth/internal/types"
	"auth/pkg/common"
	"auth/pkg/ent"
	"auth/pkg/ent/predicate"
	"auth/pkg/ent/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户列表
func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.UserListReq) (resp *types.UserListResp, err error) {
	// 1. 构建查询条件
	where := make([]predicate.User, 0)
	if req.Name != "" {
		where = append(where, user.NameContains(req.Name))
	}
	if req.Phone != "" {
		where = append(where, user.Phone(req.Phone))
	}
	if req.Status > 0 {
		where = append(where, user.Status(req.Status))
	}

	// 3. 分页查询
	list, total, err := l.PageList(req.Current, req.PageSize, where)
	if err != nil {
		return nil, err
	}

	respList := make([]*types.UserInfoResponse, 0)
	// 4. 构建返回数据
	for _, u := range list {
		roles := make([]int, 0)
		for _, r := range u.Edges.Roles {
			roles = append(roles, r.ID)
		}
		respList = append(respList, &types.UserInfoResponse{
			ID:     u.ID,
			Name:   u.Name,
			Email:  u.Email,
			Phone:  u.Phone,
			Status: u.Status,
			Roles:  roles,
		})
	}

	return &types.UserListResp{
		List: respList,
		Page: &types.PageResponse{
			Total:           int32(total),
			PageSize:        int32(len(list)),
			Current:         req.Current,
			RequestPageSize: req.PageSize,
		},
	}, nil
}

func (l *ListLogic) PageList(current, pageSize int32, where []predicate.User) (list []*ent.User, total int, err error) {
	offset := common.Offset(current, pageSize)

	query := l.svcCtx.Orm.User.Query().Where(where...).Order(ent.Desc(user.FieldCreatedAt))
	total, err = query.Count(l.ctx)
	if err != nil || total == 0 {
		return
	}

	list, err = query.Offset(int(offset)).Limit(int(pageSize)).All(l.ctx)
	return
}
