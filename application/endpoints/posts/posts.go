package posts

import (
	"mapoteca/application/models"
	"time"
)

func GetPosts() []models.Post {
	var list []models.Post
	var post = models.Post{
		Id:        "i1ub3irhu9",
		Title:     "Um roteador portátil para chamar de seu",
		Slug:      "um-roteador-portatil-para-chamar-de-seu",
		Subtitle:  "Como usar seu Raspberry para garantir sua privacidade no trabalho",
		ImagePath: "",
		Category:  "software",
		CreatedAt: time.Now(),
	}

	var posts = append(list, post)

	return posts
}
