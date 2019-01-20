package stat

import (
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/stat/config"
	"github.com/HackIllinois/api/services/stat/controller"
	"github.com/HackIllinois/api/services/stat/service"
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
	controller.SetupController(router.PathPrefix("/stat"))

	log.Fatal(apiserver.StartServer(config.STAT_PORT, router, "stat", Initialize))
}
