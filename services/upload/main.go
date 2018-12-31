package main

import (
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/upload/config"
	"github.com/HackIllinois/api/services/upload/controller"
	"github.com/gorilla/mux"
	"log"
)

func main() {
	router := mux.NewRouter()
	controller.SetupController(router.PathPrefix("/upload"))

	log.Fatal(apiserver.StartServer(config.UPLOAD_PORT, router))
}
