package response

import "net/http"

// ref : https://stackoverflow.com/questions/73312823/wrap-api-json-response-in-go
type SuccessResponse struct {
	Data interface{} `json:"data"`
	StatusCode int `json:"status_code"`
}

func NewSuccessResponse(d interface{}, s int) *SuccessResponse {
	return &SuccessResponse{d, s}
}

func SuccessCreated(d interface{}) *SuccessResponse {
	return &SuccessResponse{d, http.StatusOK}
}