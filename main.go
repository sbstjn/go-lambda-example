package main

import (
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if len(req.Body) < 1 {
		return events.APIGatewayProxyResponse{}, errors.New("No data received")
	}

	return events.APIGatewayProxyResponse{
		Body:       "You sent some data 😍\n" + req.Body,
		StatusCode: 200,
	}, nil

}

func main() {
	lambda.Start(handler)
}
