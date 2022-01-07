package login

import (
	"learnGin/middleWare"

	"github.com/gin-gonic/gin"
)

func RoutersLogin(e *gin.Engine) {
	gro1 := e.Group("/login")
	{
		// login by cookie
		gro1.POST("/cookie", postCookie)

		// show home by cookie
		gro1.GET("/home", middleWare.Auth(), getHome)

		// login by session

		// show home by session
	}

}
