package router

import (
	"github.com/DopamineNone/bubblePro/src/router/community"
	"github.com/DopamineNone/bubblePro/src/router/middleware"
	"github.com/DopamineNone/bubblePro/src/router/post"
	"github.com/DopamineNone/bubblePro/src/router/user"
	"github.com/DopamineNone/bubblePro/src/utils/echo"
	"github.com/DopamineNone/bubblePro/src/utils/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegiserRoutes register route with no authentication first then register routes with authentication
// the register order mustn't change!!!!
func RegiserRoutes(g *gin.Engine) {
	RegisterRoutesWithNoAuthentication(g)

	RegisterRoutesWithAuthentication(g, middleware.AuthMiddleware)
}

// RegisterRoutesWithNoAuthentication route registery with no authentication
func RegisterRoutesWithNoAuthentication(g *gin.Engine) {

	user.Register(g)
	g.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": 404,
		})
	})
}

// RegisterRoutesWithAuthentication
func RegisterRoutesWithAuthentication(g *gin.Engine, authMiddleware ...gin.HandlerFunc) {
	if len(authMiddleware) <= 0 {
		panic("at least one auth middleware is required")
	}
	// route registery with authentication
	g.Use(authMiddleware...)
	community.Register(g)
	post.Register(g)

	g.GET("/ping", func(context *gin.Context) {
		uid, ok := request.MustUserID(context)
		if !ok {
			return
		} else {
			echo.SendSuccessResponse(context, uid)
		}
	})
}
