package router

import (
	"github.com/JREAMLU/j-gin/config"
	"github.com/JREAMLU/j-gin/controller"
	"github.com/JREAMLU/j-gin/middleware"

	"github.com/gin-gonic/gin"
)

// GetRouters init router
func GetRouters(router *gin.Engine, conf *config.HelloConfig) *gin.Engine {
	// hello world
	router.GET("/", middleware.Middle(), controller.NewHelloController(conf).World)

	return router
}
