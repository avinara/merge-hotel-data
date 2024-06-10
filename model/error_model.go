package model

type ErrorResponse struct {
	Code         int    `json:"code"`
	ErrorCode    uint32 `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}
