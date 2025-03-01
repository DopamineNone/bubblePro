package middleware

import (
	"github.com/DopamineNone/bubblePro/src/utils/auth"
	"github.com/DopamineNone/bubblePro/src/utils/echo"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	ContextUserIDKey = "userID"
)

func AuthMiddleware(c *gin.Context) {
	// get token
	rawToken := c.Request.Header.Get("Authorization")
	if len(rawToken) == 0 {
		echo.SendErrorResponse(c, "auth token is required")
		c.Abort()
		return
	}
	splits := strings.SplitN(rawToken, " ", 2)
	if splits[0] != "Bearer" {
		echo.SendErrorResponse(c, "invalid authorization header")
		c.Abort()
		return
	}

	tokenString := splits[1]
	// parse token
	claim, err := auth.ParseJwtToken(tokenString)
	if err != nil {
		echo.SendErrorResponse(c, err.Error())
		c.Abort()
		return
	}

	// store user id
	c.Set(ContextUserIDKey, claim.UserID)
}
