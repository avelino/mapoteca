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

func InsertPost(p models.Post) error {
	var log = logger.New()
	log.Info("inserting post to database")
	return database.DB.Transaction(func(tx *gorm.DB) error {
		var d = tx.Create(&databaseModels.Post{
			ID:        uuid.New(),
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
		}
		return d.Error
	})
}
