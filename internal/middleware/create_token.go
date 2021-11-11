package securitty

import (
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/rpinedafocus/u-library/internal/model"
	"github.com/rpinedafocus/u-library/internal/utils"
	"github.com/twinj/uuid"

	"os"
	"time"
)

func CreateToken(id string, user string, role int) (*model.TokenDetails, error) {

	td := &model.TokenDetails{}

	//Add expire time and random uuid
	td.AtExpires = time.Now().Add(time.Minute * 30).Unix()
	td.AccessUuid = uuid.NewV4().String()

	//Adding the info
	var err error
	os.Setenv("ACCESS_SECRET", utils.GoDotEnvVariable("ACCESS_SECRET"))
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["id"] = id
	atClaims["user_id"] = user
	atClaims["exp"] = td.AtExpires
	atClaims["role_id"] = strconv.Itoa(role)

	//Creating token
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return td, nil
}
