package user

import (
	"net/http"

	"auth/app/auth/internal/logic/user"
	"auth/app/auth/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取用户详情
func DetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewDetailLogic(r.Context(), svcCtx)
		resp, err := l.Detail()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
