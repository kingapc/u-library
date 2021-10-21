package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func GoDotEnvVariable(key string) string {

	if key == "" {
		return KeyNotFound.Error()
	}

	err := godotenv.Load("../../.env")

	if err != nil {
		return EnvNotLoaded.Error()
	}

	return os.Getenv(key)
}
