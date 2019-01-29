package dynamock

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type ExpectedPutItemRequest struct {
	requestExpectation
	complete bool
	err      error
	table    string
	items map[string]dynamodb.AttributeValue
	conditionExpression *string
	output   dynamodb.GetItemOutput
}

func (e *ExpectedPutItemRequest) WithItems(items map[string]dynamodb.AttributeValue) *ExpectedPutItemRequest {
	e.items = items
	return e
}

func (e *ExpectedPutItemRequest) WithConditionExpression(expression *string) *ExpectedPutItemRequest {
	e.conditionExpression = expression
	return e
}
 
func (e *ExpectedPutItemRequest) WithOutput(output dynamodb.GetItemOutput) *ExpectedPutItemRequest {
	e.output = output
	return e
}

func (e *dynamo) PutItemRequest(input *dynamodb.PutItemInput) dynamodb.PutItemRequest {
	mock := dynamodb.PutItemRequest{
		&aws.Request{
			Data:  nil,
			Error: nil,
		},
		Input: input,
	}

	for index, request := range e.mock.requests {
		if err := comparePutItemRequests(request, input); err == nil {
			e.mock.requests[index].complete = true
			mock.Request.Data = request.output
			return mock
		}

		if !e.mock.ordered {
			continue
		} else {
			e.mock.errors = append(e.mock.errors, fmt.Errorf("Expected %+v\n but found PutItemRequest %+v", request, input))
			return mock
		}
	}

	e.mock.errors = append(e.mock.errors, fmt.Errorf("Unexpected PutItemRequest %+v\n", input))
	return mock
}

func comparePutItemRequests(expected ExpectedPutItemRequest, actual dynamodb.PutItemInput) error {
	if putItemRequest, ok := request.(*ExpectedPutItemRequest); ok {
		if putItemRequest.table != actual.TableName {
			return fmt.Errorf("Expected PutItemRequest table %s but found %s", putItemRequest.table, actual.TableName)
		}

		if putItemRequest.conditionExpression != nil {
			if *putItemRequest.conditionExpression != *actual.ConditionExpression {
				return fmt.Errorf("Expected PutItemRequest ConditionExpression %s but found %s", *putitemRequest.conditionExpression, *actual.ConditionExpression)
			}
		}

		if putItemRequest.items != nil {
			if err, ok := DeepEqualAttributeValues(putItemRequest.items, actual.Items); !ok {
				return err
			}
		}

		return nil
	}

	return fmt.Errorf("Expected PutItemRequest \"%+v\"\nBut Found \"%+v\"", expected, actual)
}