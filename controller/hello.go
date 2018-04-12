package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloController hello controller
type HelloController struct {
	Controller
}

// NewHelloController new hello
func NewHelloController() *HelloController {
	return &HelloController{}
}

// World world
func (h *HelloController) World(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}
