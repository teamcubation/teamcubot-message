package message

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type messageDAO struct {
	ID     string `json:"id"`
	Status uint   `json:"status"`
	Value  string `json:"value"`
}

type dynamoRepository struct {
	client    *dynamodb.DynamoDB
	tableName string
}

func NewDynamoRepository(sess *session.Session, tableName string) Repository {
	return &dynamoRepository{client: dynamodb.New(sess), tableName: tableName}
}

func (r *dynamoRepository) SaveMessage(ctx context.Context, message Message) error {
	av, err := dynamodbattribute.MarshalMap(messageDAO(message))
	if err != nil {
		return fmt.Errorf("error marshalling item: %w", err)
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item:      av,
	}

	_, err = r.client.PutItem(input)
	if err != nil {
		return fmt.Errorf("error saving item in DB: %w", err)
	}
	return nil
}

func (r *dynamoRepository) GetMessage(field, value string, resultado interface{}) (*Message, error) {
	key := map[string]*dynamodb.AttributeValue{
		field: {
			S: aws.String(value),
		},
	}

	input := &dynamodb.GetItemInput{
		TableName: aws.String(r.tableName),
		Key:       key,
	}

	result, err := r.client.GetItem(input)
	if err != nil {
		return nil, fmt.Errorf("error getting item from DB: %w", err)
	}

	message := new(messageDAO)
	if err := dynamodbattribute.UnmarshalMap(result.Item, message); err != nil {
		return nil, fmt.Errorf("error unmarshalling item: %w", err)
	}

	return (*Message)(message), nil
}
