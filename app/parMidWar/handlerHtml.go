package parMidWar

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getPart(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title": "i am title",
		"ce":    "110",
	})
}
