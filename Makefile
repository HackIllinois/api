BASE_PACKAGE := github.com/HackIllinois/api
REPO_ROOT := $(shell git rev-parse --show-toplevel)

SERVICES := auth user registration decision rsvp checkin upload mail event stat notifications
GATEWAYS := gateway

UTILITIES := accountgen tokengen

.PHONY: all
all: api utilities

.PHONY: api
api:
	@echo 'Building api'
	@mkdir -p $(REPO_ROOT)/bin
	@go build -i -o $(REPO_ROOT)/bin/hackillinois-api $(BASE_PACKAGE)

.PHONY: test
test:
	@echo 'Testing services'
	@$(foreach service,$(SERVICES),go test $(BASE_PACKAGE)/services/$(service)/tests;)
	@echo 'Testing gateway'
	@$(foreach gateway,$(GATEWAYS),go test $(BASE_PACKAGE)/$(gateway)/tests;)

.PHONY: utilities
utilities:
	@echo 'Building utilities'
	@mkdir -p $(REPO_ROOT)/bin
	@$(foreach utility,$(UTILITIES),go build -i -o $(REPO_ROOT)/bin/hackillinois-utility-$(utility) $(BASE_PACKAGE)/utilities/$(utility);)

.PHONY: setup
setup: all
	@echo 'Generating API admin account'
	@export HI_CONFIG=file://$(REPO_ROOT)/config/dev_config.json; \
	$(REPO_ROOT)/bin/hackillinois-utility-accountgen
	@echo 'Generating token for admin account'
	@export HI_CONFIG=file://$(REPO_ROOT)/config/dev_config.json; \
	$(REPO_ROOT)/bin/hackillinois-utility-tokengen

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: docs
docs:
	$(MAKE) -C $(BASE_DIR)/documentation build