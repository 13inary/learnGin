package upload

import (
	"github.com/gin-gonic/gin"
)

func RoutersHtml(e *gin.Engine) {
	gro1 := e.Group("/html")
	{
		// file
		gro1.GET("/file", uploadFile)

		// picture
		gro1.POST("/picture", uploadPicture)

		// more file
		gro1.POST("/more", uploadMore)
	}
}
