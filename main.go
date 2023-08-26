package main

import (
	"fiber_test/configs"
	"fiber_test/pkg"
	"fiber_test/router"
)

func main() {
	conf := configs.Init()

	app := pkg.InitFiber(conf)

	router.InitApiRoutes(app)

	pkg.ServeFiber(app, conf)
}
