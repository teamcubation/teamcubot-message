package bootstrap

import (
	"context"
	"log"

	"github.com/teamcubation/teamcubot-message/internal/platform/aws"
	"github.com/teamcubation/teamcubot-message/internal/platform/config"
	"github.com/teamcubation/teamcubot-message/internal/watcher"
	"github.com/teamcubation/teamcubot-message/internal/watcher/message"
)

func Run() error {
	err := config.LoadConfigs()
	if err != nil {
		return err
	}

	session, err := aws.NewAWSSession()
	if err != nil {
		return err
	}

	messageSQSAdapter := message.NewSQSQueue(session, config.GetQueueURL())
	messageRepoAdapter := message.NewDynamoRepository(session, config.GetTableName())

	messageUsecase := watcher.NewMessageUsecase(messageRepoAdapter, messageSQSAdapter)

	go messageUsecase.ReadMessages(context.Background())

	log.Println("Running...")

	select {}
}
