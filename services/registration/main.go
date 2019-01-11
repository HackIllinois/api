package registration

import (
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/registration/config"
	"github.com/HackIllinois/api/services/registration/controller"
	"github.com/gorilla/mux"
	"log"
)

func Entry() {
	router := mux.NewRouter()
	controller.SetupController(router.PathPrefix("/registration"))

	log.Fatal(apiserver.StartServer(config.REGISTRATION_PORT, router, "registration"))
}
