package model

type Author struct {
	AuthorId int    `json:"author_id"`
	Name     string `json:"name"`
	Active   bool   `json:"active"`
}

type AuthorEntity struct {
	Entity
	Author
}
