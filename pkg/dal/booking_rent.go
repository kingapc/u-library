package dal

import (
	"database/sql"
	"errors"
	"fmt"

	conn "github.com/rpinedafocus/u-library/pkg/db"
	"github.com/rpinedafocus/u-library/pkg/model"
	"github.com/rpinedafocus/u-library/pkg/utils"
)

func FetchBookingRent(id string) ([]*model.MyBookingRent, error) {

	var stmt = `SELECT a.id ID, a.booking_rent_by, a.book_id BOOK_ID, b.title TITLE, b.publish_date PUBLISH_DATE, c.name AUTHOR, d.name GENRE, a.booking_date, a.rent_date,
				CASE WHEN a.booking_date IS NULL THEN 'RENTED' ELSE 'BOOKED' END STATUS
				FROM university.BOOKING_RENT a 
				INNER JOIN university.books b ON a.book_id = b.id 
				INNER JOIN university.authors c ON c.author_id = b.author_id 
				INNER JOIN university.genres d ON d.genre_id = b.genre_id 
				WHERE a.returned = false AND a.is_deleted = 'N' AND a.booking_rent_by = $1`

	db, errc := conn.GetConnection()
	if errc != nil {
		return nil, errc
	}

	rows, err := db.Query(stmt, id)

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
		if err := rows.Scan(&e.ID, &e.BookingRentBy, &e.BookId, &e.Title, &e.PublishDate, &e.Author, &e.Genre, &e.BookingDate, &e.RentDate, &e.Status); err != nil {
			return nil, fmt.Errorf(utils.FetchQueryC.Error()+"%w", err)
		}

		entities = append(entities, e)
	}
	db.Close()

	return entities, nil
}

func ExistValidRent(bookId string, rentBy string) (bool, error) {

	const stmt = `SELECT 'X' FROM UNIVERSITY.BOOKING_RENT 
					WHERE RENT_DATE IS NOT NULL 
					AND RETURN_DATE IS NULL 
					AND ACTIVE = true 
					AND RETURNED = false 
					AND IS_DELETED = 'N'
					AND BOOK_ID = $1
					AND BOOKING_RENT_BY = $2`

	db, err := conn.GetConnection()
	if err != nil {
		return false, err
	}

	var result string
	err = db.QueryRow(stmt, bookId, rentBy).Scan(&result)
	if err != nil && !(errors.Is(err, sql.ErrNoRows)) {
		db.Close()
		return false, err
	}
	db.Close()

	return (result == "X"), nil
}

func ExistValidBooking(bookId string, rentBy string) (bool, error) {

	const stmt = `SELECT 'X' FROM UNIVERSITY.BOOKING_RENT 
					WHERE BOOKING_DATE IS NOT NULL 
					AND RETURN_DATE IS NULL 
					AND ACTIVE = true 
					AND RETURNED = false 
					AND IS_DELETED = 'N'
					AND BOOK_ID = $1
					AND BOOKING_RENT_BY = $2`

	db, err := conn.GetConnection()
	if err != nil {
		return false, err
	}

	var result string
	err = db.QueryRow(stmt, bookId, rentBy).Scan(&result)
	if err != nil && !(errors.Is(err, sql.ErrNoRows)) {
		db.Close()
		return false, err
	}
	db.Close()

	return (result == "X"), nil
}
