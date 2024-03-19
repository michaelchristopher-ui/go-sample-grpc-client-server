package main

import (
	"flag"
	"fmt"
	apigrpc "go-sample-grpc-server/api/grpc"
	"go-sample-grpc-server/internal/conf"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

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

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	var wg sync.WaitGroup
	//Start server
	wg.Add(1)
	go func() {
		log.Printf("Starting server at %s", lis.Addr())
		grpcsrv.Serve(lis)
		wg.Done()
	}()
	<-termChan
	log.Print("Waiting for other waitgroups to finish")
	wg.Wait()

	//Stop Server
	grpcsrv.GracefulStop()

}
