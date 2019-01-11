package notifications

import (
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/notifications/config"
	"github.com/HackIllinois/api/services/notifications/controller"
	"github.com/gorilla/mux"
	"log"
)

func Entry() {
	router := mux.NewRouter()
	controller.SetupController(router.PathPrefix("/notifications"))

	log.Fatal(apiserver.StartServer(config.NOTIFICATIONS_PORT, router, "notifications"))
}
