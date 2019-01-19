package auth

import (
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/auth/config"
	"github.com/HackIllinois/api/services/auth/controller"
	"github.com/HackIllinois/api/services/auth/service"
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
	controller.SetupController(router.PathPrefix("/auth"))

	log.Fatal(apiserver.StartServer(config.AUTH_PORT, router, "auth"))
}
