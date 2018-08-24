package service

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/HackIllinois/api/services/print/config"
	"github.com/HackIllinois/api/services/print/models"
)

var sess *session.Session
var client *sns.SNS
var QR_PREFIX = "hackillinois://qrcode/user?id=%d&identifier=%s"

func init() {
	sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String(config.SNS_REGION),
	}))
	client = sns.New(sess)
}

/*
	Returns the response from the SNS publish request
*/
func PublishPrintJob(job models.PrintJob) (*sns.PublishOutput, error) {
	identifier, err := GetUserInfo(job.ID)
	if err != nil {
		return nil, err
	}
	if !config.IS_PRODUCTION {
		return &sns.PublishOutput { MessageId : aws.String("printjob-uuid") }, nil
	}

	payload := fmt.Sprintf(QR_PREFIX, job.ID, identifier)
	request, resp := client.PublishRequest(&sns.PublishInput {
		MessageStructure: aws.String("json"),
		TopicArn: aws.String(config.PRINT_TOPIC),
		Subject: aws.String(job.Location),
		Message: aws.String(payload),
	})

	err := request.Send()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
