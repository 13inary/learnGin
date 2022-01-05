package token

import "github.com/gin-gonic/gin"

func RoutersToken(e *gin.Engine) {
	gro1 := e.Group("/token")
	{
		// set
		gro1.GET("get", getGet)

		// get
		gro1.GET("set", getSet)
	}
}
