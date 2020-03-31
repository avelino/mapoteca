package server

import (
	"github.com/gofiber/fiber"
	"mapoteca/config"
	"mapoteca/database"
	"mapoteca/endpoints/post"
	"mapoteca/logger"
)

func Run() {
	var log = logger.New()
	log.Info("initializing the application")
	config.Init()

	var app = fiber.New()

	var db = database.Connect()
	defer db.Close()

	app.Get("/ping", func(c *fiber.Ctx) {
		c.Send("pong")
	})

	app.Get("/post", post.GetPosts)
	app.Get("/post/:slug", post.GetPostBySlug)

	app.Post("/post/new", post.NewPost)

	log.Info("application is running at port " + config.ServerConfig.Port)
	app.Listen(config.ServerConfig.Port)
}
