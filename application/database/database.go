package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	// importing to be run at Global
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"mapoteca/application/config"
	"mapoteca/application/database/models"
	"mapoteca/application/logger"
	"os"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	var log = logger.New()
	log.Info("connecting to database")

	var dbConfig = fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.DatabaseConfig.Host,
		config.DatabaseConfig.Port,
		config.DatabaseConfig.User,
		config.DatabaseConfig.Name,
		config.DatabaseConfig.Password,
		config.DatabaseConfig.SslMode,
	)

	var db, dbError = gorm.Open("postgres", dbConfig)

	if dbError != nil {
		var message = fmt.Sprintf("problem initializing database. Error: %d", dbError)
		log.Error(message)
		os.Exit(1)
	}

	autoMigrate(db)
	DB = db
	return db
}

func autoMigrate(db *gorm.DB) {
	var log = logger.New()
	log.Info("running AutoMigrate")

	var migration = db.AutoMigrate(&models.Post{})

	if migration.Error != nil {
		var msg = fmt.Sprintf("problem running migration: %d", migration.Error)
		log.Error(msg)
		os.Exit(1)
	}
}
