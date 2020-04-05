package admin

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
	"mapoteca/database/models/authToken"
	"mapoteca/database/models/physicalPubKey"
	"mapoteca/logger"
	"mapoteca/models"
	"mapoteca/services/yubico"
	"time"
)

type successResponse struct {
	Status string `json:"success"`
}

func GenerateToken(context *fiber.Ctx) {
	var log = logger.New()
	log.Info("generating authentication token")

	var otp = context.Get("otp")
	fmt.Println(len(otp))
	fmt.Println(otp)
	if len(otp) < 13 {
		log.Info(fmt.Sprintf("Problem with otp length of %d", len(otp)))
		context.Status(401).JSON(models.HttpError{
			ErrorMessage: "Not authorized",
		})
		return
	}

	var keyIdentifier = otp[0:12]

	var _, keyErr = physicalPubKey.GetKeyByIdentifier(keyIdentifier)
	if keyErr != nil {
		var msg = fmt.Sprintf("key with id %s not found in database", keyIdentifier)
		log.Info(msg)
		context.Status(401).JSON(models.HttpError{
			ErrorMessage: "Not authorized",
		})
		return
	}

	var ok, err = yubico.OTPValidation(otp)
	if err != nil {
		log.Error(err)
		context.Status(401).JSON(models.HttpError{
			ErrorMessage: "Not authorized",
		})
		return
	}
	if ok {
		log.Info("authentication was successful")
		var authToken = authToken.InsertToken(keyIdentifier)

		var cookie = new(fiber.Cookie)
		cookie.Name = "robson"
		cookie.Value = uuid.UUID.String(authToken)
		cookie.Expires = time.Now().Add(24 * time.Hour)

		context.Cookie(cookie)
		context.JSON(successResponse{
			Status: "success",
		})
	}
}
