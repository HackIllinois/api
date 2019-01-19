package gateway

import (
	"github.com/HackIllinois/api/gateway/config"
	"github.com/HackIllinois/api/gateway/services"
	"github.com/arbor-dev/arbor/server"
	"fmt"
	"os"
)

func Entry() {
	err := config.Initialize()

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)

	}

	config.LoadArborConfig()

	Routes := services.RegisterAPIs()
	server.StartUnsecuredServer(Routes.ToServiceRoutes(), "0.0.0.0", config.GATEWAY_PORT)
}
