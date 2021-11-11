package securitty

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rpinedafocus/u-library/internal/app/u-library/utils"
)

func ValidateTokenSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		t, err := TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusUnauthorized), false, err.Error(), false)})
			c.Abort()
			return
		}

		claims, ok := t.Claims.(jwt.MapClaims)
		if ok && t.Valid {
			uuidValue, ok := claims["access_uuid"].(string)
			if !ok {
				c.JSON(http.StatusUnprocessableEntity, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusUnprocessableEntity), false, err.Error(), false)})
				c.Abort()
				return
			}

			err := ValidateSession(uuidValue)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusUnauthorized), false, err.Error(), false)})
				c.Abort()
				return
			}
		}

		//Gettin the information
		e, err := ExtractTokenMetadata(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusUnauthorized), false, err.Error(), false)})
			c.Abort()
			return
		}

		myData := map[string]interface{}{
			"uuid":   e.AccessUuid,
			"myId":   e.MyId,
			"userId": e.UserId,
			"roleId": e.RoleId,
		}

		c.Set("MyInformation", myData)
		c.Next()
	}
}

func TokenValid(r *http.Request) (*jwt.Token, error) {
	token, err := VerifyToken(r)

	if err != nil {
		return nil, err
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if ok && !token.Valid {
		return nil, err
	}

	return token, nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(utils.ErrSingMethod.Error()+" %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func ExtractToken(r *http.Request) string {

	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func ValidateSession(uuid string) error {

	_, err := Client.Get(uuid).Result()

	if err != nil {
		return err
	}

	return nil
}
