package post

import (
	"github.com/jinzhu/gorm"
	"mapoteca/database"
	"mapoteca/logger"
	"mapoteca/models"
)

func GetPostBySlug(slug string) (models.Post, error) {
	var log = logger.New()
	log.Info("querying database")
	var post models.Post
	var err = database.DB.Transaction(func(tx *gorm.DB) error {
		var d = tx.Where("slug = ?", slug).First(&post)

		return d.Error
	})

	return post, err
}
