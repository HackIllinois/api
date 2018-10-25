package service

import (
	"github.com/HackIllinois/api/services/notifications/config"
    "github.com/HackIllinois/api/services/notifications/models"
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
func GetAllTopics() (*models.TopicList, error) {
	out, err := client.ListTopics(&sns.ListTopicsInput{})

	if err != nil {
		return nil, err
	}

	var topic_list models.TopicList
	for _, topic := range out.Topics {
		topic_list.Topics = append(topic_list.Topics, *topic.TopicArn)
	}

	return &topic_list, nil
}
