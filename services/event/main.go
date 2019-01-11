package event

import (
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/event/config"
	"github.com/HackIllinois/api/services/event/controller"
	"github.com/gorilla/mux"
	"log"
)

func Entry() {
	router := mux.NewRouter()
	controller.SetupController(router.PathPrefix("/event"))

	log.Fatal(apiserver.StartServer(config.EVENT_PORT, router, "event"))
}
