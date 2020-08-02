package success_response

import "net/http"

type SuccessResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// Found returns status OK
func Found(message string) *SuccessResponse {
	return &SuccessResponse{
		Message: message,
		Status:  http.StatusFound,
	}
}

// StatusOk returns status OK
func StatusOk(message string) *SuccessResponse {
	return &SuccessResponse{
		Message: message,
		Status:  http.StatusOK,
	}
}
