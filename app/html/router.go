package html

import (
	"github.com/gin-gonic/gin"
)

func RoutersHtml(e *gin.Engine) {
	gro1 := e.Group("/html")
	{
		// return html
		gro1.GET("/index", getHtml)

		// redirect
		gro1.GET("/redirect", getRedirect)
	}
}
