package parameter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

func getProtobuf(c *gin.Context) {
	rep := []int64{int64(1), int64(2)}

	label := "label"

	data := &protoexample.Test{
		Label: &label,
		Reps:  rep,
	}

	c.ProtoBuf(http.StatusOK, data)
}
