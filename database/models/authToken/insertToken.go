package authToken

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"mapoteca/database"
	databaseModels "mapoteca/database/models"
	"mapoteca/logger"
	"time"
)

func InsertToken(keyPubId string) uuid.UUID {
	var log = logger.New()
	log.Info("inserting token to database")
	var token uuid.UUID
	database.DB.Transaction(func(tx *gorm.DB) error {
		var id = uuid.New()
		var d = tx.Create(&databaseModels.AuthToken{
			ID:              id,
			KeyIdentifierId: keyPubId,
			CreatedAt:       time.Now(),
		})

		if d.Error != nil {
			var msg = fmt.Sprintf("problem at transaction with error: %d", d.Error)
			log.Error(msg)
		}

		fmt.Sprintln(d)
		token = id
		return d.Error
	})
	return token
}
