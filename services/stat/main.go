package stat

import (
	"fmt"
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/stat/config"
	"github.com/HackIllinois/api/services/stat/controller"
	"github.com/HackIllinois/api/services/stat/service"
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
	controller.SetupController(router.PathPrefix("/stat"))

	log.Fatal(apiserver.StartServer(config.STAT_PORT, router, "stat"))
}
