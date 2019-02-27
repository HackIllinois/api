package rsvp

import (
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/rsvp/config"
	"github.com/HackIllinois/api/services/rsvp/controller"
	"github.com/HackIllinois/api/services/rsvp/service"
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
	controller.SetupController(router.PathPrefix("/rsvp"))

	log.Fatal(apiserver.StartServer(config.RSVP_PORT, router, "rsvp", Initialize))
}
