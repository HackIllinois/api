package config

import (
	"github.com/arbor-dev/arbor/proxy"
	"github.com/arbor-dev/arbor/security"
)

const TestURL string = "http://localhost:8001"
const AuthURL string = "http://localhost:8002"
const UserURL string = "http://localhost:8003"
const RegistrationURL string = "http://localhost:8004"

func LoadArborConfig() {
	security.AccessLogLocation = "log/access.log"
	security.ClientRegistryLocation = "clients.db"
	proxy.AccessControlPolicy = "*"
}
