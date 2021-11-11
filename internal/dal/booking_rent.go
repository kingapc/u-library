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

func CreateBookingRent(user string, br *model.BookingRent) (*model.BookingRentEntity, error) {

	if br == nil {
		return nil, utils.EmtpyModel
	}

	now := time.Now()

	if !br.IsBooking {
		br.ProcessDate = now.Format("01-02-2006")
	}

	e := &model.BookingRentEntity{
		Entity: model.Entity{
			ID:        utils.RemoveHyphens(uuid.NewV4().String()),
			CreatedBy: user,
			CreatedAt: now.Format("01-02-2006"),
		},
		BookingRent: model.BookingRent{
			BookId:      br.BookId,
			ProcessBy:   br.ProcessBy,
			ProcessDate: br.ProcessDate,
			IsBooking:   br.IsBooking,
		},
	}

	const stmt = `INSERT INTO university.booking_rent (id,book_id,process_by,process_date,is_booking,created_by,created_at) VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING is_deleted`

	db, errc := conn.GetConnection()
	if errc != nil {
		return nil, errc
	}

	var temps utils.Flags
	err := db.QueryRow(stmt, e.Entity.ID, e.BookId, e.ProcessBy, e.ProcessDate, e.IsBooking, e.Entity.CreatedBy, e.Entity.CreatedAt).Scan(&temps.IsDeleted)
	if err != nil {
		db.Close()
		return nil, err
	}

	db.Close()

	e.IsDeleted = temps.IsDeleted

	return e, nil
}

func ReleaseBookingRent(user string, id string) (*model.BookingRentEntity, error) {

	now := time.Now()

	e := &model.BookingRentEntity{
		Entity: model.Entity{
			ID:        id,
			UpdatedBy: user,
			UpdatedAt: now.Format("01-02-2006"),
		},
		BookingRent: model.BookingRent{
			ReturnDate: now.Format("01-02-2006"),
		},
	}

	const stmt = `UPDATE UNIVERSITY.BOOKING_RENT SET return_date = $1, updated_by = $2, updated_at = $3 WHERE id = $4 
					RETURNING book_id,process_by,process_date,is_deleted`

	db, errc := conn.GetConnection()
	if errc != nil {
		return nil, errc
	}

	temp := &model.BookingRentEntity{}
	err := db.QueryRow(stmt, e.BookingRent.ReturnDate, e.Entity.UpdatedBy, e.Entity.UpdatedAt, id).Scan(&temp.BookId, &temp.ProcessBy, &temp.ProcessDate, &temp.Entity.IsDeleted)

	if err != nil {
		db.Close()
		return nil, err
	}

	db.Close()

	e.BookId = temp.BookId
	e.ProcessBy = temp.ProcessBy
	e.ProcessDate = temp.ProcessDate
	e.Entity.IsDeleted = temp.Entity.IsDeleted

	return e, nil
}

func FetchBookingRent(id string) ([]*model.MyBookingRent, error) {

	var sqlStmt = `SELECT a.id ID, a.process_by, a.book_id BOOK_ID, b.title TITLE, b.publish_date PUBLISH_DATE, c.name AUTHOR, d.name GENRE, a.process_date,
				CASE WHEN a.is_booking IS NULL THEN 'booked' ELSE 'rented' END STATUS
				FROM university.BOOKING_RENT a 
				INNER JOIN university.books b ON a.book_id = b.id 
				INNER JOIN university.authors c ON c.author_id = b.author_id 
				INNER JOIN university.genres d ON d.genre_id = b.genre_id 
				WHERE a.return_date IS NULL AND a.is_deleted = 'N' AND a.process_by = $1`

	db, errc := conn.GetConnection()
	if errc != nil {
		return nil, errc
	}

	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		return nil, fmt.Errorf((utils.ErrStmt.Error() + "%w"), err)
	}

	rows, err := stmt.Query(id)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, utils.ErrNoDataFoun
	case err != nil:
		return nil, fmt.Errorf((utils.FetchQueryC.Error() + "%w"), err)
	default:
	}
	defer rows.Close()

	entities := []*model.MyBookingRent{}
	for rows.Next() {
		e := &model.MyBookingRent{}
		if err := rows.Scan(&e.ID, &e.ProcessBy, &e.BookId, &e.Title, &e.PublishDate, &e.Author, &e.Genre, &e.ProcessDate, &e.Status); err != nil {
			return nil, fmt.Errorf(utils.FetchQueryC.Error()+"%w", err)
		}

		entities = append(entities, e)
	}
	db.Close()

	return entities, nil
}

func ExistValidBookingRent(bookId string, processBy string) (bool, error) {

	const sqlStmt = `SELECT 'X' 
					FROM UNIVERSITY.BOOKING_RENT
					WHERE 
						RETURN_DATE IS NULL
					AND IS_DELETED = 'N'
					AND BOOK_ID = $1
					AND PROCESS_BY = $2`

	db, err := conn.GetConnection()
	if err != nil {
		return false, err
	}

	var result string
	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		return false, fmt.Errorf((utils.ErrStmt.Error() + "%w"), err)
	}

	err = stmt.QueryRow(bookId, processBy).Scan(&result)

	if err != nil && !(errors.Is(err, sql.ErrNoRows)) {
		db.Close()
		return false, err
	}
	db.Close()

	return (result == "X"), nil
}

func IsValidBookingRent(id string) (bool, error) {

	const sqlStmt = `SELECT 'X' FROM UNIVERSITY.BOOKING_RENT WHERE RETURN_DATE IS NOT NULL AND IS_DELETED = 'N' AND ID = $1`

	db, errc := conn.GetConnection()
	if errc != nil {
		return false, errc
	}

	var result string
	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		return false, fmt.Errorf((utils.ErrStmt.Error() + "%w"), err)
	}

	if stmt.QueryRow(id).Scan(&result); errc != nil {
		return false, errc
	}
	db.Close()

	return (result == "X"), nil
}
