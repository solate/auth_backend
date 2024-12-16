package svc

import (
	"auth/app/auth/internal/config"
	"auth/app/auth/internal/middleware"
	"auth/pkg/ent"
	"auth/pkg/ent/migrate"
	"context"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config         config.Config
	PingMiddleware rest.Middleware // manual added
	Orm            *ent.Client
}

func NewServiceContext(c config.Config) *ServiceContext {

	orm := initOrm(c)
	return &ServiceContext{
		Config:         c,
		PingMiddleware: middleware.NewPingMiddleware().Handle, // manual added
		Orm:            orm,
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
