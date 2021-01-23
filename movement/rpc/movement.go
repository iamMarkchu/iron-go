// Code generated by goctl. DO NOT EDIT!
// Source: movement.proto

package main

import (
	"flag"
	"fmt"

	"iron-go/movement/rpc/internal/config"
	"iron-go/movement/rpc/internal/server"
	"iron-go/movement/rpc/internal/svc"
	"iron-go/movement/rpc/movement"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/movement.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewMovementServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		movement.RegisterMovementServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}