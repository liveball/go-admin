package user

import (
	"github.com/yangjian/mainsite/conf"
	"github.com/yangjian/mainsite/dao/user"
)

//Service struct
type Service struct {
	c    *conf.Config
	user *user.Dao
}

//New get service
func New(c *conf.Config) *Service {
	s := &Service{
		c:    c,
		user: user.New(c),
	}

	return s
}

//Close close db...
func (s *Service) Close() {
	//
}
