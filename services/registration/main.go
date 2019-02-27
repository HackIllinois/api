package registration

import (
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/registration/config"
	"github.com/HackIllinois/api/services/registration/controller"
	"github.com/HackIllinois/api/services/registration/service"
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
	controller.SetupController(router.PathPrefix("/registration"))

	log.Fatal(apiserver.StartServer(config.REGISTRATION_PORT, router, "registration", Initialize))
}
