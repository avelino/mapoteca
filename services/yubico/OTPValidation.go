package yubico

import (
	"github.com/GeertJohan/yubigo"
	"mapoteca/config"
	"mapoteca/logger"
)

func OTPValidation(otp string) (bool, error) {
	var log = logger.New()
	log.Info("creating a new yubico auth instance")
	var yubicoAuth, _ = yubigo.NewYubiAuth(config.YubicoConfig.ClientId, config.YubicoConfig.ApiKey)
	log.Info("verificating OTP with yubico instance")
	var _, ok, err = yubicoAuth.Verify(otp)

	return ok, err
}
