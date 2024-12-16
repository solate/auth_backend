package middleware

import (
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingMiddleware struct {
}

func NewPingMiddleware() *PingMiddleware {
	return &PingMiddleware{}
}

// need to implement logic
func (m *PingMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("开始ping中间件") // 填充逻辑

		// Passthrough to next handler if need
		next(w, r)

		logx.Info("结束ping中间件") // 填充逻辑
	}
}
