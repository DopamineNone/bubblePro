package main

import (
	"github.com/DopamineNone/bubblePro/src/config"
	"github.com/DopamineNone/bubblePro/src/infra"
	"github.com/DopamineNone/bubblePro/src/router"
	"github.com/gin-gonic/gin"
	"strconv"
)

var (
	g *gin.Engine
)

func initApp() {
	// initiate infrastructure
	infra.Init()

	// create http handler and regiser routes
	g = gin.New()
	router.RegiserRoutes(g)
}
func main() {
	initApp()

	// start http service
	if err := g.Run(":" + strconv.Itoa(config.GetConf().Port)); err != nil {
		panic(err)
	}
}
