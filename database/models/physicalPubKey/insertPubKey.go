package physicalPubKey

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"mapoteca/database"
	"mapoteca/database/models"
	"mapoteca/logger"
	"time"
)

func InsertPubKey(pubId string) error {
	var log = logger.New()
	log.Info("inserting new pub key into database")
	return database.DB.Transaction(func(tx *gorm.DB) error {
		var d = tx.Create(&models.PhysicalPubKey{
			ID:            uuid.New(),
			CreatedAt:     time.Now(),
			KeyIdentifier: pubId,
		})

		return d.Error
	})
}
