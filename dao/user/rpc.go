package user

import (
	"context"

	pb "github.com/yangjian/api/grpc/proto/user"
	"github.com/yangjian/lib/log"
)

//UsersByRPC get all the user by grpc.
func (d *Dao) UsersByRPC() (res []*pb.User, err error) {
	resp, err := d.rpc.Users(context.Background(), &pb.Empty{})
	if err != nil {
		log.Errorf("d.rpc.Users error(%v)", err)
		return
	}
	res = resp.Users
	return
}
