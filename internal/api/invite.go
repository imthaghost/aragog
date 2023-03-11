package api

import (
	"context"

	pb "github.com/imthaghost/aragog/rpc/aragog"
)

// InviteUser ...
func (a *asusServer) InviteUser(c context.Context, req *pb.UserReq) (*pb.InviteResp, error) {

	return &pb.InviteResp{}, nil
}
