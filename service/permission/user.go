package permission

import (
	"log"

	"github.com/gin-gonic/gin"
	"go-admin/model/permission"
)

//Users get all the users
func (s *Service) AddUser(c *gin.Context, user *permission.User) (err error) {
	if err = s.per.AddUser(user); err != nil {
		log.Printf("s.user.Users error(%v)", err)
	}
	return
}
