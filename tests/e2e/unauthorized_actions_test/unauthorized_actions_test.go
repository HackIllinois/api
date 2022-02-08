// // Attendee user attempts admin actions (delete event, etc.) should result in failure
package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/HackIllinois/api/common/configloader"
	"github.com/HackIllinois/api/services/mail/models"
	"github.com/HackIllinois/api/tests/common"
	"github.com/dghubble/sling"
	"gopkg.in/mgo.v2"
)

var attendee_client *sling.Sling
var session *mgo.Session

func TestMain(m *testing.M) {

	cfg, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	attendee_client = common.GetSlingClient("Attendee")

	session = common.GetLocalMongoSession()

	decision_db_name, err := cfg.Get("MAIL_DB_NAME")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
	session.DB(decision_db_name).DropDatabase()

	return_code := m.Run()
	os.Exit(return_code)
}

func TestAttendeeUnauthorizedCalls(t *testing.T) {
	received_mail_list := models.MailList{}
	response, _ := attendee_client.New().Get("/mail/list/").Receive(&received_mail_list, &received_mail_list)

	if response.StatusCode != 403 {
		t.Errorf("Attendee able to access admin-only endpoints")
	}
}
