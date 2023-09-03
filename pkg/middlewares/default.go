package middlewares

import (
	"fiber_test/configs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func InitDefaultMiddlewares(app *fiber.App, conf configs.Config) {
	//https://docs.gofiber.io/api/middleware/cors
	app.Use(cors.New())

	//https://docs.gofiber.io/api/middleware/encryptcookie
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: conf.EncryptCookieKey,
	}))

	//https://docs.gofiber.io/api/middleware/requestid
	app.Use(requestid.New())

	//https://docs.gofiber.io/api/middleware/logger
	app.Use(logger.New())

	//https://docs.gofiber.io/api/middleware/helmet
	app.Use(helmet.New())

	//https://docs.gofiber.io/api/middleware/idempotency
	app.Use(idempotency.New())

	//https://docs.gofiber.io/api/middleware/limiter
	app.Use(limiter.New())

	//https://docs.gofiber.io/api/middleware/recover
	app.Use(recover.New())
}
