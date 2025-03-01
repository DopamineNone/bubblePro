package user

import (
	"github.com/DopamineNone/bubblePro/src/controller"
	"github.com/gin-gonic/gin"
)

func Register(g *gin.Engine) {
	// auth related
	g.POST("/signup", controller.SignUpHandler)
	g.POST("/signin", controller.SignInHandler)
	g.POST("/refresh", controller.RefreshHandler)
	// user related

}
