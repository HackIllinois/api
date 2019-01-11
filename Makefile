BASE_PACKAGE := github.com/HackIllinois/api

SERVICES := auth user registration decision rsvp checkin upload mail event stat notifications
GATEWAYS := gateway

.PHONY: api
api:
	go build -i -o hackillinois-api $(BASE_PACKAGE)

.PHONY: test
test:
	@echo 'Testing services'
	@$(foreach service,$(SERVICES),go test $(BASE_PACKAGE)/services/$(service)/tests;)
	@echo 'Testing gateway'
	@$(foreach gateway,$(GATEWAYS),go test $(BASE_PACKAGE)/$(gateway)/tests;)

.PHONY: fmt
fmt:
	go fmt ./...
