package svc

import (
	"auth/app/auth/internal/config"
	"auth/app/auth/internal/middleware"
	"auth/pkg/ent"
	"auth/pkg/ent/migrate"
	"auth/pkg/utils/cache"
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config          config.Config
	Orm             *ent.Client
	Redis           *redis.Redis
	AuthMiddleware  rest.Middleware
	PermissionCache *cache.PermissionCache
}

func NewServiceContext(c config.Config) *ServiceContext {

	client := initOrm(c)
	rdb := initRedis(c)

	return &ServiceContext{
		Config:          c,
		Orm:             client,
		Redis:           initRedis(c),
		AuthMiddleware:  middleware.NewAuthMiddleware(c, client).Handle,
		PermissionCache: cache.NewPermissionCache(rdb),
	}
}

// initOrm
func initOrm(c config.Config) *ent.Client {
	ops := make([]ent.Option, 0)
	if c.ShowSQL {
		ops = append(ops, ent.Debug())
	}
	client, err := ent.Open(dialect.MySQL, c.DataSource, ops...)
	if err != nil {
		logx.Errorf("ent.Open error: %v", err)
		panic(err)
	}
	if err := client.Schema.Create(context.Background(),
		migrate.WithDropIndex(true),
	); err != nil {
		logx.Errorf("ent.Schema.Create error: %v", err)
		panic(err)
	}

	return client
}

func initRedis(c config.Config) *redis.Redis {
	return redis.New(fmt.Sprintf("%s:%d", c.Redis.Host, c.Redis.Port), func(r *redis.Redis) {
		r.Type = c.Redis.Type
		r.Pass = c.Redis.Pass
	})
}
