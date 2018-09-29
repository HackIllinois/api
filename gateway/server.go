package main

import (
	"github.com/ethan-lord/api/gateway/config"
	"github.com/ethan-lord/api/gateway/services"
	"github.com/arbor-dev/arbor"
)

func main() {
	config.LoadArborConfig()
	Routes := services.RegisterAPIs()
	arbor.Boot(Routes, "0.0.0.0", config.GATEWAY_PORT)
}
