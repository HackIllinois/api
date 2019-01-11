BASE_PACKAGE := github.com/HackIllinois/api

SERVICES := auth user registration decision rsvp checkin upload mail event stat notifications
GATEWAYS := gateway

UTILITIES := accountgen tokengen

.PHONY: all
all: api utilities

.PHONY: api
api:
	@echo 'Building api'
	@go build -i -o hackillinois-api $(BASE_PACKAGE)

.PHONY: test
test:
	@echo 'Testing services'
	@$(foreach service,$(SERVICES),go test $(BASE_PACKAGE)/services/$(service)/tests;)
	@echo 'Testing gateway'
	@$(foreach gateway,$(GATEWAYS),go test $(BASE_PACKAGE)/$(gateway)/tests;)

.PHONY: utilities
utilities:
	@echo 'Building utilities'
	@$(foreach utility,$(UTILITIES),go build $(BASE_PACKAGE)/utilities/$(utility);)

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: docs
docs:
	$(MAKE) -C $(BASE_DIR)/documentation build