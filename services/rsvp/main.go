package rsvp

import (
	"fmt"
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/rsvp/config"
	"github.com/HackIllinois/api/services/rsvp/controller"
	"github.com/HackIllinois/api/services/rsvp/service"
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
	controller.SetupController(router.PathPrefix("/rsvp"))

	log.Fatal(apiserver.StartServer(config.RSVP_PORT, router, "rsvp"))
}
