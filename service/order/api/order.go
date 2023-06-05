package main

import (
	"flag"
	"fmt"

	"mall/service/order/api/internal/config"
	"mall/service/order/api/internal/handler"
	"mall/service/order/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	_ "github.com/dtm-labs/driver-gozero"
)

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad("etc/order.yaml", &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
