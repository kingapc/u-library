package utils

import "errors"

var (
	KeyNotFound               = errors.New("Key env required")
	EnvNotLoaded              = errors.New("Unable to load env")
	EmtpyModel                = errors.New("Model is empty")
	ErrCreatingRow            = errors.New("Unable to create the register")
	DBConnectionError         = errors.New("Unable to connect to the data base")
	ErrNoDataFoun             = errors.New("No data found")
	FetchQueryC               = errors.New("Error fetching data")
	InvalidId                 = errors.New("Invalid Id")
	RowScanError              = errors.New("Row scan error")
	ErrSingMethod             = errors.New("unexpected signing method")
	ErrDeleteSession          = errors.New("Unable to delete session")
	ErrRoleIdIinvalid         = errors.New("Invalid Role Id")
	WarningNoBooksAvailables  = errors.New("No books availables")
	WarningNoRentAvailable    = errors.New("The rent selected is not available")
	WarningNoBookingAvailable = errors.New("The booking selected is not available")
	WarningRentExist          = errors.New("Book is already rented")
	WarningBookingExist       = errors.New("Book is already booked")
)
