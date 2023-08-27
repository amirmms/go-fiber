package handlers

import (
	"fiber_test/configs"
	"fiber_test/entities/services"
	"fiber_test/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService services.UserService
}

func InitUserHandler(conf configs.Config) UserHandler {
	return UserHandler{
		UserService: services.InitUserService(conf),
	}
}

func (h UserHandler) GetUserList(c *fiber.Ctx) error {
	return c.Status(200).JSON(utils.Response(true, "test", "test"))
}
