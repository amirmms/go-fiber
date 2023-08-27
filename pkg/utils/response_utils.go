package utils

import "github.com/gofiber/fiber/v2"

func Response(isSuccess bool, message string, data any) fiber.Map {
	status := "nok"
	if isSuccess {
		status = "ok"
	}

	return fiber.Map{
		"status":  status,
		"message": message,
		"data":    data,
	}
}
