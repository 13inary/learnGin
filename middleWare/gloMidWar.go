package middleWare

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Before() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("enter before middleware")
		c.Set("before", "middleware")
		//c.Next()
		status := c.Writer.Status()
		log.Println("end before middleware", status)
	}
}
