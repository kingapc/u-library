package dal

import (
	"time"

	conn "github.com/rpinedafocus/u-library/pkg/db"
	"github.com/rpinedafocus/u-library/pkg/model"
	"github.com/rpinedafocus/u-library/pkg/utils"
	"github.com/twinj/uuid"
)

func CreateBooking(booking *model.BookingRent) (*model.BookingRentEntity, error) {

	if booking == nil {
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
			BookId:        booking.BookId,
			BookingRentBy: booking.BookingRentBy,
			BookingDate:   booking.BookingDate,
		},
	}

	const stmt = `INSERT INTO university.booking_rent (id,book_id,booking_rent_by,booking_date,created_by,created_at) VALUES($1,$2,$3,$4,$5,$6) RETURNING active, is_deleted`

	db, errc := conn.GetConnection()
	if errc != nil {
		return nil, errc
	}

	var temps utils.Flags
	err := db.QueryRow(stmt, e.Entity.ID, e.BookId, e.BookingRentBy, e.BookingDate, e.Entity.CreatedBy, e.Entity.CreatedAt).Scan(&temps.Active, &temps.IsDeleted)
	if err != nil {
		db.Close()
		return nil, err
	}

	db.Close()

	e.Active = temps.Active
	e.IsDeleted = temps.IsDeleted

	return e, nil
}

func ReleaseBookedBook(id string) (*model.BookingRentEntity, error) {

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

	const stmt = `UPDATE UNIVERSITY.BOOKING_RENT SET return_date = $1, updated_by = $2, updated_at = $3, returned = true WHERE BOOKING_DATE IS NOTNULL AND id = $4 
					RETURNING book_id,booking_rent_by,booking_date, returned, active`

	db, errc := conn.GetConnection()
	if errc != nil {
		return nil, errc
	}

	temp := &model.BookingRent{}
	err := db.QueryRow(stmt, e.BookingRent.ReturnDate, e.Entity.UpdatedBy, e.Entity.UpdatedAt, id).Scan(&temp.BookId, &temp.BookingRentBy, &temp.BookingDate, &temp.Returned, &temp.Active)

	if err != nil {
		db.Close()
		return nil, err
	}

	db.Close()

	e.BookId = temp.BookId
	e.BookingRentBy = temp.BookingRentBy
	e.RentDate = temp.BookingDate
	e.Returned = temp.Returned
	e.Active = temp.Active

	return e, nil
}

func IsValidBookin(id string) (bool, error) {

	const stmt = `SELECT 'X' FROM UNIVERSITY.BOOKING_RENT WHERE BOOKING_DATE IS NOT NULL AND ACTIVE = true AND IS_DELETED = 'N' AND RETURNED = false AND ID = $1`

	db, errc := conn.GetConnection()
	if errc != nil {
		return false, errc
	}

	var result string

	if db.QueryRow(stmt, id).Scan(&result); errc != nil {
		return false, errc
	}
	db.Close()

	return (result == "X"), nil
}

// func FetchAll() ([]*model.UserEntity, error) {

// 	const stmt = `SELECT id, first_name, last_name, created_at, updated_at, deleted_at FROM user`
// 	rows, err := u.DB.QueryContext(ctx, stmt)
// 	switch {
// 	case errors.Is(err, sql.ErrNoRows):
// 		return nil, errorx.ErrNoUser
// 	case err != nil:
// 		return nil, fmt.Errorf("user fetch query %w", err)
// 	default:
// 	}
// 	defer rows.Close()

// 	entities := []*model.UserEntity{}
// 	for rows.Next() {
// 		e := &model.UserEntity{}
// 		deletedAt := sql.NullTime{}
// 		if err := rows.Scan(&e.ID, &e.FirstName, &e.LastName, &e.CreatedAt, &e.UpdatedAt, &deletedAt); err != nil {
// 			return nil, fmt.Errorf("user row scan error %w", err)
// 		}
// 		if deletedAt.Valid == false {
// 			entities = append(entities, e)
// 		}
// 	}

// 	return entities, nil
// }
