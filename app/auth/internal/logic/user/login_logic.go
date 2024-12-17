package user

import (
	"context"
	"errors"
	"time"

	"auth/app/auth/internal/svc"
	"auth/app/auth/internal/types"
	"auth/pkg/ent"

	"github.com/dgrijalva/jwt-go"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginResp, error) {
	// 查询用户
	user, err := l.svcCtx.EntClient.User.Query().
		Where(
			user.Username(req.Username),
			user.Status(1),
		).First(l.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("用户名或密码错误")
		}
		return nil, err
	}

	// 验证密码
	if !utils.ComparePassword(user.Password, req.Password) {
		return nil, errors.New("用户名或密码错误")
	}

	// 获取用户角色
	roles, err := user.QueryRoles().All(l.ctx)
	if err != nil {
		return nil, err
	}

	roleNames := make([]string, len(roles))
	roleIds := make([]int64, len(roles))
	for i, role := range roles {
		roleNames[i] = role.Name
		roleIds[i] = role.ID
	}

	// 生成 token
	token, err := l.generateToken(user.ID, user.Username, roleIds)
	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		Token:    token,
		UserId:   user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
		Roles:    roleNames,
	}, nil
}

func (l *LoginLogic) generateToken(userId int64, username string, roleIds []int64) (string, error) {
	claims := jwt.MapClaims{
		"userId":   userId,
		"username": username,
		"roleIds":  roleIds,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
}
