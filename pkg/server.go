package pkg

import (
	"fiber_test/configs"
	"fiber_test/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func InitFiber(conf configs.Config) *fiber.App {
	return fiber.New(fiber.Config{
		AppName:   conf.AppName,
		BodyLimit: conf.BodyLimit,
	})
}

func ServeFiber(app *fiber.App, conf configs.Config) {
	if conf.AppEnv == configs.PROD {
		utils.ServeFiberInProductionMode(app, conf)
	} else {
		utils.ServeFiberInDevelopMode(app, conf)
	}
}
