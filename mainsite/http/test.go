package http

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	JSONOutput(c,
		map[string]interface{}{
			"a": 111,
			"b": "vvcvv",
		},
		nil,
	)
}

func longAsync(c *gin.Context) {
	// goroutine 中只能使用只读的上下文 c.Copy()
	cCp := c.Copy()
	go func() {
		time.Sleep(5 * time.Second)
		// 注意使用只读上下文
		log.Println("Done! in path " + cCp.Request.URL.Path)
	}()
}

func longSync(c *gin.Context) {
	time.Sleep(5 * time.Second)
	// 注意可以使用原始上下文
	log.Println("Done! in path " + c.Request.URL.Path)
}
