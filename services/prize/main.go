package prize

import (
	"log"

	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/prize/config"
	"github.com/HackIllinois/api/services/prize/controller"
	"github.com/HackIllinois/api/services/prize/service"
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
	controller.SetupController(router.PathPrefix("/prize"))

	log.Fatal(apiserver.StartServer(config.PRIZE_PORT, router, "prize", Initialize))
}
