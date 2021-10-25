package model

type User struct {
	User      string `json:"user_name"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Role      int    `json:"role"`
	Active    bool   `json:"active"`
}

type Credentials struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type UserEntity struct {
	Entity
	User
}
