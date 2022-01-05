package parameter

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id   string `uri:"id"`
	Name string `uri:"name"`
	Age  int64  `uri:"age"`
}

func postJson(c *gin.Context) {
	// json
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if user.Name != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad user",
		})
		return
	}

	// response yaml
	c.YAML(http.StatusOK, gin.H{
		"success": "user is ok",
	})

}

func postForm(c *gin.Context) {
	// form
	var user User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// response json
	c.JSON(http.StatusOK, gin.H{
		"success": "success receive",
	})
}

func getUri(c *gin.Context) {
	// form
	var user User
	if err := c.ShouldBindUri(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
		return
	}
	fmt.Printf("user = %+v\n", user)

	// response struct json
	c.JSON(http.StatusOK, user)
}
