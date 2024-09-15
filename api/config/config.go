package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	utils "github.com/ncephamz/efishery-be-test/api/pkg"
)

type (
	Database struct {
		DSN          string
		MaxIdleConns int
		MaxOpenConns int
	}

	Config struct {
		Database  Database
		JwtSecret string
		AllowCors []string
		Port      string
	}
)

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}

	database := Database{
		DSN:          os.Getenv("POSTGRES_DSN"),
		MaxIdleConns: utils.StringToInt(os.Getenv("POSTGRES_MAX_IDLE_CONNS")),
		MaxOpenConns: utils.StringToInt(os.Getenv("POSTGRES_OPEN_CONNS")),
	}

	allowCors := os.Getenv("ALLOW_CORS")

	return Config{
		Database:  database,
		JwtSecret: os.Getenv("SECRET_JWT"),
		Port:      os.Getenv("PORT"),
		AllowCors: strings.Split(allowCors, ","),
	}
}
