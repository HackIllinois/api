package decision

import (
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/decision/config"
	"github.com/HackIllinois/api/services/decision/controller"
	"github.com/HackIllinois/api/services/decision/service"
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
	controller.SetupController(router.PathPrefix("/decision"))

	log.Fatal(apiserver.StartServer(config.DECISION_PORT, router, "decision"))
}
