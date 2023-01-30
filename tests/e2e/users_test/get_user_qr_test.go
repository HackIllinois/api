package tests

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/common/utils"
	"github.com/HackIllinois/api/services/user/models"
)

// This test + TestStaffCheckinExpiredToken verify that an expiring token will be generated
// and that an expired token will not checkin, respectively.
func TestGetUserQR(t *testing.T) {
	CreateUserInfo()
	defer ClearUserInfo()

	received_qr_info_container := models.QrInfoContainer{}
	received_qr_info_error := errors.ApiError{}

	response, err := user_client.New().Get(fmt.Sprintf("/user/qr/")).Receive(&received_qr_info_container, &received_qr_info_error)

	if err != nil {
		t.Fatalf("Failed to parse qr info container: %v", err)
		return
	}

	if response.StatusCode != http.StatusOK {
		t.Fatalf("QR code request failed: %v, %v", response.Status, received_qr_info_error)
		return
	}

	u, err := url.Parse(received_qr_info_container.QrInfo)

	if err != nil {
		t.Fatalf("Failed to parse url: %v", err)
		return
	}

	query_map, err := url.ParseQuery(u.RawQuery)

	if err != nil {
		t.Fatalf("Failed to parse query string: %v", err)
		return
	}

	user_token := query_map.Get("userToken")

	res, err := utils.ExtractFieldFromJWT(string(TOKEN_SECRET), user_token, "userId")

	if err != nil {
		t.Fatalf("Failed to get userId field from jwt: %v", err)
		return
	}

	userId := res[0]

	if userId != TEST_USER_ID {
		t.Fatalf("Expected user id %v, got %v", TEST_USER_ID, userId)
		return
	}

	res, err = utils.ExtractFieldFromJWT(string(TOKEN_SECRET), user_token, "exp")

	if err != nil {
		t.Fatalf("Failed to get exp field from jwt: %v", err)
		return
	}

	exp_float, err := strconv.ParseFloat(res[0], 64)

	if err != nil {
		t.Fatalf("Failed to parse exp field from jwt: %v", err)
		return
	}

	exp := int64(exp_float)

	exp_time := time.Unix(exp, 0)
	should_be_expired := time.Now().Add(time.Minute)

	now := time.Now()
	expires_properly := now.Before(exp_time) && should_be_expired.After(exp_time)

	if !expires_properly {
		t.Fatalf("Token does not expire properly! Got exp field of %v, should expire between %v and %v", exp, now.Unix(), should_be_expired.Unix())
		return
	}
}
