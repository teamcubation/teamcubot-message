package message

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type sqsQueue struct {
	svc      *sqs.SQS
	queueURL *string
}

func NewSQSQueue(sess *session.Session, queueURL string) Queue {
	return &sqsQueue{svc: sqs.New(sess), queueURL: &queueURL}
}

func (q *sqsQueue) Consume(ctx context.Context) ([]Message, error) {
	result, err := q.svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            q.queueURL,
		MaxNumberOfMessages: aws.Int64(10),
		WaitTimeSeconds:     aws.Int64(20),
	})

	if err != nil {
		log.Println("error receiving messages:", err)
		return nil, err
	}

	var messages []Message

	for _, message := range result.Messages {
		var myMessage Message
		err := json.Unmarshal([]byte(*message.Body), &myMessage)
		if err != nil {
			log.Println("error decoding message:", err)
			continue
		}

		myMessage.ID = *message.ReceiptHandle

		log.Println("received message:", myMessage.ID)

		messages = append(messages, myMessage)
	}

	return messages, nil
}

func (q *sqsQueue) DeleteMessage(ctx context.Context, id string) error {
	_, err := q.svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      q.queueURL,
		ReceiptHandle: aws.String(id),
	})

	if err != nil {
		log.Println("error deleting message:", err)
		return fmt.Errorf("error deleting message: %w", err)
	}

	return nil
}
