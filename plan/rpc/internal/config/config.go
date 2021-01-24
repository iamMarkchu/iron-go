package config

import "github.com/tal-tech/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	ZrpcUser zrpc.RpcClientConf
	Mysql    struct {
		DataSource string
	}
}
