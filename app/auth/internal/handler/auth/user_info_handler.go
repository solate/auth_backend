package auth

import (
	"net/http"

	"auth/app/auth/internal/logic/auth"
	"auth/app/auth/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取当前用户信息
func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := auth.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
