package dal

import (
	"database/sql"
	"errors"
	"fmt"

	conn "github.com/rpinedafocus/u-library/internal/db"
	"github.com/rpinedafocus/u-library/internal/model"
	"github.com/rpinedafocus/u-library/internal/utils"
)

func Login(user string, password string) (*model.UserEntity, error) {

	const sqlStmt = ` SELECT id, user_name, first_name, last_name, email, role, active FROM UNIVERSITY.users where user_name = $1 and pwd = crypt($2, pwd)`

	db, err := conn.GetConnection()
	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		return nil, fmt.Errorf((utils.ErrStmt.Error() + "%w"), err)
	}

	e := &model.UserEntity{}
	err = stmt.QueryRow(user, password).Scan(&e.ID, &e.User.User, &e.FirstName, &e.LastName, &e.Email, &e.Role, &e.Active)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, utils.ErrNoDataFoun
	case err != nil:
		return nil, fmt.Errorf(utils.FetchQueryC.Error()+"%w", err)
	default:
		return e, nil
	}
}
