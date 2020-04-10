package config

import (
	"github.com/joho/godotenv"
	"os"
)

type databaseConfig struct {
	Name     string
	Host     string
	Port     string
	User     string
	Password string
	SslMode  string
}

type serverConfig struct {
	Port string
}

type yubicoConfig struct {
	ClientId string
	ApiKey   string
}

type adminConfig struct {
	MasterPublicKey string
}

var DatabaseConfig databaseConfig
var ServerConfig serverConfig
var YubicoConfig yubicoConfig
var AdminConfig adminConfig
var ClientUrl string

func Init() {
	var _ = godotenv.Load()

	DatabaseConfig = databaseConfig{
		Name:     os.Getenv("POSTGRESQL_DATABASE_NAME"),
		Host:     os.Getenv("POSTGRESQL_DATABASE_HOST"),
		Port:     os.Getenv("POSTGRESQL_DATABASE_PORT"),
		User:     os.Getenv("POSTGRESQL_DATABASE_USER"),
		Password: os.Getenv("POSTGRESQL_DATABASE_PASSWORD"),
		SslMode:  os.Getenv("POSTGRESQL_DATABASE_SSL_MODE"),
	}

	ServerConfig = serverConfig{
		Port: os.Getenv("SERVER_PORT"),
	}

	YubicoConfig = yubicoConfig{
		ClientId: os.Getenv("YUBICO_CLIENT_ID"),
		ApiKey:   os.Getenv("YUBICO_API_KEY"),
	}

	AdminConfig = adminConfig{
		MasterPublicKey: os.Getenv("ADMIN_MASTER_PUBLIC_KEY"),
	}

	ClientUrl = os.Getenv("CLIENT_URL")
}
