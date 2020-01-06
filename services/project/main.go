package project

import (
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/project/config"
	"github.com/HackIllinois/api/services/project/controller"
	"github.com/HackIllinois/api/services/project/service"
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
	controller.SetupController(router.PathPrefix("/project"))

	log.Fatal(apiserver.StartServer(config.PROJECT_PORT, router, "project", Initialize))
}
