package dynamock

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type ExpectedGetItemRequest struct {
	request              expectation
	err                  error
	table                string
	keys                 map[string]dynamodb.AttributeValue
	projectionExpression *string
	ignoreFields         []string
	output               dynamodb.GetItemOutput
}

func (e *ExpectedGetItemRequest) WithKeys(keys map[string]dynamodb.AttributeValue) *ExpectedGetItemRequest {
	e.keys = keys
	return e
}

func (e *ExpectedGetItemRequest) WithProjectionExpression(expression *string) *ExpectedGetItemRequest {
	e.projectionExpression = expression
	return e
}

func (e *ExpectedGetItemRequest) WithOutput(output dynamodb.GetItemOutput) *ExpectedGetItemRequest {
	e.output = output
	return e
}

func (e *dynamo) GetItemRequest(input *dynamodb.GetItemInput) dynamodb.GetItemRequest {
	mock := dynamodb.GetItemRequest{
		Input: input,
	}

	output := e.matchExpectation(input, compareGetItemRequests)
	if output != nil {
		mock.Data = output
	}

	return mock
}

func compareGetItemRequests(expected expectation, input interface{}) (error, interface{}) {
	actual := input.(dynamodb.GetItemInput)
	if getItemRequest, ok := expected.(ExpectedGetItemRequest); ok {
		if getItemRequest.table != *actual.TableName {
			return fmt.Errorf("expected GetItemRequest table %s but found %s", getItemRequest.table, actual.TableName), nil
		}

		if getItemRequest.projectionExpression != nil {
			if escapeSequences.Replace(*getItemRequest.projectionExpression) != escapeSequences.Replace(*actual.ProjectionExpression) {
				return fmt.Errorf("expected GetItemRequest ConditionExpression %s but found %s", *getItemRequest.projectionExpression, *actual.ProjectionExpression), nil
			}
		}

		if getItemRequest.keys != nil {
			if err, ok := DeepEqualAttributeValues(getItemRequest.keys, actual.Key, getItemRequest.ignoreFields); !ok {
				return err, nil
			}
		}

		return nil, &getItemRequest.output
	}

	return fmt.Errorf("expected GetItemRequest \"%+v\"\nBut Found \"%+v\"", expected, actual), nil
}
