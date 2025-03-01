package post

import (
	"github.com/DopamineNone/bubblePro/src/controller"
	"github.com/gin-gonic/gin"
)

func Register(g *gin.Engine) {
	g.POST("/post", controller.CreatePostHandler)
	g.GET("/post/:pid", controller.GetPostHandler)
}
