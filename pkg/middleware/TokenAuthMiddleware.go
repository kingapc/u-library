package securitty

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rpinedafocus/u-library/pkg/model"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t, err := TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		claims, ok := t.Claims.(jwt.MapClaims)
		if ok && t.Valid {
			uuidValue, ok := claims["access_uuid"].(string)
			if !ok {
				c.JSON(http.StatusUnprocessableEntity, err)
				c.Abort()
				return
			}

			err := ValidateSession(uuidValue)
			if err != nil {
				c.JSON(http.StatusUnauthorized, "Invalid Session Test")
				c.Abort()
				return
			}
		}
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
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
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

/************ LOGOUT	*******************/
func ExtractTokenMetadata(r *http.Request) (*model.AccessDetails, error) {

	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}

		fmt.Print(claims)

		userId := claims["user_id"].(string)
		//refresh := claims["refresh_uuid"].(string)

		return &model.AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
			//RefreshUuid: refresh,
		}, nil
	}

	return nil, err
}

func DeleteTokens(access string) error {

	//delete access token
	deletedAt, err := Client.Del(access).Result()
	if err != nil {
		return err
	}

	//delete refresh token
	// deletedRt, err := Client.Del(refresh).Result()
	// if err != nil {
	// 	return err
	// }

	//When the record is deleted, the return value is 1
	if deletedAt != 1 {
		return errors.New("something went wrong")
	}

	return nil
}
