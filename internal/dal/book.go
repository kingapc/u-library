package dal

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	conn "github.com/rpinedafocus/u-library/internal/db"
	"github.com/rpinedafocus/u-library/internal/model"
	"github.com/rpinedafocus/u-library/internal/utils"
	"github.com/twinj/uuid"
)

func CreateBook(user string, book *model.Book) (*model.BookEntity, error) {

	if book == nil {
		return nil, utils.EmtpyModel
	}

	now := time.Now()

	e := &model.BookEntity{
		Entity: model.Entity{
			ID:        utils.RemoveHyphens(uuid.NewV4().String()),
			CreatedBy: user,
			CreatedAt: now.Format("01-02-2006"),
		},
		Book: model.Book{
			Title:          book.Title,
			AuthorId:       book.AuthorId,
			GenreId:        book.GenreId,
			PublishDate:    book.PublishDate,
			TotalAvailable: book.TotalAvailable,
		},
	}

	const stmt = `INSERT INTO university.books (id,title,author_id,genre_id,publish_date,total_available,created_by,created_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8) RETURNING active, is_deleted`

	db, errc := conn.GetConnection()
	if errc != nil {
		return nil, errc
	}

	var temps utils.Flags
	err := db.QueryRow(stmt, e.Entity.ID, e.Book.Title, e.Book.AuthorId, e.Book.GenreId, e.Book.PublishDate, e.Book.TotalAvailable, e.Entity.CreatedBy, e.Entity.CreatedAt).Scan(&temps.Active, &temps.IsDeleted)
	if err != nil {
		db.Close()
		return nil, err
	}

	db.Close()

	e.Active = temps.Active
	e.IsDeleted = temps.IsDeleted

	return e, nil
}

func FetchAll(p string, v string) ([]*model.FetchBook, error) {

	var rows *sql.Rows
	var err error
	var condition string

	switch p {
	case "title":
		condition = `AND LOWER(b.title) like $1`
	case "author":
		condition = `AND LOWER(a.name) like $1`
	case "genre":
		condition = `AND LOWER(g.name) like $1`
	default:
		condition = ""
	}

	var sqlStmt = `SELECT b.id, b.title, a.name author, g.name genre, b.publish_date, b.total_available 
					FROM university.books b
					INNER JOIN university.authors a on a.author_id = b.author_id
					INNER JOIN university.genres g on g.genre_id = b.genre_id
					WHERE b.active = true AND b.is_deleted = 'N' ` + condition

	db, errc := conn.GetConnection()
	if errc != nil {
		return nil, errc
	}

	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		return nil, fmt.Errorf((utils.ErrStmt.Error() + "%w"), err)
	}

	if condition != "" {
		rows, err = stmt.Query(`%` + v + `%`)
	} else {
		rows, err = stmt.Query()
	}

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, utils.ErrNoDataFoun
	case err != nil:
		return nil, fmt.Errorf((utils.FetchQueryC.Error() + "%w"), err)
	default:
	}
	defer rows.Close()

	entities := []*model.FetchBook{}
	for rows.Next() {
		e := &model.FetchBook{}
		if err := rows.Scan(&e.ID, &e.Title, &e.Author, &e.Genre, &e.PublishDate, &e.TotalAvailable); err != nil {
			return nil, fmt.Errorf(utils.FetchQueryC.Error()+"%w", err)
		}

		entities = append(entities, e)
	}

	return entities, nil
}

func FetchBookById(id string) (*model.FetchBook, error) {

	var sqlStmt = `SELECT b.id, b.title, a.name author, g.name genre, b.publish_date, b.total_available 
					FROM university.books b
					INNER JOIN university.authors a on a.author_id = b.author_id
					INNER JOIN university.genres g on g.genre_id = b.genre_id
					WHERE b.active = true AND b.is_deleted = 'N' AND b.id = $1`

	db, errc := conn.GetConnection()
	if errc != nil {
		return nil, errc
	}

	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		return nil, fmt.Errorf((utils.ErrStmt.Error() + "%w"), err)
	}

	row := stmt.QueryRow(id)

	switch {
	case errors.Is(row.Err(), sql.ErrNoRows):
		return nil, utils.ErrNoDataFoun
	case row.Err() != nil:
		return nil, fmt.Errorf((utils.FetchQueryC.Error() + "%w"), row.Err().Error())
	default:
	}

	e := &model.FetchBook{}
	err = row.Scan(&e.ID, &e.Title, &e.Author, &e.Genre, &e.PublishDate, &e.TotalAvailable)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, utils.ErrNoDataFoun
	case err != nil:
		return nil, fmt.Errorf(utils.FetchQueryC.Error()+"%w", err)
	default:
		return e, nil
	}
}
