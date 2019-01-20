package notifications

import (
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/notifications/config"
	"github.com/HackIllinois/api/services/notifications/controller"
	"github.com/HackIllinois/api/services/notifications/service"
	"github.com/gorilla/mux"
	"log"
)

func Initialize() error {
	err := config.Initialize()

	if err != nil {
		return err

	}

	err = service.Initialize()

	if err != nil {
		return err
	}

	return nil
}

func Entry() {
	err := Initialize()

	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	controller.SetupController(router.PathPrefix("/notifications"))

	log.Fatal(apiserver.StartServer(config.NOTIFICATIONS_PORT, router, "notifications", Initialize))
}
