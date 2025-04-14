package config 

import (
	"github.com/joho/godotenv"
	"sync"
	"os"
	"log"
)

type EnvConfig struct{
	ENV_TYPE string
	DB_URL string
	JWT_SECRET string
	ZOHO_MAIL_PASSWORD string 
}

var Env *EnvConfig
var once sync.Once

func LoadEnv(){
	once.Do(func() {
		//Getting the env type (is set externally when runnning the main.go file )
		envType := os.Getenv("ENV_TYPE")

		var envFile string 
		switch envType{
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
			ENV_TYPE:        os.Getenv("ENV_TYPE"),
			DB_URL:          os.Getenv("DATABASE_URL"),
			JWT_SECRET:      os.Getenv("JWT_SECRET"),
			ZOHO_MAIL_PASSWORD: os.Getenv("ZOHO_MAIL_PASSWORD"),
		}
	})
}