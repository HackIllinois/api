package decision

import (
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/decision/config"
	"github.com/HackIllinois/api/services/decision/controller"
	"github.com/gorilla/mux"
	"log"
)

func Entry() {
	router := mux.NewRouter()
	controller.SetupController(router.PathPrefix("/decision"))

	log.Fatal(apiserver.StartServer(config.DECISION_PORT, router, "decision"))
}
