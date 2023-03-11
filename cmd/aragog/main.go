package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/imthaghost/aragog/config"
	"github.com/imthaghost/aragog/internal/bus/rabbitmq"
	"github.com/imthaghost/aragog/internal/clients"
	"github.com/imthaghost/aragog/internal/errors/sentry"
	"github.com/imthaghost/aragog/internal/logger/zap"
)

var startup time.Time

func main() {
	startup = time.Now()
	rand.Seed(time.Now().UTC().UnixNano())

	configService := config.New{}
	configService.Load()
	cfg := configService.Get()

	// init error monitor service
	errorMonitorService := sentry.NewService(&cfg)

	// init logger service
	ls := zap.NewService()

	httpClients, err := clients.SetupClients(&cfg)
	if err != nil {

		ls.Error("Failed to initialize clients")
	}

	// init bus service
	busService := rabbitmq.NewService(&cfg, ls, errorMonitorService, httpClients.Monopoly)
	busService.Consume()

	// start message
	startMsg := fmt.Sprintf("starting http server took %v", time.Since(startup))
	ls.Msg(startMsg)

	// init TradingView service
	//tvService := crawler.New(cfg)
	//
	////grpc server
	//apiServer := api.NewGRPCServer(busService, tvService)
	//
	//lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//var opts []grpc.ServerOption
	//
	//grpcServer := grpc.NewServer(opts...)
	//pb.RegisterAragogServer(grpcServer, apiServer)
	//// listen and server
	//grpcServer.Serve(lis)
	log.Printf("starting http server took %v", time.Since(startup))

}
