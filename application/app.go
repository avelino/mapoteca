package application

import (
	"encoding/json"
	"github.com/gofiber/fiber"
	"mapoteca/application/endpoints/posts"
)

func Run() {
	app := fiber.New()

	app.Get("/posts", func(c *fiber.Ctx) {
		var posts = posts.GetPosts()
		var response, _ = json.Marshal(posts)

		c.Send(response)
	})

	app.Get("/ping", func(c *fiber.Ctx) {
		c.Send("pong")
	})

	app.Listen(3001)
}
