package router

import (
	"fiber_test/configs"
	"fiber_test/entities/handlers"
	"fiber_test/entities/router"
	"github.com/gofiber/fiber/v2"
)

func InitApiRoutes(app *fiber.App, conf configs.Config) {
	app.Get("/", handlers.Index)

	v1 := app.Group("v1")

	router.InitUserRoutes(v1, conf)
}
