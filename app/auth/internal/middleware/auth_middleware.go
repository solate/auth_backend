package middleware

import (
	"auth/app/auth/internal/config"
	"auth/pkg/ent"
	"auth/pkg/ent/permission"
	"auth/pkg/ent/role"
	"auth/pkg/ent/user"
	"auth/pkg/utils/jwt"
	"auth/pkg/utils/xerr"
	"context"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type AuthMiddleware struct {
	Config    config.Config
	EntClient *ent.Client
}

func NewAuthMiddleware(c config.Config, client *ent.Client) *AuthMiddleware {
	return &AuthMiddleware{
		Config:    c,
		EntClient: client,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. 获取token
		auth := r.Header.Get("Authorization")
		if auth == "" {
			httpx.Error(w, xerr.NewErrCodeMsg(http.StatusUnauthorized, "未授权访问"))
			return
		}

		// 处理Bearer token
		parts := strings.SplitN(auth, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			httpx.Error(w, xerr.NewErrCodeMsg(http.StatusUnauthorized, "无效的token格式"))
			return
		}

		// 2. 验证token
		claims, err := jwt.ParseToken(parts[1])
		if err != nil {
			httpx.Error(w, xerr.NewErrCodeMsg(http.StatusUnauthorized, "token无效"))
			return
		}

		// 3. 获取用户权限
		permissions, err := m.getUserPermissions(r.Context(), claims.UserId)
		if err != nil {
			logx.Errorf("获取用户权限失败: %v", err)
			httpx.Error(w, xerr.NewErrCodeMsg(http.StatusInternalServerError, "系统错误"))
			return
		}

		// 4. 检查权限
		if !m.checkPermission(r, permissions) {
			httpx.Error(w, xerr.NewErrCodeMsg(http.StatusForbidden, "无访问权限"))
			return
		}

		// 5. 将用户信息存入context
		ctx := context.WithValue(r.Context(), "userId", claims.UserId)
		ctx = context.WithValue(ctx, "phone", claims.Phone)
		ctx = context.WithValue(ctx, "roleIds", claims.RoleIds)
		ctx = context.WithValue(ctx, "permissions", permissions)

		// next(w, r.WithContext(ctx))

		next(w, r)
	}
}

// getUserPermissions 获取用户的所有权限
func (m *AuthMiddleware) getUserPermissions(ctx context.Context, userId int) ([]string, error) {
	// 1. 查询用户角色
	user, err := m.EntClient.User.Query().
		Where(user.ID(userId)).
		WithRoles(func(q *ent.RoleQuery) {
			q.Where(role.StatusEQ(1))
		}).
		First(ctx)

	if err != nil {
		return nil, err
	}

	roles := user.Edges.Roles

	// 2. 查询角色关联的权限
	permissionMap := make(map[string]struct{})
	for _, role := range roles {
		perms, err := role.QueryPermissions().
			Where(permission.StatusEQ(1)).
			All(ctx)
		if err != nil {
			return nil, err
		}

		for _, perm := range perms {
			permissionMap[perm.Code] = struct{}{}
		}
	}

	// 3. 转换为字符串数组
	permissions := make([]string, 0, len(permissionMap))
	for perm := range permissionMap {
		permissions = append(permissions, perm)
	}

	return permissions, nil
}

// checkPermission 检查用户是否有访问权限
func (m *AuthMiddleware) checkPermission(r *http.Request, permissions []string) bool {
	// 1. 获取路由需要的权限
	requiredPerm := m.getRequiredPermission(r)
	if requiredPerm == "" {
		return true // 如果路由没有配置权限要求，默认放行
	}

	// 2. 检查用户是否有该权限
	for _, perm := range permissions {
		if perm == requiredPerm {
			return true
		}
	}

	return false
}

// getRequiredPermission 获取路由需要的权限
func (m *AuthMiddleware) getRequiredPermission(r *http.Request) string {
	// 这里可以通过路由信息获取所需权限
	// 例如：可以通过路由注解或配置文件定义权限要求
	// 示例实现：
	path := r.URL.Path
	method := r.Method

	// 权限映射示例
	permissionMap := map[string]string{
		"GET:/api/v1/users":    "user:read",
		"POST:/api/v1/users":   "user:create",
		"PUT:/api/v1/users":    "user:update",
		"DELETE:/api/v1/users": "user:delete",
		"GET:/api/v1/roles":    "role:read",
		"POST:/api/v1/roles":   "role:create",
		"PUT:/api/v1/roles":    "role:update",
		"DELETE:/api/v1/roles": "role:delete",
	}

	return permissionMap[method+":"+path]
}
