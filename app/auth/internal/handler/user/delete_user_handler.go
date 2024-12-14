package user

import (
	"net/http"

	"auth/app/auth/internal/logic/user"
	"auth/app/auth/internal/svc"
	"hdz-cloud-service/pkg/result"
)

func DeleteUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewDeleteUserLogic(r.Context(), svcCtx)
		err := l.DeleteUser()
		result.HttpResult(r, w, nil, err)
	}
}
