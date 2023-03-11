package api

import (
	"context"

	pb "github.com/imthaghost/aragog/rpc/aragog"
)

// RemoveUser ...
func (a *asusServer) RemoveUser(c context.Context, req *pb.UserReq) (*pb.RemoveResp, error) {

	return &pb.RemoveResp{}, nil
}
