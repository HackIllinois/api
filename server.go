package main

import (
	"./config"
	"./services"
	"github.com/ASankaran/arbor"
)

func main() {
	config.LoadArborConfig()
	Routes := services.RegisterAPIs()
	arbor.Boot(Routes, "0.0.0.0", 8000)
}
