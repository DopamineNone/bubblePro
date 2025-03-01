package echo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func SendSuccessResponse(c *gin.Context, data any) {
	c.JSON(http.StatusOK, &Response{
		Message: "success",
		Data:    data,
	})
}

func SendErrorResponse(c *gin.Context, data any) {
	c.JSON(http.StatusBadRequest, &Response{
		Message: "error",
		Data:    data,
	})
}
