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

/*
	Creates an SNS Topic
*/
func CreateTopic(name string) (*models.TopicArn, error) {
	out, err := client.CreateTopic(&sns.CreateTopicInput{Name: &name})

	if err != nil {
		return nil, err
	}

	topic_arn := models.TopicArn{Arn: *out.TopicArn}

	return &topic_arn, nil
}

/*
	Deletes an SNS Topic
*/
func DeleteTopic(arn string) error {
	_, err := client.DeleteTopic(&sns.DeleteTopicInput{TopicArn: &arn})

	if err != nil {
		return err
	}

	return nil
}

/*
	Dispatches a notification to a given SNS Topic
*/
func PublishNotification(notification models.Notification) (*models.MessageId, error) {
	out, err := client.Publish(&sns.PublishInput{
		TopicArn: &notification.Arn,
		Message:  &notification.Message,
	})

	if err != nil {
		return nil, err
	}

	message_id := models.MessageId{MessageId: *out.MessageId}

	return &message_id, nil
}
