package dto

type ResDTO struct {
	StatusCode int      `json:"status_code"`
	Status     string   `json:"status"`
	Data       any      `json:"data"`
	Error      ErrorDTO `json:"error"`
}

type ErrorDTO struct {
	Message any `json:"msg"`
}

type RegisterRequest struct {
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ApiToken string `json:"api-token"`
	Username string `json:"username"`
	Email    string `json:"email"`
	ID       string `json:"id"`
}

type ActiveNewUser struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	Token       string `json:"token"`
	ExpiredTime int64  `json:"expiredTime"`
}
