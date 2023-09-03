package handlers

import (
	"fiber_test/pkg/logger"
	"fiber_test/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	logger.Infof("simple zap logger example")

	return c.Status(200).JSON(utils.Response(false, "user", "test"))
}
