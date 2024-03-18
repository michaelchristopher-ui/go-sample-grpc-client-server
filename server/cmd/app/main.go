package main

import (
	"flag"
	"fmt"
	apigrpc "go-sample-grpc-server/api/grpc"
	"go-sample-grpc-server/internal/conf"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	//Initialize server
	//srv := transport.NewServer()
	cfgPath := flag.String("configpath", "config.yaml", "path to config file")
	flag.Parse()

	err := conf.Init(*cfgPath)
	if err != nil {
		panic(fmt.Errorf("error parsing config. %w", err))
	}
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", conf.GetConfig().Server.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcsrv := grpc.NewServer()

	//Register APIs
	apigrpc.API(grpcsrv)

	//Start server
	grpcsrv.Serve(lis)
	//srv.StartServer()
}
