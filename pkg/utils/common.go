package utils

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type Flags struct {
	Active    bool
	IsDeleted string
}

func GoDotEnvVariable(key string) string {

	if key == "" {
		return KeyNotFound.Error()
	}

	err := godotenv.Load(".env")

	if err != nil {
		return EnvNotLoaded.Error()
	}

	return os.Getenv(key)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func RemoveHyphens(ui string) string {
	return strings.Replace(ui, "-", "", 1)
}
