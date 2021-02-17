package profile

import (
	"log"

	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/profile/config"
	"github.com/HackIllinois/api/services/profile/controller"
	"github.com/HackIllinois/api/services/profile/service"
	"github.com/gorilla/mux"
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
	controller.SetupController(router.PathPrefix("/profile"))

	log.Fatal(apiserver.StartServer(config.PROFILE_PORT, router, "profile", Initialize))
}
