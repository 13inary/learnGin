package asynch

import (
	"github.com/gin-gonic/gin"
)

func RoutersAsynch(e *gin.Engine) {
	gro1 := e.Group("/sync")
	{
		// sync
		gro1.GET("/log", getSync)
	}
}
