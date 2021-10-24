package model

// User is the structure for an user
type AccessRole struct {
	Endpoint string `json:"endpoint"`
	RoleId   int    `json:"role_id"`
	Active   bool   `jsong:"active"`
}

// UserEntity is the user entity for the database
type AccessRoleEntinty struct {
	Entity
	AccessRole
}
