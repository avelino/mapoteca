package admin

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
	"mapoteca/database/models/authToken"
	"mapoteca/logger"
	"mapoteca/models"
	"time"
)

func VerifyAuth(context *fiber.Ctx) {
	var log = logger.New()
	log.Info("verifying authentication")
	var robson = context.Cookies("robson")
	var uuidCookie, uuidCookieErr = uuid.Parse(robson)
	fmt.Println(">>>>>>>>>>>>>> %d", robson)
	fmt.Println(">>>>>>>>>>>>>> %s", robson)
	fmt.Println(">>>>>>>>>>>>>> %d", context.Get("otp"))
	fmt.Println(">>>>>>>>>>>>>> %s", context.Get("otp"))
	if uuidCookieErr != nil {
		log.Error(fmt.Sprintf("Problem parsing token. Error: %d", uuidCookieErr))
		context.Status(401).JSON(models.HttpError{
			ErrorMessage: "Not authorized",
		})
		return
	}

	var token, tokenErr = authToken.GetToken(uuidCookie)
	if tokenErr != nil {
		log.Error(fmt.Sprintf("Problem getting token at database. Error: %d", tokenErr))
		context.Status(401).JSON(models.HttpError{
			ErrorMessage: "Not authorized",
		})
		return
	}

	var tokenValidity = token.CreatedAt.Local().Add(time.Hour * 24)

	log.Info("validating token age")
	if tokenValidity.Before(time.Now()) {
		log.Error("Token is expired")
		context.Status(401).JSON(models.HttpError{
			ErrorMessage: "Not authorized",
		})
		return
	}

	context.JSON(models.HttpSuccess{
		Greetings: "Authorized",
	})
}
