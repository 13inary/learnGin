package asynch

import (
	"log"

	"github.com/gin-gonic/gin"
)

func getSync(c *gin.Context) {
	co := c.Copy()
	go func() {
		log.Println(co.Request.URL.Path)

	}()
}
