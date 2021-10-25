package model

import "database/sql"

//basic structure
type BookingRent struct {
	BookId      string `json:"book_id"`
	ProcessBy   string `json:"process_by"`
	ProcessDate string `json:"process_date"`
	ReturnDate  string `json:"return_date"`
	IsBooking   bool   `json:"is_booking"`
}

type MyBookingRent struct {
	ID          string         `json:"id"`
	ProcessBy   string         `json:"process_by"`
	BookId      string         `json:"book_id"`
	Title       string         `json:"title"`
	PublishDate sql.NullString `json:"publish_date,omitempty"`
	Author      string         `json:"author"`
	Genre       string         `json:"genre"`
	ProcessDate sql.NullString `json:"process_date,omitempty"`
	Status      string         `json:"status"`
}

type BookingRentEntity struct {
	Entity
	BookingRent
}
