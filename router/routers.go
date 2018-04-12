package router

import (
	"github.com/JREAMLU/gin-example/controller"
	"github.com/gin-gonic/gin"
)

// GetRouters init router
func GetRouters(router *gin.Engine) *gin.Engine {
	// hello world
	router.GET("/", controller.NewHelloController().World)

	return router
}
