package authToken

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"mapoteca/database"
	"mapoteca/database/models"
	"mapoteca/logger"
)

func GetToken(token uuid.UUID) (models.AuthToken, error) {
	var log = logger.New()
	log.Info("consulting token into database")
	var authToken models.AuthToken
	var err = database.DB.Transaction(func(tx *gorm.DB) error {
		var d = tx.Where("id = ?", token).First(&authToken)

		return d.Error
	})

	return authToken, err
}
