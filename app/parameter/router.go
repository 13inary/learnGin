package parameter

import (
	"log"

	"github.com/gin-gonic/gin"
)

func RoutersBasic(e *gin.Engine) {
	log.Println("2. set group")
	gro1 := e.Group("/parameter")
	log.Println("3. set url interface")
	{
		// select
		gro1.GET("/get", getTest)
		// insert
		gro1.POST("/post", postTest)
		// update
		gro1.PUT("/put", getTest)
		// delete
		gro1.DELETE("/delete", getTest)

		// json
		gro1.POST("/json", postJson)
		// form
		gro1.POST("/form", postForm)
		// uri
		gro1.GET("/uri/:id/:name/:age", getUri)

		// protobuf
		gro1.GET("/protobuf", getProtobuf)
	}
}
