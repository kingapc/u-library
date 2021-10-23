package model

type Genre struct {
	GenreId int    `json:"genre_id"`
	Name    string `json:"name"`
	Active  bool   `json:"active"`
}

type GenreEntity struct {
	Entity
	Genre
}
