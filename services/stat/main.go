package main

import (
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/stat/config"
	"github.com/HackIllinois/api/services/stat/controller"
	"github.com/gorilla/mux"
	"log"
)

func main() {
	router := mux.NewRouter()
	controller.SetupController(router.PathPrefix("/stat"))

	log.Fatal(apiserver.StartServer(config.STAT_PORT, router))
}
