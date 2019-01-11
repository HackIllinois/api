package gateway

import (
	"github.com/HackIllinois/api/gateway/config"
	"github.com/HackIllinois/api/gateway/services"
	"github.com/arbor-dev/arbor/server"
)

func Entry() {
	config.LoadArborConfig()
	Routes := services.RegisterAPIs()
	server.StartUnsecuredServer(Routes.ToServiceRoutes(), "0.0.0.0", config.GATEWAY_PORT)
}
