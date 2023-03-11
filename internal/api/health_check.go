package api

import (
	"context"

	pb "github.com/imthaghost/aragog/rpc/aragog"
)

// HealthCheck ...
func (a *asusServer) HealthCheck(c context.Context, req *pb.HealthReq) (*pb.HealthResp, error) {

	return &pb.HealthResp{Status: httpOk}, nil
}
