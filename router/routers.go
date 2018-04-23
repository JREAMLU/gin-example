package router

import (
	"github.com/JREAMLU/gin-example/config"
	"github.com/JREAMLU/gin-example/controller"
	"github.com/JREAMLU/gin-example/middleware"
	"github.com/gin-gonic/gin"
)

// GetRouters init router
func GetRouters(router *gin.Engine, conf *config.HelloConfig) *gin.Engine {
	// hello world
	router.GET("/", middleware.Middle(), controller.NewHelloController(conf).World)

	return router
}
