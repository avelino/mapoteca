package post

import (
	"github.com/gofiber/fiber"
	"mapoteca/application/database/models/post"
	"mapoteca/application/logger"
)

func GetPosts(context *fiber.Ctx) {
	var log = logger.New()
	log.Info("getting all posts")
	var posts, _ = post.GetPosts()

	context.JSON(posts)
}
