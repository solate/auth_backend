package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	UserId   int64   `json:"userId"`
	Username string  `json:"username"`
	RoleIds  []int64 `json:"roleIds"`
	jwt.StandardClaims
}

// middleware/rbac.go
func RBACMiddleware(requiredPerms []string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// claims := r.Context().Value("claims").(*JWTClaims)

			// // 检查用户权限
			// if !hasPermission(claims.RoleIds, requiredPerms) {
			// 	http.Error(w, "Forbidden", http.StatusForbidden)
			// 	return
			// }

			next.ServeHTTP(w, r)
		}
	}
}
