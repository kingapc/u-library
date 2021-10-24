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

	//const stmt = `SELECT id, "user", first_name, last_name, email, role, active FROM UNIVERSITY.USERS where "user" = 'user1' and "password" = '$2a$14$HRdKWvlphoAMtVtNZNXmZOmhifOVBnXwjy9Q6V7gZ.DKkRAxRNM8m'`
	// const crypto = `CREATE EXTENSION pgcrypto`
	// _, err = db.Exec(crypto)

	// if err != nil {
	// 	return nil, err
	// }

	const stmt = ` SELECT id, user_name, first_name, last_name, email, role, active FROM UNIVERSITY.users where user_name = $1 and pwd = crypt($2, pwd)`
	e := &model.UserEntity{}
	//err = db.QueryRow(stmt).Scan(&e.Entity.ID, &e.User.User, &e.User.FirstName, &e.User.LastName, &e.User.Email, &e.User.Role, &e.User.Active)
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
