package community

import (
	"github.com/DopamineNone/bubblePro/src/controller"
	"github.com/gin-gonic/gin"
)

func Register(g *gin.Engine) {
	g.GET("/community", controller.ListCommunityHandler)
	g.GET("/community/:cid", controller.GetCommunityHandler)
}
