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

func GetResponseDTO(code int, data any, err ErrorDTO) any {
	var res ResDTO
	res.StatusCode = code
	res.Data = data
	res.Error = err
	res.Status = map[bool]string{true: "SUCCESS", false: "ERROR"}[code/100 == 2]

	return res

}

type RegisterRequest struct {
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginResponse struct {
	ApiToken string `json:"api-token"`
}

type ActiveNewUser struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	Token       string `json:"token"`
	ExpiredTime int64  `json:"expiredTime"`
}
