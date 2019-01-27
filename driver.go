package dynamock

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbiface"
)

type dynamo struct {
	dynamodbiface.DynamoDBAPI
	mock *dynamock
}

func New() (dynamodbiface.DynamoDBAPI, *dynamock) {
	db = new(dynamo)
	db.mock = new(dynamock)
	return db, db.mock
}
