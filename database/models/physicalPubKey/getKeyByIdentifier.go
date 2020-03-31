package physicalPubKey

import (
	"github.com/jinzhu/gorm"
	"mapoteca/database"
	"mapoteca/database/models"
	"mapoteca/logger"
)

func GetKeyByIdentifier(pubId string) (models.PhysicalPubKey, error) {
	var log = logger.New()
	var key models.PhysicalPubKey
	var err = database.DB.Transaction(func(tx *gorm.DB) error {
		var d = tx.Where("key_identifier = ?", pubId).First(&key)

		return d.Error
	})

	if err != nil {
		log.Error(err)
	}

	return key, err
}
