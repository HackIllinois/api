package main

import (
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/checkin/config"
	"github.com/HackIllinois/api/services/checkin/controller"
	"github.com/gorilla/mux"
	"log"
)

func main() {
	router := mux.NewRouter()
	controller.SetupController(router.PathPrefix("/checkin"))

	log.Fatal(apiserver.StartServer(config.CHECKIN_PORT, router))
}
