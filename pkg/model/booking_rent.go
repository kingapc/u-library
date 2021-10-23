package model

import "database/sql"

//basic structure
type BookingRent struct {
	BookId        string `json:"book_id"`
	BookingRentBy string `json:"booking_rent_by"`
	BookingDate   string `json:"booking_date"`
	RentDate      string `json:"rent_date"`
	ReturnDate    string `json:"return_date"`
	Active        bool   `json:"active"`
	Returned      bool   `json:"returned"`
}

type MyBookingRent struct {
	ID            string         `json:"id"`
	BookingRentBy string         `json:"booking_rent_by"`
	BookId        string         `json:"book_id"`
	Title         string         `json:"title"`
	PublishDate   sql.NullString `json:"publish_date,omitempty"`
	Author        string         `json:"author"`
	Genre         string         `json:"genre"`
	BookingDate   sql.NullString `json:"booking_date,omitempty"`
	RentDate      sql.NullString `json:"rent_date,omitempty"`
	Status        string         `json:"status"`
}

type BookingRentEntity struct {
	Entity
	BookingRent
}
