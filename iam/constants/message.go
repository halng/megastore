package constants

// Message constants
const (
	MessageSuccess       = "Success"
	MessageErrorBindJson = "Please check your input. Something went wrong"

	BadRequest          = "Bad Request"
	Unauthorized        = "Unauthorized"
	Forbidden           = "Forbidden"
	NotFound            = "Not Found"
	Conflict            = "Conflict"
	InternalServerError = "Internal Server Error"

	// account constant
	AccountCreated   = "Account created successfully"
	AccountNotFound  = "Account not found"
	AccountUpdated   = "Account updated successfully"
	AccountDeleted   = "Account deleted successfully"
	DefaultCreator   = "REGISTER"
	AccountExists    = "Account with username: %s or email: %s already exists"
	PasswordNotMatch = "Password doesn't match"
)
