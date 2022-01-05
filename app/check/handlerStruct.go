package check

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	// required : not null
	Age      int       `form:"age" binding:"required,gt=18"`
	Name     string    `form:"name" binding:"required"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

// getStruct have some problem TODO
func getStruct(c *gin.Context) {
	var person Person

	err := c.ShouldBind(&person)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, person)
}
