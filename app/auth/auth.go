package main

import (
	"flag"
	"fmt"

	"auth/app/auth/internal/config"
	"auth/app/auth/internal/handler"
	"auth/app/auth/internal/middleware"
	"auth/app/auth/internal/svc"
	"auth/pkg/utils/response"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/auth.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 配置log, 必须放到MustNewServer之前，不然就会会导致log配置失效，默认日志级别info
	c.LoadLogConf()

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	server.Use(middleware.LoggerMiddleware)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 使用拦截器，处理返回值
	httpx.SetOkHandler(response.OkHanandler)
	httpx.SetErrorHandlerCtx(response.ErrHandler(c.Name))

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
