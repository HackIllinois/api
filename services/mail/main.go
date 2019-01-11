package mail

import (
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/mail/config"
	"github.com/HackIllinois/api/services/mail/controller"
	"github.com/gorilla/mux"
	"log"
)

func Entry() {
	router := mux.NewRouter()
	controller.SetupController(router.PathPrefix("/mail"))

	log.Fatal(apiserver.StartServer(config.MAIL_PORT, router, "mail"))
}
