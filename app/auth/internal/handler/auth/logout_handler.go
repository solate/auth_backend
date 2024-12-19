package auth

import (
	"net/http"

	"auth/app/auth/internal/logic/auth"
	"auth/app/auth/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 退出登录
func LogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := auth.NewLogoutLogic(r.Context(), svcCtx)
		err := l.Logout()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
