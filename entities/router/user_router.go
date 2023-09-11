package router

import (
	"fiber_test/configs"
	"fiber_test/entities/handlers"
	"github.com/gofiber/fiber/v2"
)

func InitUserRoutes(parentRoute fiber.Router, conf configs.Config) {
	userGroup := parentRoute.Group("users")

	userHandler := handlers.InitUserHandler(conf)

	userGroup.Get("/", userHandler.Index)
	userGroup.Get("/:id", userHandler.Show)
	userGroup.Delete("/:id", userHandler.Delete)
	userGroup.Patch("/:id", userHandler.Update)
	userGroup.Post("/", userHandler.Store)
}
