package models

type ErrorResponse struct {
	ErrorCode int    `json:"error_code"`
	Status    string `json:"status"`
	Message   string `json:"message"`
}

type SuccessResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
