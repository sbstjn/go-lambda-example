package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-lambda-go/events"
)

func TestWithoutParameter(t *testing.T) {
	res, err := handleRequest(events.APIGatewayProxyRequest{})

	assert.NotEmpty(t, res)
	assert.Empty(t, err)
	assert.Equal(t, 404, res.StatusCode)
}

func TestWithParameter(t *testing.T) {
	res, errRequest := handleRequest(events.APIGatewayProxyRequest{
		PathParameters: map[string]string{"id": "d1"},
	})

	var data PersonResponse
	errUnmarshal := json.Unmarshal([]byte(res.Body), &data)

	assert.Equal(t, nil, errRequest)
	assert.Equal(t, nil, errUnmarshal)

	assert.Equal(t, "{\"data\":{\"id\":\"d1\",\"name\":\"Anton\",\"age\":31}}", res.Body)
	assert.Equal(t, "Anton", data.Person.Name)
}
