package http

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go-admin/conf"
	"go-admin/service/permission"
)

var (
	err error
	// r   = gin.Default()// use Logger(), Recovery()

	g   = gin.New()
	perSvc *permission.Service
)

//Init for service router init.
func Init(c *conf.Config) {
	pprof.Register(g)

	perSvc = permission.New(c)

	route(c)
}

func route(c *conf.Config) {
	// demo test
	t := g.Group("/test")
	{
		t.GET("/ping", ping)
		t.GET("/long_async", longAsync) //1. 异步
		t.GET("/long_sync", longSync)   //2. 同步
	}

	u := g.Group("/x/admin")
	{
		u.POST("/user/add", addUser)
		//u.POST("/user/edit", addEdit)

	}

	if err := g.Run(c.HTTP.Addr); err != nil {
		panic(err)
	}
}

//Close close conn resource
func Close() {
	//close some source //mysql redis..
	perSvc.Close()
}
