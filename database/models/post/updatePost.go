package post

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"mapoteca/database"
	"mapoteca/logger"
	"mapoteca/models"
)

func UpdatePost(payload models.Post) error {
	var log = logger.New()
	log.Info(fmt.Sprintf("Updating post of id %d", payload.ID))
	var post models.Post
	var err = database.DB.Transaction(func(tx *gorm.DB) error {
		var d = tx.Model(&post).Where("id = ?", payload.ID).Update(models.Post{
			ID:        payload.ID,
			CreatedAt: payload.CreatedAt,
			Title:     payload.Title,
			Subtitle:  payload.Subtitle,
			ImagePath: payload.ImagePath,
			Category:  payload.Category,
			Content:   payload.Content,
		})

		return d.Error
	})

	return err
}
