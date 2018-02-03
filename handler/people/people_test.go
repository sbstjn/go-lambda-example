package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-lambda-go/events"
)

func TestWithoutData(t *testing.T) {
	res, errRequest := handleRequest(events.APIGatewayProxyRequest{})

	var data PeopleResponse
	errUnmarshal := json.Unmarshal([]byte(res.Body), &data)

	assert.Equal(t, nil, errRequest)
	assert.Equal(t, nil, errUnmarshal)

	assert.Equal(t, "{\"data\":[{\"id\":\"d1\",\"name\":\"Anton\",\"age\":31},{\"id\":\"c2\",\"name\":\"Frank\",\"age\":28},{\"id\":\"b1\",\"name\":\"Horst\",\"age\":42},{\"id\":\"e3\",\"name\":\"Steve\",\"age\":34}]}", res.Body)
	assert.Equal(t, 4, len(data.People))
	assert.Equal(t, "Anton", data.People[0].Name)
}
