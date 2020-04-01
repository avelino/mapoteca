package post

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"mapoteca/database"
	databaseModels "mapoteca/database/models"
	"mapoteca/logger"
	"mapoteca/models"
)

func InsertPost(p models.Post) (uuid.UUID, error) {
	var log = logger.New()
	log.Info("inserting post to database")
	var postId uuid.UUID
	var err = database.DB.Transaction(func(tx *gorm.DB) error {
		var id = uuid.New()
		var d = tx.Create(&databaseModels.Post{
			ID:        id,
			CreatedAt: p.CreatedAt,
			Title:     p.Title,
			Subtitle:  p.Subtitle,
			Slug:      p.Slug,
			ImagePath: p.ImagePath,
			Category:  p.Category,
			Content:   p.Content,
		})

		if d.Error != nil {
			var msg = fmt.Sprintf("problem at transaction with error: %d", d.Error)
			log.Error(msg)
		} else {
			postId = id
		}
		return d.Error
	})

	return postId, err
}
