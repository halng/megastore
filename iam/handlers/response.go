package handlers

type ResDTO struct {
	StatusCode int      `json:"status_code"`
	Status     string   `json:"status"`
	Data       any      `json:"data"`
	Error      ErrorDTO `json:"error"`
}

type ErrorDTO struct {
	Message string `json:"msg"`
}

func GetResponseDTO(code int, data any, err ErrorDTO) any {
	var res ResDTO
	res.StatusCode = code
	res.Data = data
	res.Error = err
	res.Status = map[bool]string{true: "SUCCESS", false: "ERROR"}[code/100 == 2]

	return res

}
