package dal

import (
	"errors"
	"fmt"
	"time"

	model "../model"
	conn "github.com/rpinedafocus/mylib-dbconn"
)

const uuidLength = 36

type User struct {
	GenerateUUID GenerateUUID
}

func (u *User) Create(user *model.User) (*model.UserEntity, error) {

	if user == nil {
		return nil, errors.New("user can not be nil")
	}

	now := time.Now()

	e := &model.UserEntity{
		Entity: model.Entity{
			ID:        u.GenerateUUID(),
			CreatedBy: "",
			CreatedAt: now,
		},
		User: model.User{
			User:      user.User,
			Password:  user.Password,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Role:      user.Role,
			Active:    true,
		},
	}

	const stmt = `INSERT INTO university.users (id,user,password,first_name,last_name,email,role,created_by, created_at,active) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`

	db, errc := conn.GetConnection()

	if _, err := u.DB.ExecContext(ctx, stmt, e.ID, e.FirstName, e.LastName, e.CreatedAt, e.UpdatedAt); err != nil {
		return nil, fmt.Errorf("user create insert %w", err)
	}
	return e, nil
}
