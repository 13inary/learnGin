package check

import "github.com/gin-gonic/gin"

func RoutersCheck(e *gin.Engine) {
	gro1 := e.Group("/check")
	{
		// check struct
		gro1.GET("/struct", getStruct)

		// custom validator 1
		gro1.GET("/admin", getCustom1)
		// custom validator 2
		gro1.GET("/time", getCustom2)
	}

}
