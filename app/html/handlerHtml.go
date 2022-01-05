package html

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "index/index.html", gin.H{
		"title": "i am title",
		"ce":    "110",
	})
}

func getRedirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://www.liwenzhou.com")
}
