package registration

import (
	"fmt"
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/registration/config"
	"github.com/HackIllinois/api/services/registration/controller"
	"github.com/HackIllinois/api/services/registration/service"
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
	controller.SetupController(router.PathPrefix("/registration"))

	log.Fatal(apiserver.StartServer(config.REGISTRATION_PORT, router, "registration"))
}
