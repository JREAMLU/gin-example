package main

import (
	"github.com/JREAMLU/gin-example/router"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.New()
	g.Use(gin.Recovery(), gin.Logger())
	g = router.GetRouters(g)
	g.Run(":8001")
}
