package application

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	// importing to be run at Global
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"mapoteca/application/config"
	"mapoteca/application/endpoints/posts"
	"mapoteca/application/logger"
)

func Run() {
	var log = logger.New()
	log.Info("initializing the application")
	config.Init()

	var app = fiber.New()
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
		log.Error("problem initializing database")
		fmt.Println(dbError, config.DatabaseConfig)
		fmt.Println(db, config.DatabaseConfig)
	}

	defer db.Close()

	app.Get("/posts", func(c *fiber.Ctx) {
		var posts = posts.GetPosts()
		var response, _ = json.Marshal(posts)

		c.Send(response)
	})

	app.Get("/ping", func(c *fiber.Ctx) {
		c.Send("pong")
	})

	app.Listen(3001)
}
