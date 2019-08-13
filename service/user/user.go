package user

import (
	pb "github.com/yangjian/api/grpc/proto/user"
	"github.com/yangjian/lib/log"
)

//Users get all the users
func (s *Service) Users() (res []*pb.User, err error) {
	if res, err = s.user.Users(); err != nil {
		log.Errorf("s.user.Users error(%v)", err)
	}
	return
}

//UsersByRPC get all the users
func (s *Service) UsersByRPC() (res []*pb.User, err error) {
	if res, err = s.user.UsersByRPC(); err != nil {
		log.Errorf("s.user.UsersByRPC error(%v)", err)
	}
	return
}
