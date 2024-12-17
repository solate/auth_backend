package middleware

import (
	"net/http"
)

// middleware/rbac.go
func RBACMiddleware(requiredPerms []string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			claims := r.Context().Value("claims").(*JWTClaims)

			// 检查用户权限
			if !hasPermission(claims.RoleIds, requiredPerms) {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		}
	}
}
