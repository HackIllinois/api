package checkin

import (
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/checkin/config"
	"github.com/HackIllinois/api/services/checkin/controller"
	"github.com/HackIllinois/api/services/checkin/service"
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
	controller.SetupController(router.PathPrefix("/checkin"))

	log.Fatal(apiserver.StartServer(config.CHECKIN_PORT, router, "checkin", Initialize))
}
