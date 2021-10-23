package securitty

import (
	"github.com/rpinedafocus/u-library/pkg/dal"
	"github.com/rpinedafocus/u-library/pkg/model"
	"github.com/rpinedafocus/u-library/pkg/utils"
)

func PrepareLogin(user string, password string) (*model.UserEntity, error) {

	//hash password
	pass, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	e, err := dal.Login(user, pass)
	if err != nil {
		return nil, err
	}

	return e, nil
}
