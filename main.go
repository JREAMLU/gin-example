package main

import (
	"github.com/JREAMLU/j-gin/config"
	"github.com/JREAMLU/j-gin/router"
	"github.com/JREAMLU/j-gin/service"
	"github.com/JREAMLU/j-kit/http"
)

func main() {
	// load config
	conf, err := config.Load()
	if err != nil {
		panic(err)
	}

	RunHTTPService(conf)
}

// RunHTTPService run http service
func RunHTTPService(conf *config.HelloConfig) {
	ms, g, t := http.NewHTTPService(conf.Config)

	// init micro client
	service.InitMicroClient(ms.Client())

	// init http client
	service.InitHTTPClient(t)

	g = router.GetRouters(g, conf)
	g.Run(conf.Web.URL)
}
