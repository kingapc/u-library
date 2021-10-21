package utils

import "errors"

var (
	KeyNotFound   = errors.New("Key env required")
	EnvNotLoaded  = errors.New("Unable to load env")
	ErrDeleteUser = errors.New("user has been deleted")
)
