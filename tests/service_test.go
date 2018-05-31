package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/HackIllinois/api-auth/models"
	"github.com/HackIllinois/api-auth/service"
)

/*
	Tests that GetUserInfo returns the correct user information
*/
func TestGetUserInfo(t *testing.T) {
	actualUserInfo := models.UserInfo{
		ID:       "1984",
		Username: "jane_smith",
		Email:    "jane.smith@gmail.com",
	}
	fmt.Println(actualUserInfo)
	body := bytes.Buffer{}
	json.NewEncoder(&body).Encode(&actualUserInfo)

	resp, err := http.Post("http://localhost:8000/user/", "application/json", &body)
	fmt.Println(resp)

	if err != nil {
		t.Fatal(err)
	}

	fetchedUserInfo, err := service.GetUserInfo(actualUserInfo.ID)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(actualUserInfo, fetchedUserInfo) {
		t.Errorf("Wrong user info. Expected %v, got %v", actualUserInfo, fetchedUserInfo)
	}
}
