package dynamock

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type ExpectedPutItemRequest struct {
	expectation
	err                 error
	table               string
	items               map[string]dynamodb.AttributeValue
	conditionExpression *string
	ignoreFields        []string
	output              dynamodb.PutItemOutput
}

func (e *ExpectedPutItemRequest) WithItems(items map[string]dynamodb.AttributeValue) *ExpectedPutItemRequest {
	e.items = items
	return e
}

func (e *ExpectedPutItemRequest) WithConditionExpression(expression *string) *ExpectedPutItemRequest {
	e.conditionExpression = expression
	return e
}

func (e *ExpectedPutItemRequest) WithOutput(output dynamodb.PutItemOutput) *ExpectedPutItemRequest {
	e.output = output
	return e
}

func (e *dynamo) PutItemRequest(input *dynamodb.PutItemInput) dynamodb.PutItemRequest {
	mock := dynamodb.PutItemRequest{
		Input: input,
	}

	output := e.matchExpectation(*input, comparePutItemRequests)
	if output != nil {
		mock.Data = output
	}

	return mock
}

func comparePutItemRequests(expected expectation, input interface{}) (error, interface{}) {
	actual := input.(dynamodb.PutItemInput)
	if putItemRequest, ok := expected.(ExpectedPutItemRequest); ok {
		if putItemRequest.table != *actual.TableName {
			return fmt.Errorf("expected PutItemRequest table %s but found %s", putItemRequest.table, actual.TableName), nil
		}

		if putItemRequest.conditionExpression != nil {
			if escapeSequences.Replace(*putItemRequest.conditionExpression) != escapeSequences.Replace(*actual.ConditionExpression) {
				return fmt.Errorf("expected PutItemRequest ConditionExpression %s but found %s", *putItemRequest.conditionExpression, *actual.ConditionExpression), nil
			}
		}

		if putItemRequest.items != nil {
			if err, ok := DeepEqualAttributeValues(putItemRequest.items, actual.Item, putItemRequest.ignoreFields); !ok {
				return err, nil
			}
		}

		return nil, putItemRequest.output
	}

	return fmt.Errorf("expected PutItemRequest \"%+v\"\nBut Found \"%+v\"", expected, actual), nil
}
