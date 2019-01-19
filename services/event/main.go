package event

import (
	"fmt"
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/event/config"
	"github.com/HackIllinois/api/services/event/controller"
	"github.com/HackIllinois/api/services/event/service"
	"github.com/gorilla/mux"
	"log"
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
	controller.SetupController(router.PathPrefix("/event"))

	log.Fatal(apiserver.StartServer(config.EVENT_PORT, router, "event"))
}
