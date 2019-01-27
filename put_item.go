package dynamock

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type ExpectedPutItemRequest struct {
	requestExpectation
	table string
	key    map[string]dynamodb.AttributeValue
	output dynamodb.GetItemOutput
}

func (e *ExpectedPutItemRequest) WithKeys(key map[string]dynamodb.AttributeValue) *ExpectedPutItemRequest {
	e.key = key
	return e
}

func (e *ExpectedPutItemRequest) WithOutput(output dynamodb.GetItemOutput) *ExpectedPutItemRequest {
	e.output = output
	return e
}

func (d *dynamo) PutItemRequet() {
	
}