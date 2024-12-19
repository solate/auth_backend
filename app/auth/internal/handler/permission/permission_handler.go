package permission

import (
	"net/http"

	"auth/app/auth/internal/logic/permission"
	"auth/app/auth/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取权限树
func PermissionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := permission.NewPermissionLogic(r.Context(), svcCtx)
		resp, err := l.Permission()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
