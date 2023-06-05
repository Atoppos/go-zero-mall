package main

import (
	"flag"
	"fmt"

	"mall/service/product/api/internal/config"
	"mall/service/product/api/internal/handler"
	"mall/service/product/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"

)

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad("etc/product.yaml", &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
