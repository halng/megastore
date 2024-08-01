package constants

// Message constants
const (
	MessageSuccess = "Success"
	MessageError   = "Something went wrong. Please try again later."

	BadRequest          = "Bad Request"
	Unauthorized        = "Unauthorized"
	Forbidden           = "Forbidden"
	NotFound            = "Not Found"
	Conflict            = "Conflict"
	InternalServerError = "Internal Server Error"

	// account constant
	AccountCreated  = "Account created successfully"
	AccountExist    = "Account already exist"
	AccountNotFound = "Account not found"
	AccountUpdated  = "Account updated successfully"
	AccountDeleted  = "Account deleted successfully"
	DefaultCreator  = "REGISTER"
	AccountExists   = "Account with username: %s or email: %s already exists"
)
