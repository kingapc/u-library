package utils

import "errors"

var (
	KeyNotFound       = errors.New("Key env required")
	EnvNotLoaded      = errors.New("Unable to load env")
	EmtpyModel        = errors.New("Model is empty")
	ErrCreatingRow    = errors.New("Unable to create the register")
	DBConnectionError = errors.New("Unable to connect to the data base")
)
