package main

import (
	"flag"
	"fmt"

	"mall/service/pay/api/internal/config"
	"mall/service/pay/api/internal/handler"
	"mall/service/pay/api/internal/svc"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad("etc/pay.yaml", &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
