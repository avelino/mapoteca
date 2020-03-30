package application

import (
	"encoding/json"
	"github.com/gofiber/fiber"
	"mapoteca/application/config"
	"mapoteca/application/database"
	"mapoteca/application/endpoints/post"
	"mapoteca/application/logger"
)

func Run() {
	var log = logger.New()
	log.Info("initializing the application")
	config.Init()

	var app = fiber.New()

	var db = database.Connect()
	defer db.Close()

	app.Get("/posts", func(c *fiber.Ctx) {
		log.Info("request made at /posts")
		var posts = post.GetPosts()
		var response, _ = json.Marshal(posts)

		c.Send(response)
	})

	app.Get("/ping", func(c *fiber.Ctx) {
		c.Send("pong")
	})

	app.Post("/post/new", post.NewPost)

	log.Info("application is running at port " + config.ServerConfig.Port)
	app.Listen(config.ServerConfig.Port)
}
