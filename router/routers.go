package router

import (
	"github.com/JREAMLU/gin-example/config"
	"github.com/JREAMLU/gin-example/controller"
	"github.com/gin-gonic/gin"
)

// GetRouters init router
func GetRouters(router *gin.Engine, conf *config.HelloConfig) *gin.Engine {
	// hello world
	router.GET("/", controller.NewHelloController(conf).World)

	return router
}
