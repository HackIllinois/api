package auth

import (
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/auth/config"
	"github.com/HackIllinois/api/services/auth/controller"
	"github.com/HackIllinois/api/services/auth/service"
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
	controller.SetupController(router.PathPrefix("/auth"))

	log.Fatal(apiserver.StartServer(config.AUTH_PORT, router, "auth", Initialize))
}
