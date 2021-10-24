package securitty

import (
	"time"

	"github.com/rpinedafocus/u-library/pkg/model"
)

func CreateSession(id string, td *model.TokenDetails) error {

	at := time.Unix(td.AtExpires, 0)
	//rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := Client.Set(td.AccessUuid, id, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}

	// errRefresh := Client.Set(td.RefreshUuid, id, rt.Sub(now)).Err()
	// if errRefresh != nil {
	// 	return errRefresh
	// }

	return nil
}
