package dynamock

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type ExpectedScanRequest struct {
	request              expectation
	err                  error
	table                string
	limit                *int64
	expressionValues     map[string]dynamodb.AttributeValue
	filterExpression     *string
	projectionExpression *string
	ignoreFields         []string
	output               dynamodb.ScanOutput
}

func (e *ExpectedScanRequest) WithExpressionValues(expressions map[string]dynamodb.AttributeValue) *ExpectedScanRequest {
	e.expressionValues = expressions
	return e
}

func (e *ExpectedScanRequest) WithProjectionExpression(expression *string) *ExpectedScanRequest {
	e.projectionExpression = expression
	return e
}

func (e *ExpectedScanRequest) WithFilterExpression(expression *string) *ExpectedScanRequest {
	e.filterExpression = expression
	return e
}

func (e *ExpectedScanRequest) WithOutput(output dynamodb.ScanOutput) *ExpectedScanRequest {
	e.output = output
	return e
}

func (e *ExpectedScanRequest) WithLimit(limit *int64) *ExpectedScanRequest {
	e.limit = limit
	return e
}

func (e *dynamo) ScanRequest(input *dynamodb.ScanInput) dynamodb.ScanRequest {
	mock := dynamodb.ScanRequest{
		Input: input,
	}

	return mock
}

func compareScanRequests(expected expectation, input interface{}) (error, interface{}) {
	actual := input.(dynamodb.ScanInput)
	if scanRequest, ok := expected.(ExpectedScanRequest); ok {
		if scanRequest.table != *actual.TableName {
			return fmt.Errorf("expected scanRequest table %s but found %s", scanRequest.table, actual.TableName), nil
		}

		if scanRequest.filterExpression != nil {
			if escapeSequences.Replace(*scanRequest.filterExpression) != escapeSequences.Replace(*actual.FilterExpression) {
				return fmt.Errorf("expected scanRequest FilterExpression %s but found %s", *scanRequest.filterExpression, *actual.FilterExpression), nil
			}
		}

		if scanRequest.projectionExpression != nil {
			if escapeSequences.Replace(*scanRequest.projectionExpression) != escapeSequences.Replace(*actual.ProjectionExpression) {
				return fmt.Errorf("expected scanRequest ProjectionExpression %s but found %s", *scanRequest.projectionExpression, *actual.ProjectionExpression), nil
			}
		}

		if scanRequest.expressionValues != nil {
			if err, ok := DeepEqualAttributeValues(scanRequest.expressionValues, actual.ExpressionAttributeValues, scanRequest.ignoreFields); !ok {
				return err, nil
			}
		}

		return nil, &scanRequest.output
	}

	return fmt.Errorf("expected scanRequest \"%+v\"\nBut Found \"%+v\"", expected, actual), nil
}
