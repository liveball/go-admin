package permission

import (
	"go-admin/conf"
	"go-admin/dao/permission"
)

//Service struct
type Service struct {
	c   *conf.Config
	per *permission.Dao
}

//New get service
func New(c *conf.Config) *Service {
	s := &Service{
		c:   c,
		per: permission.New(c),
	}

	return s
}

//Close close db...
func (s *Service) Close() {
	s.per.Close()
}
