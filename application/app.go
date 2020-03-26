package application

import (
	"github.com/gofiber/fiber"
)

func Run() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	app.Listen(3001)
}
