package checkin

import (
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/checkin/config"
	"github.com/HackIllinois/api/services/checkin/controller"
	"github.com/gorilla/mux"
	"log"
)

func Entry() {
	router := mux.NewRouter()
	controller.SetupController(router.PathPrefix("/checkin"))

	log.Fatal(apiserver.StartServer(config.CHECKIN_PORT, router, "checkin"))
}
