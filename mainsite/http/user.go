package http

import (
	"github.com/gin-gonic/gin"
)

func users(c *gin.Context) {
	users, err := userSvc.Users()
	if err != nil {
		JSONOutput(c, nil, err)
		return
	}
	JSONOutput(c, users, nil)
}

func usersByRPC(c *gin.Context) {
	users, err := userSvc.UsersByRPC()
	if err != nil {
		JSONOutput(c, nil, err)
		return
	}
	JSONOutput(c, users, nil)
}
