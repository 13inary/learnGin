package upload

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func uploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(500, "upload error")
	}

	c.SaveUploadedFile(file, file.Filename)

	c.String(http.StatusOK, file.Filename)
}

func uploadPicture(c *gin.Context) {
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(500, "upload error")
	}

	// file size
	if header.Size > 1024*1024*2 {
		c.String(500, "size too big")
	}

	// file type
	if header.Header.Get("Context-Type") != "image/png" {
		c.String(500, "only allow picture")
	}

	c.SaveUploadedFile(header, header.Filename)

	c.String(http.StatusOK, header.Filename)
}

func uploadMore(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	files := form.File["file"]

	for _, file := range files {
		err = c.SaveUploadedFile(file, file.Filename)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}
	}

	c.String(200, fmt.Sprintf("%d", len(files)))
}
