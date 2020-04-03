package post

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber"
	postModel "mapoteca/database/models/post"
	"mapoteca/logger"
	"mapoteca/models"
)

func UpdatePost(context *fiber.Ctx) {
	var log = logger.New()
	log.Info("updating post")
	var post models.Post
	var bytePostData = []byte(context.Body())
	json.Unmarshal(bytePostData, &post)

	var success, fields = ValidateFields(&HttpRequestPost{
		Title:    post.Title,
		Subtitle: post.Subtitle,
		Content:  post.Content,
		Category: post.Category,
	})
	if !success {
		var msg = models.HttpError{
			ErrorMessage:    "One or more fields were not acceptable",
			FormErrorFields: fields,
		}
		log.Info(msg.ErrorMessage)
		context.Status(406).JSON(msg)
		return
	}

	var err = postModel.UpdatePost(post)

	if err != nil {
		context.Status(500).JSON(models.HttpError{
			ErrorMessage: fmt.Sprintf("Problem updating post. Error: %s", err),
		})
		return
	}

	context.JSON(models.HttpSuccess{
		Greetings: "All good",
	})
	return
}
