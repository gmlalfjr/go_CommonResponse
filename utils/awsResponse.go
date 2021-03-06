package utils

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func SuccessResponse(data *SuccessWrapper) (events.APIGatewayProxyResponse, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return FailResponse(&ErrorWrapper{
			Error:      err.Error(),
			Message:    "Gor Error When Marshaling",
			StatusCode: 400,
		})
	}

	return events.APIGatewayProxyResponse{
		StatusCode:      data.StatusCode,
		IsBase64Encoded: false,
		Body:            string(body),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func FailResponse(data *ErrorWrapper) (events.APIGatewayProxyResponse, error) {
	body, _ := json.Marshal(data)

	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: data.StatusCode,
	}, nil
}

func SuccessResponsePagination(data *SuccessPaginationWrapper) (events.APIGatewayProxyResponse, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return FailResponse(&ErrorWrapper{
			Error:      err.Error(),
			Message:    "Gor Error When Marshaling",
			StatusCode: 400,
		})
	}

	return events.APIGatewayProxyResponse{
		StatusCode:      data.StatusCode,
		IsBase64Encoded: false,
		Body:            string(body),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}
