package dal

import (
	"time"

	conn "github.com/rpinedafocus/u-library/pkg/db"
	"github.com/rpinedafocus/u-library/pkg/model"
	"github.com/rpinedafocus/u-library/pkg/utils"
	"github.com/twinj/uuid"
)

func CreateGenre(user string, genre *model.Genre) (*model.GenreEntity, error) {

	if genre == nil {
		return nil, utils.EmtpyModel
	}

	now := time.Now()

	e := &model.GenreEntity{
		Entity: model.Entity{
			ID:        utils.RemoveHyphens(uuid.NewV4().String()),
			CreatedBy: user,
			CreatedAt: now.Format("01-02-2006"),
		},
		Genre: model.Genre{
			Name: genre.Name,
		},
	}

	const stmt = `INSERT INTO university.genres (id, name, created_by, created_at) VALUES($1,$2,$3,$4) RETURNING genre_id, active, is_deleted`

	db, errc := conn.GetConnection()
	if errc != nil {
		return nil, errc
	}

	var tempid int
	var temps utils.Flags
	err := db.QueryRow(stmt, e.Entity.ID, e.Genre.Name, e.Entity.CreatedBy, e.Entity.CreatedAt).Scan(&tempid, &temps.Active, &temps.IsDeleted)
	if err != nil {
		db.Close()
		return nil, err
	}

	db.Close()

	e.GenreId = tempid
	e.Active = temps.Active
	e.IsDeleted = temps.IsDeleted

	return e, nil
}
