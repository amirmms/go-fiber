package handlers

import (
	"fiber_test/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	return c.Status(200).JSON(utils.Response(false, "user", "test"))
}
