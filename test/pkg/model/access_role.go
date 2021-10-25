package model

type AccessRole struct {
	Endpoint string `json:"endpoint"`
	RoleId   int    `json:"role_id"`
	Active   bool   `jsong:"active"`
}

type AccessRoleEntinty struct {
	Entity
	AccessRole
}
