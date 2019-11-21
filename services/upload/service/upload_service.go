package service

import (
	"errors"
	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/upload/config"
	"github.com/HackIllinois/api/services/upload/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"time"
)

var db database.Database

var sess *session.Session
var client *s3.S3

func Initialize() error {
	sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String(config.S3_REGION),
	}))
	client = s3.New(sess)

	if db != nil {
		db.Close()
		db = nil
	}

	var err error
	db, err = database.InitDatabase(config.UPLOAD_DB_HOST, config.UPLOAD_DB_NAME)

	if err != nil {
		return err
	}

	return nil
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

/*
	Returns the blob with the given id
*/
func GetBlob(id string) (*models.Blob, error) {
	query := database.QuerySelector{
		"id": id,
	}

	var blob models.Blob
	err := db.FindOne("blobstore", query, &blob)

	if err != nil {
		return nil, err
	}

	return &blob, nil
}

/*
	Creates and stores a blob
*/
func CreateBlob(blob models.Blob) error {
	_, err := GetBlob(blob.ID)

	if err != database.ErrNotFound {
		if err != nil {
			return err
		}
		return errors.New("Blob already exists.")
	}

	err = db.Insert("blobstore", &blob)

	return err
}

/*
	Updates the blob with the given id
*/
func UpdateBlob(blob models.Blob) error {
	selector := database.QuerySelector{
		"id": blob.ID,
	}

	err := db.Update("blobstore", selector, &blob)

	return err
}

/*
Deletes the blob with the given id
Returns the blob that was deleted
*/
func DeleteBlob(id string) (*models.Blob, error) {
	blob, err := GetBlob(id)

	if err != nil {
		return nil, err
	}

	selector := database.QuerySelector{
		"id": id,
	}

	err = db.RemoveOne("blobstore", selector)

	return blob, err
}
