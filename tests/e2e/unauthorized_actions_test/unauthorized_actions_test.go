// // Attendee user attempts admin actions (delete event, etc.) should result in failure
package tests

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/HackIllinois/api/common/configloader"
	event_models "github.com/HackIllinois/api/services/event/models"
	mail_models "github.com/HackIllinois/api/services/mail/models"
	notification_models "github.com/HackIllinois/api/services/notifications/models"
	"github.com/HackIllinois/api/tests/common"
	"github.com/dghubble/sling"
	"go.mongodb.org/mongo-driver/mongo"
)

var attendee_client *sling.Sling
var unauthenticated_client *sling.Sling
var client *mongo.Client

func TestMain(m *testing.M) {

	cfg, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	attendee_client = common.GetSlingClient("Attendee")
	unauthenticated_client = sling.New().Base("http://localhost:8000").Client(nil).Add("Authorization", "FAKE_TOKEN")

	client = common.GetLocalMongoSession()

	mail_db_name, err := cfg.Get("MAIL_DB_NAME")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
	client.Database(mail_db_name).Drop(context.Background())

	notification_db_name, err := cfg.Get("NOTIFICATIONS_DB_NAME")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
	client.Database(notification_db_name).Drop(context.Background())

	event_db_name, err := cfg.Get("EVENT_DB_NAME")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
	client.Database(event_db_name).Drop(context.Background())

	return_code := m.Run()
	os.Exit(return_code)
}

func TestAttendeeUnauthorizedCalls(t *testing.T) {
	received_mail_list := mail_models.MailList{}
	response, _ := attendee_client.New().Get("/mail/list/").Receive(&received_mail_list, &received_mail_list)

	if response.StatusCode != 403 {
		t.Errorf("Attendee able to access admin-only endpoints")
	}
}

func TestUnauthenticatedCalls(t *testing.T) {
	// 1. Staff endpoint
	received_mail_list := mail_models.MailList{}
	response, _ := unauthenticated_client.New().Get("/mail/list/").Receive(&received_mail_list, &received_mail_list)

	if response.StatusCode != 403 {
		t.Errorf("Unauthenticated attendee able to access endpoint that requires authentication")
	}

	// 2. Attendee endpoint
	received_notifications_list := notification_models.NotificationList{}
	response, _ = unauthenticated_client.New().Get("/notifications/topic/all/").Receive(&received_notifications_list, &received_notifications_list)

	if response.StatusCode != 403 {
		t.Errorf("Unauthenticated attendee able to access endpoint that requires authentication")
	}

	// 3. Public endpoint
	received_events_list := event_models.EventList[event_models.EventPublic]{}
	response, _ = unauthenticated_client.New().Get("/event/filter/").Receive(&received_events_list, &received_events_list)

	if response.StatusCode != 200 {
		t.Errorf("Unauthenticated attendee can not access public endpoint.")
	}
}
