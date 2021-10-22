package model

// User is the structure for an user
type Author struct {
	AuthorId int    `json:"author_id"`
	Name     string `json:"name"`
	Active   bool   `json:"active"`
}

// UserEntity is the user entity for the database
type AuthorEntity struct {
	Entity
	Author
}
