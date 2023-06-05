package main

import (
	"flag"
	"fmt"

	"mall/service/pay/rpc/internal/config"
	"mall/service/pay/rpc/internal/server"
	"mall/service/pay/rpc/internal/svc"
	"mall/service/pay/rpc/types"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad("etc/pay.yaml", &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pay.RegisterPayServer(grpcServer, server.NewPayServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	_ = consul.RegisterService(c.ListenOn, c.Consul)
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
