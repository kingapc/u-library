package model

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

type FetchBook struct {
	ID             string `json:"id"`
	Title          string `json:"title"`
	Author         string `json:"author"`
	Genre          string `json:"genre"`
	PublishDate    string `json:"publish_date"`
	TotalAvailable int    `json:"total_available"`
}

type BookEntity struct {
	Entity
	Book
}
