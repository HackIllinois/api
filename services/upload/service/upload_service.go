package service

import (
	"github.com/HackIllinois/api/services/upload/config"
	"github.com/HackIllinois/api/services/upload/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"time"
)

var sess *session.Session
var client *s3.S3

func init() {
	sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String(config.S3_REGION),
	}))
	client = s3.New(sess)
}

/*
	Returns a presigned link to user requested user's resume
*/
func GetUserResumeLink(id string) (*models.UserResume, error) {
	var signed_url string
	var err error

	if config.IS_PRODUCTION {
		request, _ := client.GetObjectRequest(&s3.GetObjectInput{
			Bucket: aws.String(config.S3_BUCKET),
			Key:    aws.String("resumes/" + id + ".pdf"),
		})

		signed_url, err = request.Presign(15 * time.Minute)

		if err != nil {
			return nil, err
		}
	} else {
		signed_url = "/tmp/uploads/" + id + ".pdf"
	}

	resume := models.UserResume{
		ID:     id,
		Resume: signed_url,
	}

	return &resume, nil
}

/*
	Update the given user's resume
*/
func GetUpdateUserResumeLink(id string) (*models.UserResume, error) {
	var signed_url string
	var err error

	if config.IS_PRODUCTION {
		request, _ := client.PutObjectRequest(&s3.PutObjectInput{
			Bucket: aws.String(config.S3_BUCKET),
			Key:    aws.String("resumes/" + id + ".pdf"),
		})

		signed_url, err = request.Presign(15 * time.Minute)

		if err != nil {
			return nil, err
		}
	} else {
		signed_url = "/tmp/uploads/" + id + ".pdf"
	}

	resume := models.UserResume{
		ID:     id,
		Resume: signed_url,
	}

	return &resume, nil
}
