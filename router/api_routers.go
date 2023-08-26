package router

import (
	"fiber_test/handlers/home"
	"github.com/gofiber/fiber/v2"
)

func InitApiRoutes(app *fiber.App) {
	app.Get("/", home.Index)

	v1 := app.Group("v1")

	v1.Get("/test", func(ctx *fiber.Ctx) error {
		return ctx.SendString("v1/test")
	})
}
