package http

import (
	"github.com/yangjian/mainsite/conf"
	"github.com/yangjian/mainsite/service/user"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

var (
	err error
	g   = gin.New()
	// r   = gin.Default()// use Logger(), Recovery()
	userSvc *user.Service
)

//Init for service router init.
func Init(c *conf.Config) {
	pprof.Register(g)
	userSvc = user.New(c)
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

	u := g.Group("/x/web")
	{
		u.GET("/users", users)
		u.GET("/rpc/users", usersByRPC)
	}

	if err := g.Run(c.HTTP.Addr); err != nil {
		panic(err)
	}
}

//Close close conn resource
func Close() {
	//close some source //mysql redis..
	userSvc.Close()
}
