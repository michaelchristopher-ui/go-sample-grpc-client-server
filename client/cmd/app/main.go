package main

import (
	"flag"
	"fmt"
	apihttp "go-sample-grpc-client/api/http"
	"go-sample-grpc-client/internal/conf"
	"go-sample-grpc-client/internal/pkg/platform/pbclient"
	"go-sample-grpc-client/internal/pkg/transport"
)

func main() {
	//Initialize server
	cfgPath := flag.String("configpath", "config.yaml", "path to config file")
	flag.Parse()

	err := conf.Init(*cfgPath)
	if err != nil {
		panic(fmt.Errorf("error parsing config. %w", err))
	}
	srv := transport.NewServer()

	//Initialize services
	pbclient := pbclient.NewPBClient()

	//Register APIs
	apihttp.API(srv.GetEcho(), pbclient)

	//Start server
	srv.StartServer()
}
