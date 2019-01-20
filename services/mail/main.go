package mail

import (
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/mail/config"
	"github.com/HackIllinois/api/services/mail/controller"
	"github.com/HackIllinois/api/services/mail/service"
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
	controller.SetupController(router.PathPrefix("/mail"))

	log.Fatal(apiserver.StartServer(config.MAIL_PORT, router, "mail", Initialize))
}
