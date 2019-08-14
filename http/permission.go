package http

import (
	"log"

	"github.com/gin-gonic/gin"
	"go-admin/model/permission"
)

func addUser(c *gin.Context) {
	var err error
	v := new(permission.User)
	if err = c.Bind(v); err != nil {
		log.Printf("addUser c.Bind v(%+v) error(%v)", v, err)
		return
	}

	err = perSvc.AddUser(c, v)
	if err != nil {
		JSONOutput(c, nil, err)
		return
	}
	JSONOutput(c, nil, nil)
}
