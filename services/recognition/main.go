package recognition

import (
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/recognition/config"
	"github.com/HackIllinois/api/services/recognition/controller"
	"github.com/HackIllinois/api/services/recognition/service"
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
	controller.SetupController(router.PathPrefix("/recognition"))

	log.Fatal(apiserver.StartServer(config.RECOGNITION_PORT, router, "recognition", Initialize))
}
