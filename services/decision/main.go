package decision

import (
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/decision/config"
	"github.com/HackIllinois/api/services/decision/controller"
	"github.com/HackIllinois/api/services/decision/service"
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
	controller.SetupController(router.PathPrefix("/decision"))

	log.Fatal(apiserver.StartServer(config.DECISION_PORT, router, "decision", Initialize))
}
