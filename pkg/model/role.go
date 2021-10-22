package model

// User is the structure for an user
type Role struct {
	RoleId int    `json:"role_id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

// UserEntity is the user entity for the database
type RoleEntity struct {
	Entity
	Role
}
