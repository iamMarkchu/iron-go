package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/rest/httpx"
	"iron-go/common/library/error"
	"iron-go/internal/config"
	"iron-go/internal/handler"
	"iron-go/internal/svc"
)

var configFile = flag.String("f", "etc/iron-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandler(error.ZeroErrorHandler)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
