package rsvp

import (
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/rsvp/config"
	"github.com/HackIllinois/api/services/rsvp/controller"
	"github.com/gorilla/mux"
	"log"
)

func Entry() {
	router := mux.NewRouter()
	controller.SetupController(router.PathPrefix("/rsvp"))

	log.Fatal(apiserver.StartServer(config.RSVP_PORT, router, "rsvp"))
}
