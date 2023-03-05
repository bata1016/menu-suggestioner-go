package dynamodb

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// DynamoDB は、DynamoDBのクライアントです。
type DynamoDB struct {
	client    *dynamodb.DynamoDB
	tableName string
}

// NewDynamoDBは、DynamoDBのクライアントを生成します。
func NewDynamoDB(tableName string) *DynamoDB {
	ddb := dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))
	return &DynamoDB{
		client:    ddb,
		tableName: tableName,
	}
}

// PutItem は、DynamoDBへデータを一件登録します。
func (ddb *DynamoDB) PutItem(params interface{}) (*dynamodb.PutItemOutput, error) {
	input := &dynamodb.PutItemInput{}
	av, err := dynamodbattribute.MarshalMap(params)
	if err != nil || av == nil {
		return nil, errors.New(fmt.Sprintf("dynamodbattribute.MarshalMap got Error. %s", err))
	}

	input = &dynamodb.PutItemInput{
		Item:                   av,
		ReturnConsumedCapacity: aws.String("TOTAL"),
		TableName:              aws.String(ddb.tableName),
	}
	result, err := ddb.client.PutItem(input)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("PutItem got Error. %s", err))
	}
	if result == nil {
		return nil, errors.New(fmt.Sprintf("PutItem got nil."))
	}
	return result, nil
}
