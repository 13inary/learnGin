package login

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func postCookie(c *gin.Context) {
	cookie, err := c.Cookie("cookie_key")
	if err != nil {
		// key value ageSecond saveDir domain ifHttps ifJsGet
		c.SetCookie("cookie_key", "cookie_value", 60, "/", "localhost", false, true)

		c.JSON(http.StatusOK, gin.H{
			"msg": "login success",
		})
	} else {
		log.Printf("cookie = %+v\n", cookie)
		// override old cookie
		c.SetCookie("cookie_key", "cookie_value", 60, "/", "localhost", false, true)

		c.JSON(http.StatusOK, gin.H{
			"msg": "login success",
		})
	}

}

func getHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "home",
	})

}
