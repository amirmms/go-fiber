package handlers

import (
	"fiber_test/configs"
	"fiber_test/entities/dtos"
	"fiber_test/entities/services"
	"fiber_test/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type UserHandler struct {
	UserService services.UserService
}

func InitUserHandler(conf configs.Config) UserHandler {
	return UserHandler{
		UserService: services.InitUserService(conf),
	}
}

func (h UserHandler) Index(ctx *fiber.Ctx) error {
	errService, m := h.UserService.FetchUsers()

	if errService != nil {
		return ctx.Status(400).JSON(utils.Response(false, "Bad Request", errService.Error()))
	}

	return ctx.Status(200).JSON(utils.Response(true, "user", m))
}

func (h UserHandler) Show(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	reqDto := dtos.GetUserDto{Id: id}

	err, msg := utils.Validator(ctx, &reqDto)

	if err != nil {
		return ctx.Status(422).JSON(utils.ValidationResponse(msg))
	}

	errService, m := h.UserService.GetUser(reqDto)

	if errService != nil {
		return ctx.Status(400).JSON(utils.Response(false, "Bad Request", errService.Error()))
	}

	return ctx.Status(200).JSON(utils.Response(true, "user", m))
}

func (h UserHandler) Store(ctx *fiber.Ctx) error {
	reqDto := dtos.CreateUserDto{}

	err, msg := utils.Validator(ctx, &reqDto)

	if err != nil {
		return ctx.Status(422).JSON(utils.ValidationResponse(msg))
	}

	errService, m := h.UserService.CreateUser(reqDto)

	if errService != nil {
		return ctx.Status(400).JSON(utils.Response(false, "Bad Request", errService.Error()))
	}

	return ctx.Status(200).JSON(utils.Response(true, "userCreated", m))
}

func (h UserHandler) Delete(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	reqDto := dtos.GetUserDto{Id: id}

	err, msg := utils.Validator(ctx, &reqDto)

	if err != nil {
		return ctx.Status(422).JSON(utils.ValidationResponse(msg))
	}

	errService, m := h.UserService.DeleteUser(reqDto)

	if errService != nil {
		return ctx.Status(400).JSON(utils.Response(false, "Bad Request", errService.Error()))
	}

	return ctx.Status(200).JSON(utils.Response(true, "userDeleted", m))
}

func (h UserHandler) Update(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	reqDto := dtos.UpdateUserDto{Id: id}

	err, msg := utils.Validator(ctx, &reqDto)

	if err != nil {
		return ctx.Status(422).JSON(utils.ValidationResponse(msg))
	}

	errService, m := h.UserService.UpdateUser(reqDto)

	if errService != nil {
		return ctx.Status(400).JSON(utils.Response(false, "Bad Request", errService.Error()))
	}

	return ctx.Status(200).JSON(utils.Response(true, "userUpdated", m))
}
