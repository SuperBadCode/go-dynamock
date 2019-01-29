package dynamock

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbiface"
)

type dynamo struct {
	dynamodbiface.DynamoDBAPI
	mock *Dynamock
}

func New() (dynamodbiface.DynamoDBAPI, *Dynamock) {
	db = new(dynamo)
	db.mock = new(Dynamock)
	return db, &db.mock
}
