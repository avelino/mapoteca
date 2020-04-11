package server

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
	authTokenModel "mapoteca/database/models/authToken"
	"mapoteca/logger"
	"mapoteca/models"
	"mapoteca/services/yubico"
)

func Authenticate(context *fiber.Ctx, endpointCallback func(*fiber.Ctx)) {
	var log = logger.New()
	log.Info("authenticating request")
	var otp = context.Get("otp")
	var authorization = context.Get("Authorization")
	var authTokenUUID, authTokenErr = uuid.Parse(authorization)
	var _, tokenErr = authTokenModel.GetToken(authTokenUUID)

	if authTokenErr != nil || tokenErr != nil {
		log.Error(fmt.Sprintf("authentication failed: %d", tokenErr))
		context.Status(401).JSON(models.HttpError{
			ErrorMessage: "Not authorized",
		})
		return
	}

	var ok, err = yubico.OTPValidation(otp)
	if ok {
		log.Info("all good")
		endpointCallback(context)
		return
	}
	log.Error(fmt.Sprintf("OTP not valid. Error: %s", err))
	context.Status(401).JSON(models.HttpError{
		ErrorMessage: "Not authorized",
	})
	return
}
