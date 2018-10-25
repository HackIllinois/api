package service

import (
	"fmt"
	"github.com/HackIllinois/api/services/notifications/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

var sess *session.Session
var client *sns.SNS

func init() {
	sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String(config.SNS_REGION),
	}))
	client = sns.New(sess)
}

/*
	Returns a list of available SNS Topics
*/
func GetAllTopics() ([]string, error) {
	out, err := client.ListTopics(&sns.ListTopicsInput{})

	if err != nil {
		return nil, err
	}

	var topic_names []string
	for _, topic := range out.Topics {
		topic_names = append(topic_names, *topic.TopicArn)
	}

	return topic_names, nil
}
