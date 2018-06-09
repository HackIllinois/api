package service

import (
	"errors"
	"bytes"
	"time"
	"net/http"
	"github.com/HackIllinois/api-upload/config"
	"github.com/HackIllinois/api-upload/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var sess *session.Session
var uploader *s3manager.Uploader
var client *s3.S3

func init() {
	sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String(config.S3_REGION),
	}))
	uploader = s3manager.NewUploader(sess)
	client = s3.New(sess)
}

/*
	Returns a presigned link to user requested user's resume
*/
func GetUserResumeLink(id string) (*models.UserResume, error) {
	request, _ := client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(config.S3_BUCKET),
		Key: aws.String("resumes/" + id + ".pdf"),
	})

	signed_url, err := request.Presign(15 * time.Minute)

	if err != nil {
		return nil, err
	}

	resume  := models.UserResume{
		ID: id,
		Resume: signed_url,
	}

	return &resume, nil
}

/*
	Update the given user's resume
*/
func UpdateUserResume(id string, file_buffer []byte) error {
	content_type := http.DetectContentType(file_buffer)

	if(content_type != "application/pdf") {
		return errors.New("Resume upload must be a pdf")
	}

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(config.S3_BUCKET),
		Key: aws.String("resumes/" + id + ".pdf"),
		Body: bytes.NewReader(file_buffer),
	})

	return err
}
