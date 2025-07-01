package apperror

import "net/http"

// ref : https://stackoverflow.com/questions/73312823/wrap-api-json-response-in-go
type ErrResponse struct {
	Data       interface{} `json:"data"`
	StatusCode int         `json:"status_code"`
}

func NewErrResponse(d interface{}, s int) *ErrResponse {
	return &ErrResponse{d, s}
}

func ErrInvalidRequest(d interface{}) *ErrResponse {
	return &ErrResponse{d, http.StatusBadRequest}
}

func ErrInternalServer(d interface{}) *ErrResponse {
	return &ErrResponse{d, http.StatusInternalServerError}
}

func ErrUnauthorized(d interface{}) *ErrResponse{
	return &ErrResponse{d, http.StatusUnauthorized}
}