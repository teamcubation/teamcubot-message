package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/teamcubation/teamcubot-message/internal/platform/config"
)

func NewAWSSession() (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.GetAWSRegion()),
	})

	if err != nil {
		return nil, err
	}

	return sess, nil
}
