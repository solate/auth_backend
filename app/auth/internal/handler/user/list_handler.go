package user

import (
	"net/http"

	"auth/app/auth/internal/logic/user"
	"auth/app/auth/internal/svc"
	"auth/app/auth/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取用户列表
func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
