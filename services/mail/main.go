package mail

import (
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/mail/config"
	"github.com/HackIllinois/api/services/mail/controller"
	"github.com/HackIllinois/api/services/mail/service"
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
	controller.SetupController(router.PathPrefix("/mail"))

	log.Fatal(apiserver.StartServer(config.MAIL_PORT, router, "mail"))
}
