package securitty

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rpinedafocus/u-library/internal/app/u-library/dal"
	"github.com/rpinedafocus/u-library/internal/app/u-library/model"
	"github.com/rpinedafocus/u-library/internal/app/u-library/utils"
)

func PrepareLogin(user string, password string) (*model.UserEntity, error) {

	e, err := dal.Login(user, password)
	if err != nil {
		return nil, err
	}

	return e, nil
}

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

		userId := claims["user_id"].(string)
		myId := claims["id"].(string)
		roleId := claims["role_id"].(string)

		return &model.AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
			MyId:       myId,
			RoleId:     roleId,
		}, nil
	}

	return nil, err
}

func ValidateAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Get the information
		a, ok := c.Get("MyInformation")
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusNotFound), false, "", false)})
			c.Abort()
			return
		}

		//Validate the map
		b, ok := a.(map[string]interface{})
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusBadRequest), false, "", false)})
			c.Abort()
			return
		}

		//Assingning the values
		e := &model.AccessDetails{}
		e.AccessUuid = b["uuid"].(string)
		e.MyId = b["myId"].(string)
		e.UserId = b["userId"].(string)
		e.RoleId = b["roleId"].(string)

		//Is necessary to cast to int
		tempId, err := strconv.Atoi(e.RoleId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusBadRequest), true, utils.EmtpyModel.Error(), false)})
			c.Abort()
			return
		}

		//Validate Access()
		rest := dal.FetchAccessRolById(tempId, c.FullPath())
		if !rest {
			c.JSON(http.StatusUnauthorized, gin.H{"operation": utils.ErrorX(http.StatusText(http.StatusUnauthorized), false, "", false)})
			c.Abort()
			return
		}

		myData := map[string]interface{}{
			"uuid":   e.AccessUuid,
			"myId":   e.MyId,
			"userId": e.UserId,
			"roleId": e.RoleId,
		}

		c.Set("myData", myData)
		c.Next()
	}
}

func GetSessionDetail(c *gin.Context) *model.AccessDetails {

	//Get the information
	data, ok := c.Get("myData")
	if !ok {
		return nil
	}

	//Validate the map
	mapData, ok := data.(map[string]interface{})
	if !ok {
		return nil
	}

	//Assingning the values
	e := &model.AccessDetails{}
	e.AccessUuid = mapData["uuid"].(string)
	e.MyId = mapData["myId"].(string)
	e.UserId = mapData["userId"].(string)
	e.RoleId = mapData["roleId"].(string)

	return e
}
