package message

import "context"

type Message struct {
	ID     string
	Status uint
	Value  string
}

type Repository interface {
	SaveMessage(ctx context.Context, message Message) error
	GetMessage(field, value string, resultado interface{}) (*Message, error)
}

type Queue interface {
	Consume(ctx context.Context) ([]Message, error)
	DeleteMessage(ctx context.Context, id string) error
}
