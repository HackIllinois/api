package tests

import (
	"fmt"
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"github.com/HackIllinois/api-auth/models"
	"github.com/HackIllinois/api-auth/service"
)

/*
	Tests that GetUserInfo returns the correct user information
*/
func TestGetUserInfo(t *testing.T) {
	newId := "1984"
	newUsername := "jane_smith"
	newEmail := "jane.smith@gmail.com"
	actualUserInfo = models.UserInfo{
		ID: newId,
		Username: newUsername,
		Email: newEmail,
	}

	postBody := []byte(
		fmt.Sprintf(
			`{
				"id" : "%s",
				"username" : "%s",
				"email" : "%s",
			}`,
			newId, newUsername, newEmail
		)
	)
	resp, err := http.Post("/user/")
	fetchedUserInfo := service.GetUserInfo(newId, bytes.NewBuffer(postBody))

	if err != nil {
		t.Fatal(err)
	}
	
	if !reflect.DeepEqual(actualUserInfo, fetchedUserInfo) {
		t.Errorf("Wrong user info. Expected %v, got %v", actualUserInfo, fetchedUserInfo)
	}
}
