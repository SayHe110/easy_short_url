package response

import (
	"github.com/gin-gonic/gin"
	"go/types"
	"net/http"
)

type response struct {
	Code    int
	Message string
	Data    interface{}
}

func JSON(c *gin.Context, status int, message string, data interface{}) {
	if data == nil {
		data = types.Interface{}
	}

	c.JSON(http.StatusOK, &response{
		Code:    status,
		Message: message,
		Data:    data,
	})
}
