package main

import (
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/user/config"
	"github.com/HackIllinois/api/services/user/controller"
	"github.com/gorilla/mux"
	"log"
)

func main() {
	router := mux.NewRouter()
	controller.SetupController(router.PathPrefix("/user"))

	log.Fatal(apiserver.StartServer(config.USER_PORT, router, "user"))
}
