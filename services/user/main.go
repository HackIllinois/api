package user

import (
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/user/config"
	"github.com/HackIllinois/api/services/user/controller"
	"github.com/HackIllinois/api/services/user/service"
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
	controller.SetupController(router.PathPrefix("/user"))

	log.Fatal(apiserver.StartServer(config.USER_PORT, router, "user"))
}
