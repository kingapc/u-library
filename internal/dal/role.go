package dal

import (
	"time"

	conn "github.com/rpinedafocus/u-library/internal/db"
	"github.com/rpinedafocus/u-library/internal/model"
	"github.com/rpinedafocus/u-library/internal/utils"
	"github.com/twinj/uuid"
)

func CreateRole(user string, role *model.Role) (*model.RoleEntity, error) {

	if role == nil {
		return nil, utils.EmtpyModel
	}

	now := time.Now()

	e := &model.RoleEntity{
		Entity: model.Entity{
			ID:        utils.RemoveHyphens(uuid.NewV4().String()),
			CreatedBy: user,
			CreatedAt: now.Format("01-02-2006"),
		},
		Role: model.Role{
			Name: role.Name,
		},
	}

	const stmt = `INSERT INTO university.roles (id, name, created_by, created_at) VALUES($1,$2,$3,$4) RETURNING role_id, active, is_deleted`

	db, errc := conn.GetConnection()
	if errc != nil {
		return nil, errc
	}

	var tempid int
	var temps utils.Flags
	err := db.QueryRow(stmt, e.Entity.ID, e.Role.Name, e.Entity.CreatedBy, e.Entity.CreatedAt).Scan(&tempid, &temps.Active, &temps.IsDeleted)
	if err != nil {
		db.Close()
		return nil, err
	}

	db.Close()

	e.RoleId = tempid
	e.Active = temps.Active
	e.IsDeleted = temps.IsDeleted

	return e, nil
}
