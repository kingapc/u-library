package securitty

import (
	"github.com/rpinedafocus/u-library/internal/app/u-library/utils"
)

func DestroySession(access string) error {

	//delete access token
	deletedAt, err := Client.Del(access).Result()
	if err != nil {
		return err
	}

	//When the record is deleted, the return value is 1
	if deletedAt != 1 {
		return utils.ErrDeleteSession
	}

	return nil
}
