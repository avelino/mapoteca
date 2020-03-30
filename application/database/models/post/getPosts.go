package post

import (
	"github.com/jinzhu/gorm"
	"mapoteca/application/database"
	"mapoteca/application/logger"
	"mapoteca/application/models"
)

func GetPosts() ([]models.Post, error) {
	var log = logger.New()
	log.Info("getting posts from database")
	var posts []models.Post
	var err = database.DB.Transaction(func(tx *gorm.DB) error {
		var d = tx.Select("id, title, subtitle, slug, category, image_path, created_at").Find(&posts)

		return d.Error
	})

	if err == nil {
		return posts, nil
	}

	return nil, err
}
