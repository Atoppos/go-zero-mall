package main

import (
	"flag"
	"fmt"
	"mall/service/user/rpc/internal/config"
	"mall/service/user/rpc/internal/server"
	"mall/service/user/rpc/internal/svc"
	"mall/service/user/rpc/types"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad("etc/user.yaml", &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	_ = consul.RegisterService(c.ListenOn, c.Consul)

	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
