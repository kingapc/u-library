package model

// User is the structure for an user
type Genre struct {
	GenreId int    `json:"genre_id"`
	Name    string `json:"name"`
	Active  bool   `json:"active"`
}

// UserEntity is the user entity for the database
type GenreEntity struct {
	Entity
	Genre
}
