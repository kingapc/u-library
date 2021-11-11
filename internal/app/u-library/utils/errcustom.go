package utils

import "errors"

var (
	KeyNotFound                   = errors.New("Key env required")
	EnvNotLoaded                  = errors.New("Unable to load env")
	EmtpyModel                    = errors.New("Model is empty")
	ErrCreatingRow                = errors.New("Unable to create the register")
	DBConnectionError             = errors.New("Unable to connect to the data base")
	ErrNoDataFoun                 = errors.New("No data found")
	FetchQueryC                   = errors.New("Error fetching data")
	InvalidId                     = errors.New("Invalid Id")
	RowScanError                  = errors.New("Row scan error")
	ErrStmt                       = errors.New("Error in the statement")
	ErrSingMethod                 = errors.New("unexpected signing method")
	ErrDeleteSession              = errors.New("Unable to delete session")
	ErrRoleIdIinvalid             = errors.New("Invalid Role Id")
	ErrNoSessionInformation       = errors.New("Unable to found session information")
	ErrRequiredDateProcess        = errors.New("Booking date is required")
	ErrExpiringToken              = errors.New("Unable to expire token.")
	ErrExpiredToken               = errors.New("Token expired.")
	WarningNoBooksAvailables      = errors.New("No books available")
	WarningNoRentAvailable        = errors.New("The rent selected is not available")
	WarningNoBookingRentAvailable = errors.New("The booking or rent selected was released")
	WarningBookingRentExist       = errors.New("Book is already booked or rented")
)
