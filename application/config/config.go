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

var DatabaseConfig databaseConfig
var ServerConfig serverConfig

func Init() {
	var err = godotenv.Load()

	if err != nil {
		panic("Env vars not available")
	}

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
}
