package request

import (
	"errors"
	"github.com/DopamineNone/bubblePro/src/router/middleware"
	"github.com/DopamineNone/bubblePro/src/utils/echo"
	"github.com/gin-gonic/gin"
)

var (
	UserNotLoginErr = errors.New("user hasn't logined")
)

func MustUserID(c *gin.Context) (userID int64, ok bool) {
	value, ok := c.Get(middleware.ContextUserIDKey)
	if !ok {
		echo.SendErrorResponse(c, UserNotLoginErr.Error())
		c.Abort()
		return
	}
	userID, ok = value.(int64)
	if !ok {
		echo.SendErrorResponse(c, UserNotLoginErr.Error())
		c.Abort()
		return
	}
	return
}
