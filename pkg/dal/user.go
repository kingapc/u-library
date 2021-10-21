package dal

import (
	"time"

	conn "github.com/rpinedafocus/u-library/pkg/db"
	"github.com/rpinedafocus/u-library/pkg/model"
	"github.com/rpinedafocus/u-library/pkg/utils"
	"github.com/u-library/pkg/utils"
)

const uuidLength = 36

type User struct {
	GenerateUUID GenerateUUID
}

func (u *User) Create(user *model.User) (*model.UserEntity, error) {

	if user == nil {
		return nil, utils.EmtpyModel
	}

	now := time.Now()

	e := &model.UserEntity{
		Entity: model.Entity{
			ID:        u.GenerateUUID(),
			CreatedBy: "root",
			CreatedAt: now,
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

	const stmt = `INSERT INTO university.users (id,user,password,first_name,last_name,email,role,created_by, created_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)`

	if db, errc := conn.GetConnection(); errc != nil {
		return nil, utils.DBCDBConnectionError
	}

	if _, err := db.Exec(stmt, e.Entity.ID, e.User.User, e.User.Password, e.User.FirstName, e.User.LastName, e.User.Email, e.User.Role, e.Entity.CreatedBy, e.Entity.CreatedAt); err != nil {
		return nil, utils.ErrCreatingRow
	}

	return e, nil
}
