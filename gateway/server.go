package gateway

import (
	"github.com/HackIllinois/api/gateway/config"
	"github.com/HackIllinois/api/gateway/services"
	"github.com/arbor-dev/arbor/server"
	"log"
)

func Initialize() error {
	err := config.Initialize()

	if err != nil {
		return err
	}

	err = services.Initialize()

	if err != nil {
		return err
	}

	return nil
}

func Entry() {
	err := Initialize()

	if err != nil {
		log.Fatal(err)
	}

	config.LoadArborConfig()

	Routes := services.RegisterAPIs()
	server.StartUnsecuredServer(Routes.ToServiceRoutes(), "0.0.0.0", config.GATEWAY_PORT)
}
