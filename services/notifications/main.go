package notifications

import (
	"fmt"
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/notifications/config"
	"github.com/HackIllinois/api/services/notifications/controller"
	"github.com/HackIllinois/api/services/notifications/service"
	"github.com/gorilla/mux"
	"log"
	"os"
)

func Entry() {
	err := config.Initialize()

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)

	}

	err = service.Initialize()

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	router := mux.NewRouter()
	controller.SetupController(router.PathPrefix("/notifications"))

	log.Fatal(apiserver.StartServer(config.NOTIFICATIONS_PORT, router, "notifications"))
}
