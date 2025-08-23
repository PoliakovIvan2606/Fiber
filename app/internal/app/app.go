package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewApp() *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:       true,  // использовать несколько процессов
		CaseSensitive: true,  // чувствительность к регистру
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(500).SendString(err.Error())
		},
	})

	app.Use(logger.New())
	
	return app
}

// Middleware
// app.Use(func(c *fiber.Ctx) error {
// 	println("Request path:", c.Path())
// 	return c.Next()
// })