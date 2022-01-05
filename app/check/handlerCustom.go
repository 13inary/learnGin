package check

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type User struct {
	Age     int    `form:"age" binding:"required,gt=10"`
	Name    string `form:"name" binding:"NotNullAndAdmin"`
	Address string `form:"address" binding:"required"`
}

func getCustom1(c *gin.Context) {
	var user User
	if e := c.ShouldBind(&user); e != nil {
		c.JSON(http.StatusBadRequest, user)
		return
	}

	c.JSON(http.StatusOK, user)

}

type Book struct {
	TimeIn  time.Time `form:"timeIn" binding:"required,bookabledate" time_format:"2006-01-02"`
	TimeOut time.Time `form:"timeOut" binding:"required,gtfield=TimeIn" time_format:"2006-01-02"`
}

func getCustom2(c *gin.Context) {
	var b Book
	if err := c.ShouldBindWith(&b, binding.Query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "pass",
	})
}
