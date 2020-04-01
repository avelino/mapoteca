package post

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
	slugify "github.com/gosimple/slug"
	authTokenModel "mapoteca/database/models/authToken"
	postModel "mapoteca/database/models/post"
	"mapoteca/logger"
	"mapoteca/models"
	"reflect"
	"time"
)

type newPost struct {
	Title     string
	Subtitle  string
	Slug      string
	Content   string
	Category  string
	ImagePath string
}

func getFieldString(p *newPost, field string) string {
	r := reflect.ValueOf(p)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}

func validateFields(p *newPost) (bool, []string) {
	var list = [4]string{"Title", "Subtitle", "Content", "Category"}
	var emptyFields []string

	for i := 0; i < len(list); i++ {
		var current = list[i]
		var v = getFieldString(p, current)
		if v == "" {
			emptyFields = append(emptyFields, current)
		}
	}

	return len(emptyFields) == 0, emptyFields
}

func slugCreator(slug string, title string) string {
	if slug == "" || len(slug) < 5 {
		return slugify.Make(title)
	}

	return slugify.Make(slug)
}

func NewPost(context *fiber.Ctx) {
	var log = logger.New()
	log.Info("new post endpoint called")
	log.Info("authenticating request")
	var authTokenCookie = context.Cookies("robson")
	var authTokenUUID, authTokenErr = uuid.Parse(authTokenCookie)
	var _, tokenErr = authTokenModel.GetToken(authTokenUUID)
	if authTokenErr != nil || tokenErr != nil {
		var msg = fmt.Sprintf("authentication failed: %d", tokenErr)
		log.Error(msg)
		var response = models.HttpError{
			ErrorMessage: "Not authorized",
		}
		context.Status(401).JSON(response)
		return
	}

	log.Info("authentication successful")

	var data newPost
	var body = []byte(context.Body())
	json.Unmarshal(body, &data)

	var success, fields = validateFields(&data)

	if !success {
		var msg = models.HttpError{
			ErrorMessage:    "One or more fields were not acceptable",
			FormErrorFields: fields,
		}
		log.Info(msg.ErrorMessage)
		context.Status(406).JSON(msg)
		return
	}

	var postId, insertPostErr = postModel.InsertPost(models.Post{
		Title:     data.Title,
		Subtitle:  data.Subtitle,
		Slug:      slugCreator(data.Slug, data.Title),
		ImagePath: data.ImagePath,
		Category:  data.Category,
		CreatedAt: time.Now(),
		Content:   data.Content,
	})

	if insertPostErr != nil {
		var msg = models.HttpError{
			ErrorMessage:    fmt.Sprintln(insertPostErr),
			FormErrorFields: nil,
		}
		log.Error(msg)
		context.Status(500).JSON(msg)
		return
	}

	var generatedIds []uuid.UUID

	context.JSON(models.HttpSuccess{
		Greetings:    "All good",
		GeneratedIds: append(generatedIds, postId),
	})
}