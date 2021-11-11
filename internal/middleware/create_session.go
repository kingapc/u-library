package securitty

import (
	"time"

	"github.com/rpinedafocus/u-library/internal/model"
)

func CreateSession(id string, td *model.TokenDetails) error {

	at := time.Unix(td.AtExpires, 0)
	now := time.Now()

	errAccess := Client.Set(td.AccessUuid, id, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}

	return nil
}
