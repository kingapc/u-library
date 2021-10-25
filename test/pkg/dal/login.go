package dal

import (
	"database/sql"
	"errors"
	"fmt"

	conn "github.com/rpinedafocus/u-library/pkg/db"
	"github.com/rpinedafocus/u-library/pkg/model"
	"github.com/rpinedafocus/u-library/pkg/utils"
)

func Login(user string, password string) (*model.UserEntity, error) {

	db, err := conn.GetConnection()
	if err != nil {
		return nil, err
	}

	const stmt = ` SELECT id, user_name, first_name, last_name, email, role, active FROM UNIVERSITY.users where user_name = $1 and pwd = crypt($2, pwd)`

	e := &model.UserEntity{}
	err = db.QueryRow(stmt, user, password).Scan(&e.ID, &e.User.User, &e.FirstName, &e.LastName, &e.Email, &e.Role, &e.Active)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, utils.ErrNoDataFoun
	case err != nil:
		return nil, fmt.Errorf(utils.FetchQueryC.Error()+"%w", err)
	default:
		return e, nil
	}
}
