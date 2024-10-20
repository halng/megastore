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
	InternalServerError = "There was an error processing your request. Please try again later"
	TokenNotFount       = "Your login session has expired. Please login again"
	MissingCredentials  = "Missing credentials. X-API-SECRET-TOKEN and X-API-USER-ID are required"
	// account constant
	AccountCreated    = "Account created successfully"
	AccountNotFound   = "Account not found"
	AccountUpdated    = "Account updated successfully"
	AccountDeleted    = "Account deleted successfully"
	DefaultCreator    = "SYSTEM"
	AccountExists     = "Account with username: %s or email: %s already exists"
	PasswordNotMatch  = "Invalid credentials"
	InvalidPermission = "User does not have permission to access this resource"
	DefaultPassword   = "12345678"
	// define key
	ApiTokenRequestHeader  = "X-API-SECRET-TOKEN"
	ApiUserIdRequestHeader = "X-API-USER-ID"
	ApiUserRequestHeader   = "X-API-USER"
	ApiUserRoles           = "X-API-USER-ROLES"
)
