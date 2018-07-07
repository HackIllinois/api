package main

import (
	"github.com/HackIllinois/api-gateway/config"
	"github.com/HackIllinois/api-gateway/services"
	"github.com/arbor-dev/arbor"
)

func main() {
	config.LoadArborConfig()
	Routes := services.RegisterAPIs()
	arbor.Boot(Routes, "0.0.0.0", config.GATEWAY_PORT)
}
