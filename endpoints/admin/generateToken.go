package admin

import (
	"encoding/json"
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

type authForm struct {
	OTP string
}

type successResponse struct {
	Status string `json:"success"`
}

func GenerateToken(context *fiber.Ctx) {
	var log = logger.New()
	log.Info("generating authentication token")
	var authData authForm
	var byteBody = []byte(context.Body())
	json.Unmarshal(byteBody, &authData)

	var keyIdentifier = authData.OTP[0:12]

	var _, keyErr = physicalPubKey.GetKeyByIdentifier(keyIdentifier)
	if keyErr != nil {
		var msg = fmt.Sprintf("key with id %s not found in database", keyIdentifier)
		log.Info(msg)
		context.Status(401).JSON(models.HttpError{
			ErrorMessage: "Not authorized",
		})
		return
	}

	var ok, err = yubico.OTPValidation(authData.OTP)
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
