package user

import (
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/user/config"
	"github.com/HackIllinois/api/services/user/controller"
	"github.com/gorilla/mux"
	"log"
)

func Entry() {
	router := mux.NewRouter()
	controller.SetupController(router.PathPrefix("/user"))

	log.Fatal(apiserver.StartServer(config.USER_PORT, router, "user"))
}
