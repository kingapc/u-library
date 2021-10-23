package dal

import (
	"database/sql"
	"errors"
	"time"

	conn "github.com/rpinedafocus/u-library/pkg/db"
	"github.com/rpinedafocus/u-library/pkg/model"
	"github.com/rpinedafocus/u-library/pkg/utils"
	"github.com/twinj/uuid"
)

func CreateRent(rent *model.BookingRent) (*model.BookingRentEntity, error) {

	if rent == nil {
		return nil, utils.EmtpyModel
	}

	now := time.Now()

	e := &model.BookingRentEntity{
		Entity: model.Entity{
			ID:        utils.RemoveHyphens(uuid.NewV4().String()),
			CreatedBy: "root",
			CreatedAt: now.Format("01-02-2006"),
		},
		BookingRent: model.BookingRent{
			BookId:        rent.BookId,
			BookingRentBy: rent.BookingRentBy,
			RentDate:      rent.RentDate,
		},
	}

	const stmt = `INSERT INTO university.booking_rent (id,book_id,booking_rent_by,rent_date,created_by,created_at) VALUES($1,$2,$3,$4,$5,$6) RETURNING active, is_deleted`

	db, errc := conn.GetConnection()
	if errc != nil {
		return nil, errc //utils.DBConnectionError
	}

	var temps utils.Flags
	err := db.QueryRow(stmt, e.Entity.ID, e.BookId, e.BookingRentBy, e.RentDate, e.Entity.CreatedBy, e.Entity.CreatedAt).Scan(&temps.Active, &temps.IsDeleted)
	if err != nil {
		db.Close()
		return nil, err //utils.ErrCreatingRow
	}

	db.Close()

	e.Active = temps.Active
	e.IsDeleted = temps.IsDeleted

	return e, nil
}

func ReturnRentedBook(id string) (*model.BookingRentEntity, error) {

	now := time.Now()

	e := &model.BookingRentEntity{
		Entity: model.Entity{
			ID:        id,
			UpdatedBy: "root",
			UpdatedAt: now.Format("01-02-2006"),
		},
		BookingRent: model.BookingRent{
			ReturnDate: now.Format("01-02-2006"),
		},
	}

	const stmt = `UPDATE UNIVERSITY.BOOKING_RENT SET return_date = $1, updated_by = $2, updated_at = $3, returned = true WHERE RENT_DATE IS NOT NULL AND id = $4 
					RETURNING book_id,booking_rent_by,rent_date, returned, active`

	db, errc := conn.GetConnection()
	if errc != nil {
		return nil, errc
	}

	temp := &model.BookingRent{}
	err := db.QueryRow(stmt, e.BookingRent.ReturnDate, e.Entity.UpdatedBy, e.Entity.UpdatedAt, id).Scan(&temp.BookId, &temp.BookingRentBy, &temp.RentDate, &temp.Returned, &temp.Active)

	if err != nil {
		db.Close()
		return nil, err
	}

	db.Close()

	e.BookId = temp.BookId
	e.BookingRentBy = temp.BookingRentBy
	e.RentDate = temp.RentDate
	e.Returned = temp.Returned
	e.Active = temp.Active

	return e, nil
}

func IsValidRent(id string) (bool, error) {

	const stmt = `SELECT 'X' FROM UNIVERSITY.BOOKING_RENT WHERE RENT_DATE IS NOT NULL AND ACTIVE = true AND IS_DELETED = 'N' AND RETURNED = false AND ID = $1`

	db, errc := conn.GetConnection()
	if errc != nil {
		return false, errc
	}

	var result string
	err := db.QueryRow(stmt, id).Scan(&result)
	if err != nil && !(errors.Is(err, sql.ErrNoRows)) {
		db.Close()
		return false, err
	}
	db.Close()

	return (result == "X"), nil
}
