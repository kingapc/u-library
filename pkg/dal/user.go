package dal

import (
	"time"

	conn "github.com/rpinedafocus/u-library/pkg/db"
	"github.com/rpinedafocus/u-library/pkg/model"
	"github.com/rpinedafocus/u-library/pkg/utils"
	"github.com/twinj/uuid"
)

func CreateUser(user *model.User) (*model.UserEntity, error) {

	if user == nil {
		return nil, utils.EmtpyModel
	}

	now := time.Now()

	e := &model.UserEntity{
		Entity: model.Entity{
			ID:        utils.RemoveHyphens(uuid.NewV4().String()),
			CreatedBy: "root",
			CreatedAt: now.Format("01-02-2006"),
		},
		User: model.User{
			User:      user.User,
			Password:  user.Password,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Role:      user.Role,
		},
	}

	const stmt = `INSERT INTO university.users (id,"user",password,first_name,last_name,email,role,created_by, created_at) VALUES($1,$2,crypt($3, gen_salt('bf')),$4,$5,$6,$7,$8,$9) RETURNING active, is_deleted`

	db, errc := conn.GetConnection()
	if errc != nil {
		return nil, utils.DBConnectionError
	}

	var temps utils.Flags
	err := db.QueryRow(stmt, e.Entity.ID, e.User.User, e.User.Password, e.User.FirstName, e.User.LastName, e.User.Email, e.User.Role, e.Entity.CreatedBy, e.Entity.CreatedAt).Scan(&temps.Active, &temps.IsDeleted)
	if err != nil {
		db.Close()
		return nil, utils.ErrCreatingRow
	}

	db.Close()

	e.Active = temps.Active
	e.IsDeleted = temps.IsDeleted

	return e, nil
}
