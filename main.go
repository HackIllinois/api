package main

import (
	"flag"
	"os"
	"fmt"

	"github.com/HackIllinois/api/gateway"
	"github.com/HackIllinois/api/services/auth"
	"github.com/HackIllinois/api/services/user"
	"github.com/HackIllinois/api/services/registration"
	"github.com/HackIllinois/api/services/decision"
	"github.com/HackIllinois/api/services/rsvp"
	"github.com/HackIllinois/api/services/checkin"
	"github.com/HackIllinois/api/services/upload"
	"github.com/HackIllinois/api/services/mail"
	"github.com/HackIllinois/api/services/event"
	"github.com/HackIllinois/api/services/stat"
	"github.com/HackIllinois/api/services/notifications"
)

var SERVICE_ENTRYPOINTS = map[string](func()) {
	"gateway": gateway.Entry,
	"auth": auth.Entry,
	"user": user.Entry,
	"registration": registration.Entry,
	"decision": decision.Entry,
	"rsvp": rsvp.Entry,
	"checkin": checkin.Entry,
	"upload": upload.Entry,
	"mail": mail.Entry,
	"event": event.Entry,
	"stat": stat.Entry,
	"notifications": notifications.Entry,
}

func main() {
	var service string
	flag.StringVar(&service, "service", "", "The service to start")

	flag.Parse()

	entry, ok := SERVICE_ENTRYPOINTS[service]

	if !ok {
		fmt.Fprintf(os.Stderr, "Could not start service '%s'\n", service)
		os.Exit(1)
	}

	entry()
}
