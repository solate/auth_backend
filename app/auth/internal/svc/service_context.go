package svc

import (
	"auth/app/auth/internal/config"
	"auth/app/auth/internal/middleware"

	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config         config.Config
	PingMiddleware rest.Middleware // manual added
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		PingMiddleware: middleware.NewPingMiddleware().Handle, // manual added
	}
}
