package utils

type SuccessWrapper struct {
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	StatusCode int         `json:"statusCode"`
}

type ErrorWrapper struct {
	Error      interface{} `json:"error"`
	Message    string      `json:"message"`
	StatusCode int         `json:"statusCode"`
}
