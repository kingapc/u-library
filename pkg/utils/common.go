package utils

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/twinj/uuid"
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

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func RemoveHyphens(ui string) string {
	return strings.Replace(ui, "-", "", -1)
}

func ValidateParams(c *gin.Context, param []string) (string, string) {

	var p string
	var v string
	p = ""
	v = ""

	for i, s := range param {
		b := c.Request.URL.Query()[param[i]]
		if !(len(b) < 1) {
			p = s
			v = b[0]
		}
	}

	return p, v
}

func IsValidUUID(u string) (string, bool) {

	_, err := uuid.Parse(u)

	return u, (err == nil)
}
