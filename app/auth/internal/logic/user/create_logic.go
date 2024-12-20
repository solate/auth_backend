package user

import (
	"context"

	"auth/app/auth/internal/svc"
	"auth/app/auth/internal/types"
	"auth/pkg/common"
	"auth/pkg/constants"
	"auth/pkg/ent/role"
	"auth/pkg/ent/user"
	"auth/pkg/utils/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建用户
func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateUserReq) (resp *types.UserInfoResponse, err error) {

	// 1. 检查用户名是否已存在
	exists, err := l.svcCtx.Orm.User.Query().
		Where(user.Name(req.Name)).
		Exist(l.ctx)
	if err != nil {
		l.Error("CreateUserQuery User.Query error", err)
		return nil, err
	}

	if exists {
		l.Error("CreateUserQuery User.Exist error", err)
		return nil, xerr.NewErrMsg("用户名已存在")
	}
	// 2. 创建用户
	pwdStr := common.GenPwd(constants.UserPwdDefault)
	salt := common.GenRandomString(8)
	createUser, err := l.svcCtx.Orm.User.Create().
		SetNo(common.GenNo("U", l.svcCtx.Redis)).
		SetName(req.Name).
		SetPhone(req.Phone).
		SetEmail(req.Email).
		SetGender(req.Gender).
		SetPwdHashed(common.GenHashedPwd(pwdStr, salt)).
		SetPwdSalt(salt).
		SetDisableStatus(constants.DisableStatusEnable).
		SetParentDisableStatus(constants.DisableStatusEnable).
		SetCompany(req.Name).
		Save(l.ctx)
	if err != nil {
		l.Error("CreateUserQuery User.Save error", err)
		return nil, err
	}

	// 3. 设置角色
	if len(req.Roles) > 0 {
		roles, err := l.svcCtx.Orm.Role.Query().
			Where(role.IDIn(req.Roles...)).
			All(l.ctx)
		if err != nil {
			l.Error("CreateUserQuery Role.Query error", err)
			return nil, err
		}
		_, err = l.svcCtx.Orm.User.Update().AddRoles(roles...).Save(l.ctx)
		if err != nil {
			l.Error("CreateUserQuery Role.Update().AddRoles error", err)
			return nil, err
		}
	}

	// 4. 返回用户信息
	userInfo := &types.UserInfoResponse{
		ID:    createUser.ID,
		Uid:   createUser.ID,
		Name:  createUser.Name,
		No:    createUser.No,
		Phone: createUser.Phone,
		Email: createUser.Email,
		// Gender: createUser.Gender,
		// Icon:   createUser.Icon,
		// Roles:  createUser.Edges.Roles,
		Status: createUser.DisableStatus,
	}

	return userInfo, nil
}
