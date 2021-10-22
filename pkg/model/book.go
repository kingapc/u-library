package model

// User is the structure for an user
type Book struct {
	Title          string `json:"title"`
	AuthorId       int    `json:"author_id"`
	GenreId        int    `json:"genre_id"`
	PublishDate    string `json:"publish_date"`
	RentedNumber   int    `json:"rented_number"`
	BookingNumber  int    `json:"booking_number"`
	TotalAvailable int    `json:"total_available"`
	Active         bool   `json:"active"`
}

// UserEntity is the user entity for the database
type BookEntity struct {
	Entity
	Book
}
