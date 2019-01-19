package upload

import (
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/upload/config"
	"github.com/HackIllinois/api/services/upload/controller"
	"github.com/HackIllinois/api/services/upload/service"
	"github.com/gorilla/mux"
	"log"
	"fmt"
	"os"
)

func Entry() {
	err := config.Initialize()

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)

	}

	err = service.Initialize()

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	router := mux.NewRouter()
	controller.SetupController(router.PathPrefix("/upload"))

	log.Fatal(apiserver.StartServer(config.UPLOAD_PORT, router, "upload"))
}
