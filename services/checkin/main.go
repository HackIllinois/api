package checkin

import (
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/checkin/config"
	"github.com/HackIllinois/api/services/checkin/controller"
	"github.com/HackIllinois/api/services/checkin/service"
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
	controller.SetupController(router.PathPrefix("/checkin"))

	log.Fatal(apiserver.StartServer(config.CHECKIN_PORT, router, "checkin"))
}
