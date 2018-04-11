package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetRouters 初始化router
func GetRouters(router *gin.Engine) *gin.Engine {
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	return router
}
