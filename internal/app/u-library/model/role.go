package model

type Role struct {
	RoleId int    `json:"role_id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type RoleEntity struct {
	Entity
	Role
}
