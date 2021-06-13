package utils

import "net/http"

type RestSuccess struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Success string `json:"success"`
}

func SuccessOkResponse(message string) *RestSuccess {
	return &RestSuccess{
		Message: message,
		Status:  http.StatusOK,
		Success: "Success_Status_Ok",
	}
}
