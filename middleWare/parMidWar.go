package middleWare

// do after globle middleWare

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Par() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("do part middleWare")
		c.Writer.Status()
	}
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("cookie_key")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "not login",
			})
			// not do after function
			c.Abort()
			return
		}

		if cookie != "cookie_value" {
			c.Abort()
			return
		}

		c.Next()
		return
	}
}
