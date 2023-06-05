package main

import (
	"flag"
	"fmt"

	"mall/service/order/rpc/internal/config"
	"mall/service/order/rpc/internal/server"
	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/types"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	_ "github.com/dtm-labs/driver-gozero" 
)

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad("etc/order.yaml", &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		order.RegisterOrderServer(grpcServer, server.NewOrderServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	_ = consul.RegisterService(c.ListenOn, c.Consul)
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
