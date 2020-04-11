package server

import (
	"fmt"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"mapoteca/config"
	"mapoteca/database"
	"mapoteca/database/models/physicalPubKey"
	"mapoteca/endpoints/admin"
	"mapoteca/endpoints/post"
	"mapoteca/logger"
)

func Run() {
	var log = logger.New()
	log.Info("initializing the application")
	config.Init()

	var app = fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.ClientUrl},
		AllowCredentials: true,
	}))

	fmt.Printf("CLIENT >>>>>>>>>>>> %s \n", config.ClientUrl)

	var db = database.Connect()
	defer db.Close()
	defineAdminKey(db)

	app.Get("/ping", func(c *fiber.Ctx) {
		c.Send("pong")
	})

	app.Get("/post", post.GetPosts)
	app.Get("/post/:slug", post.GetPostBySlug)

	app.Get("/admin/auth", admin.GenerateToken)
	app.Get("/admin/auth/verify", admin.VerifyAuth)

	app.Post("/post/new", func(context *fiber.Ctx) {
		Authenticate(context, post.NewPost)
	})
	app.Put("/post/update/:id", func(context *fiber.Ctx) {
		Authenticate(context, post.UpdatePost)
	})

	log.Info("application is running at port " + config.ServerConfig.Port)
	app.Listen(config.ServerConfig.Port)
}

func defineAdminKey(db *gorm.DB) {
	var log = logger.New()
	log.Info("checking if admin public key is saved")
	var _, err = physicalPubKey.GetKeyByIdentifier(config.AdminConfig.MasterPublicKey)
	if err != nil {
		log.Info("key is not in database. Inserting it")
		physicalPubKey.InsertPubKey(config.AdminConfig.MasterPublicKey)
	}
}
