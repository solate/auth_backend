package user

import (
	"net/http"

	"auth/app/auth/internal/logic/user"
	"auth/app/auth/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
	"hdz-cloud-service/pkg/result"
)

func GetUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetUserLogic(r.Context(), svcCtx)
		resp, err := l.GetUser()
		result.HttpResult(r, w, resp, err)
	}
}
