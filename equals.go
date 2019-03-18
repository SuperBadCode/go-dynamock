package dynamock

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	escapeSequences = strings.NewReplacer("\n", "", "\t", "", "\r", "")
)

func contains(key string, values []string) bool {
	for _, value := range values {
		if value == key {
			return true
		}
	}
	return false
}

func DeepEqualAttributeValues(expected map[string]dynamodb.AttributeValue, actual map[string]dynamodb.AttributeValue, ignore []string) (error, bool) {
	for key, expectedValue := range expected {
		if actualValue, ok := actual[key]; ok {
			if !contains(key, ignore) && !reflect.DeepEqual(expectedValue, actualValue) {
				return fmt.Errorf("Expected %+v\n but found %+v", expectedValue, actualValue), false
			}
		} else {
			return fmt.Errorf("Key Value Pair Doesn't Exist"), false
		}
	}

	return nil, true
}
