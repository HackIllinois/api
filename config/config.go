package config

import (
    "github.com/ASankaran/arbor/proxy"
    "github.com/ASankaran/arbor/security"
)

const TestURL string = "http://localhost:8001"

func LoadArborConfig() {
    security.AccessLogLocation = "log/access.log"
    security.ClientRegistryLocation = "clients.db"
    proxy.AccessControlPolicy = "*"
}
