package parameter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTest(c *gin.Context) {
	id := c.Param("id")
	age := c.Query("age")
	name := c.DefaultQuery("name", "default name")

	// response xml
	c.XML(http.StatusOK, gin.H{
		"message": "hhh",
		"id":      id,
		"name":    name,
		"age":     age,
	})
}

func postTest(c *gin.Context) {
	id := c.PostForm("id")
	name := c.DefaultPostForm("name", "default name")

	// response string
	c.String(http.StatusOK, "id is "+id+" name is "+name)
}
