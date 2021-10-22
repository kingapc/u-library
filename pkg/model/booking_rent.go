package model

// User is the structure for an user
type BookingRent struct {
	BookId        string `json:"book_id"`
	BookingRentBy string `json:"booking_rent_by"`
	BookingDate   string `json:"booking_date"`
	RentDate      string `json:"rent_date"`
	ReturnDate    string `json:"return_date"`
	Active        bool   `json:"active"`
}

// UserEntity is the user entity for the database
type BookingRentEntity struct {
	Entity
	BookingRent
}
