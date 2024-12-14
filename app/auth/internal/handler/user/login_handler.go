package user

import (
	"net/http"

	"auth/app/auth/internal/logic/user"
	"auth/app/auth/internal/svc"
	"auth/app/auth/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"hdz-cloud-service/pkg/result"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		if err := svcCtx.Validate(&req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		result.HttpResult(r, w, resp, err)
	}
}
