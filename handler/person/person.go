package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sbstjn/go-lambda-example/repository"
)

// PersonResponse
type PersonResponse struct {
	Person repository.Person `json:"data"`
}

func handleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	person, err := repository.GetPerson(req.PathParameters["id"])

	if err != nil {
		return events.APIGatewayProxyResponse{Body: "User not found", StatusCode: 404}, nil
	}

	body, err := json.Marshal(PersonResponse{*person})

	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Unable to marshal JSON", StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil
}

func main() {
	lambda.Start(handleRequest)
}
