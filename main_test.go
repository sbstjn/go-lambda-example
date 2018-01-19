package main

import (
	"errors"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestWithData(t *testing.T) {
	res, err := handler(events.APIGatewayProxyRequest{
		Body: "Example 🎉",
	})

	assert.Equal(t, nil, err)
	assert.Equal(t, "You sent some data 😍\nExample 🎉", res.Body)
}

func TestWithoutData(t *testing.T) {
	res, err := handler(events.APIGatewayProxyRequest{})

	assert.Equal(t, errors.New("No data received"), err)
	assert.Equal(t, "", res.Body)
}
