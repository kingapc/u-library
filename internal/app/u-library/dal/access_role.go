package dal

import (
	"time"

	conn "github.com/rpinedafocus/u-library/internal/app/u-library/db"
	"github.com/rpinedafocus/u-library/internal/app/u-library/model"
	"github.com/rpinedafocus/u-library/internal/app/u-library/utils"
	"github.com/twinj/uuid"
)

func CreateAccessRole(user string, accrole *model.AccessRole) (*model.AccessRoleEntinty, error) {

	if accrole == nil {
		return nil, utils.EmtpyModel
	}

	now := time.Now()

	e := &model.AccessRoleEntinty{
		Entity: model.Entity{
			ID:        utils.RemoveHyphens(uuid.NewV4().String()),
			CreatedBy: user,
			CreatedAt: now.Format("01-02-2006"),
		},
		AccessRole: model.AccessRole{
			Endpoint: accrole.Endpoint,
			RoleId:   accrole.RoleId,
		},
	}

	const stmt = `INSERT INTO university.access_roles (id, endpoint, role_id, created_by, created_at) VALUES($1,$2,$3,$4,$5) RETURNING active, is_deleted`

	db, errc := conn.GetConnection()
	if errc != nil {
		return nil, errc
	}

	var temps utils.Flags
	err := db.QueryRow(stmt, e.Entity.ID, e.AccessRole.Endpoint, e.AccessRole.RoleId, e.Entity.CreatedBy, e.Entity.CreatedAt).Scan(&temps.Active, &temps.IsDeleted)
	if err != nil {
		db.Close()
		return nil, err
	}

	db.Close()

	e.Active = temps.Active
	e.IsDeleted = temps.IsDeleted

	return e, nil
}

func FetchAccessRolById(roleId int, endpoint string) bool {

	var stmt = `SELECT 'X' 
					FROM university.access_roles
					WHERE role_id = $1 AND endpoint = $2 AND active = true`

	db, errc := conn.GetConnection()
	if errc != nil {
		return false
	}

	var result string
	err := db.QueryRow(stmt, roleId, endpoint).Scan(&result)

	if err != nil {
		db.Close()
		return false
	}
	db.Close()

	return true
}
