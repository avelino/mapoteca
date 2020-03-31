package admin

import (
	"encoding/json"
	"fmt"
	"github.com/GeertJohan/yubigo"
	"github.com/gofiber/fiber"
	"mapoteca/config"
	"mapoteca/logger"
	"mapoteca/models"
)

type authForm struct {
	Username string
	OTP      string
}

func GenerateToken(context *fiber.Ctx) {
	var log = logger.New()
	log.Info("generating authentication token")
	var authData authForm
	var byteBody = []byte(context.Body())
	json.Unmarshal(byteBody, &authData)

	var keyIdentifier = authData.OTP[0:12]
	fmt.Println(len(keyIdentifier))

	log.Info("creating a new yubico auth instance")
	var yubicoAuth, _ = yubigo.NewYubiAuth(config.YubicoConfig.ClientId, config.YubicoConfig.ApiKey)

	log.Info("verificating OTP with yubico instance")
	var response, ok, err = yubicoAuth.Verify(authData.OTP)
	if err != nil {
		log.Error(err)
		context.Status(401).JSON(models.HttpError{
			ErrorMessage: "Not authorized",
		})
		return
	}
	if ok {
		log.Info("success")
		context.Send(response)
	}
}
