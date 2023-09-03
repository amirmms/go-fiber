package main

import (
	"fiber_test/configs"
	"fiber_test/database"
	"fiber_test/pkg"
	"fiber_test/pkg/middlewares"
	"fiber_test/router"
)

func main() {
	conf := configs.Init()

	app := pkg.InitFiber(conf)

	middlewares.InitDefaultMiddlewares(app, conf)

	database.ConnectDB(conf)

	router.InitApiRoutes(app, conf)

	pkg.ServeFiber(app, conf)
}
