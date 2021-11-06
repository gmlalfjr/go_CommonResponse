package utils

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func SuccessResponse(data interface{}, status int) (events.APIGatewayProxyResponse, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return FailResponse(err, 400)
	}

	return events.APIGatewayProxyResponse{
		StatusCode:      status | 200,
		IsBase64Encoded: false,
		Body:            string(body),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func FailResponse(data interface{}, status int) (events.APIGatewayProxyResponse, error) {
	body, _ := json.Marshal(data)

	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: status,
	}, nil
}
