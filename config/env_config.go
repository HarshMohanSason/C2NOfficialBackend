package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

type EnvConfig struct {
	EnvType          string
	DbUrl            string
	JwtSecret        string
	ZohoMailPassword string
}

var Env *EnvConfig
var once sync.Once

func LoadEnv() {
	once.Do(func() {
		//Getting the env type (is set externally when runnning the main.go file )
		envType := os.Getenv("ENV_TYPE")

		var envFile string
		switch envType {
		case "PROD":
			envFile = ".env.production"
		case "DEV":
			envFile = ".env.development"
		default:
			log.Fatalf("Unknown or missing ENV_TYPE: %s", envType)
		}
		//Load the env file. Not handling errors since the file will always be present there
		_ = godotenv.Load(envFile)

		Env = &EnvConfig{
			EnvType:          os.Getenv("ENV_TYPE"),
			DbUrl:            os.Getenv("DATABASE_URL"),
			JwtSecret:        os.Getenv("JWT_SECRET"),
			ZohoMailPassword: os.Getenv("ZOHO_MAIL_PASSWORD"),
		}
	})
}
