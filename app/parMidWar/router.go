package parMidWar

import (
	"testGin/middleWare"

	"github.com/gin-gonic/gin"
)

func RoutersPart(e *gin.Engine) {
	gro1 := e.Group("/part")
	{
		// user part middleWare
		gro1.GET("/part", middleWare.Par(), getPart)

	}
}
