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

	var authorization = context.Get("Authorization")
	var uuidToken, uuidTokenErr = uuid.Parse(authorization)
	if uuidTokenErr != nil {
		log.Error(fmt.Sprintf("Problem parsing token. Error: %d", uuidTokenErr))
		context.Status(401).JSON(models.HttpError{
			ErrorMessage: "Not authorized",
		})
		return
	}

	var token, tokenErr = authToken.GetToken(uuidToken)
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
