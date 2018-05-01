package config

import (
    "github.com/arbor-dev/arbor/proxy"
    "github.com/arbor-dev/arbor/security"
)

const TestURL string = "http://localhost:8001"

func LoadArborConfig() {
    security.AccessLogLocation = "log/access.log"
    security.ClientRegistryLocation = "clients.db"
    proxy.AccessControlPolicy = "*"
}
