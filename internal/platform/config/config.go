package config

import (
	"os"

	"github.com/joho/godotenv"
)

var configs struct {
	queueURL  string
	awsRegion string
	tableName string
}

func LoadConfigs() error {
	// TODO: validar para variables en produccion.
	return loadENVConfigs()
}

func loadENVConfigs() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	configs.queueURL = os.Getenv("QUEUE_URL")
	configs.awsRegion = os.Getenv("AWS_REGION")
	configs.tableName = os.Getenv("TABLE_NAME")

	return nil
}

func GetQueueURL() string {
	return configs.queueURL
}

func GetTableName() string {
	return configs.tableName
}

func GetAWSRegion() string {
	return configs.awsRegion
}
