package api

import (
	"context"

	"github.com/imthaghost/aragog/internal/bus"
	"github.com/imthaghost/aragog/internal/tradingview"
	pb "github.com/imthaghost/aragog/rpc/aragog"
)

type asusServer struct {
	Bus     bus.Service
	Crawler tradingview.Service

	pb.UnimplementedAragogServer
}

// NewGRPCServer ...
func NewGRPCServer(busService bus.Service, crawler tradingview.Service) pb.AragogServer {

	return &asusServer{
		Bus:     busService,
		Crawler: crawler,
	}
}

// Shutdown graceful shutdown of all services attached to the AragogServer
func (a *asusServer) Shutdown(ctx context.Context) error {

	return nil
}
