package dynamock

import (
	"reflect"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func DeepEqualAttributeValues(expected map[string]dynamodb.AttributeValue, actual map[string]dynamodb.AttributeValue) (error, bool) {
	if len(expected) != len(actual) {
		return fmt.Errorf("Expected %+v\n but found %+v", expected, actual), false
	}

	for key, expectedValue := range expected {
		if actualValue, ok := actual[key]; ok {
			if _, ok := expectedValue.(ignoreField); !ok {
				if !reflect.DeepEqual(expectedValue, actualValue) {
					return fmt.Errorf("Expected %+v\n but found %+v", expectedValue, actualValue), false
				}
			}
		} else {
			return fmt.Errorf("Key Value Pair Doesn't Exist"), false
		}
	}

	return nil, true
}