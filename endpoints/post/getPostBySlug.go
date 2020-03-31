package post

import (
	"github.com/gofiber/fiber"
	"mapoteca/database/models/post"
	"mapoteca/logger"
	"mapoteca/models"
)

func GetPostBySlug(context *fiber.Ctx) {
	var log = logger.New()
	log.Info("getting post")

	var slug = context.Params("slug")
	log.Info(slug)
	var post, err = post.GetPostBySlug(slug)

	if err != nil {
		context.Status(404).JSON(models.HttpError{
			ErrorMessage: "Not found",
		})
		return
	}

	context.JSON(post)
}
