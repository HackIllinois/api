package user

import (
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/user/config"
	"github.com/HackIllinois/api/services/user/controller"
	"github.com/HackIllinois/api/services/user/service"
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
	controller.SetupController(router.PathPrefix("/user"))

	log.Fatal(apiserver.StartServer(config.USER_PORT, router, "user", Initialize))
}
