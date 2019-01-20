package event

import (
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/event/config"
	"github.com/HackIllinois/api/services/event/controller"
	"github.com/HackIllinois/api/services/event/service"
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
	controller.SetupController(router.PathPrefix("/event"))

	log.Fatal(apiserver.StartServer(config.EVENT_PORT, router, "event", Initialize))
}
