package tests

import (
	"errors"
	"github.com/HackIllinois/api/services/print/models"
	"github.com/HackIllinois/api/services/print/service"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"reflect"
	"testing"
)

func TestPrintValidUser(t *testing.T) {
	service.GetUserInfo = func(id string) (*models.UserInfo, error) {
		return &models.UserInfo{
			ID:       "testid",
			Username: "testusername",
			Email:    "testemail@domain.com",
		}, nil
	}

	print_resp, _ := service.PublishPrintJob(&models.PrintJob{ID: "1", Location: models.DCL})
	expected_resp := &sns.PublishOutput{MessageId: aws.String("printjob-uuid")}
	if !reflect.DeepEqual(print_resp, expected_resp) {
		t.Errorf("Wrong sns response recieved Expected %v, got %v", print_resp, expected_resp)
	}
}

func TestPrintInvalidUser(t *testing.T) {
	service.GetUserInfo = func(id string) (*models.UserInfo, error) {
		return nil, errors.New("User service failed to return information")
	}
	_, err := service.PublishPrintJob(&models.PrintJob{ID: "1", Location: models.DCL})
	if err == nil {
		t.Errorf("Expected print job publish to fail with invalid user id")
	}
}
