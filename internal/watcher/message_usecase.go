package watcher

import (
	"context"
	"log"

	"github.com/teamcubation/teamcubot-message/internal/watcher/message"
)

type MessageUsecase interface {
	ReadMessages(context.Context)
}

type messageUsecase struct {
	repo  message.Repository
	queue message.Queue
}

func NewMessageUsecase(repo message.Repository, queue message.Queue) MessageUsecase {
	return &messageUsecase{
		repo:  repo,
		queue: queue,
	}
}

func (uc *messageUsecase) ReadMessages(ctx context.Context) {
	for {
		messages, err := uc.queue.Consume(ctx)
		if err != nil {
			log.Println("error consuming queue:", err)
			continue
		}

		for _, msg := range messages {
			err := uc.repo.SaveMessage(ctx, msg)
			if err != nil {
				log.Println("error saving message in DB:", err)
				continue
			}

			err = uc.queue.DeleteMessage(ctx, msg.ID)
			if err != nil {
				log.Println("error deleting message in queue:", err)
			}
		}
	}
}
